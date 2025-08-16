package bootstrap

import (
	"auctionsystem/internal/user"
	localLogger "auctionsystem/pkg/logger"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	Db    *gorm.DB
	Redis *redis.Client
}

var (
	dbOnce sync.Once
)

func NewDb(cfg *Env) (*DB, error) {
	var err error
	var db *gorm.DB
	var redisClient *redis.Client

	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: localLogger.NewFileLogger("gorm.log", time.Second, gormLogger.Warn),
		})
		if err != nil {
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			return
		}
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetConnMaxLifetime(time.Hour)

		user.AutoMigrate(db)

		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			Password: cfg.RedisPass,
			DB:       cfg.RedisDatabase,
		})
		_, err = redisClient.Ping(context.Background()).Result()
		if err != nil {
			return
		}
	})

	return &DB{Db: db, Redis: redisClient}, err
}

func AutoMigrate(db *DB, models ...interface{}) {
	db.Db.AutoMigrate(models...)
}
