package bootstrap

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

var (
	dbOnce sync.Once
)

func NewDb(cfg *Env) (*DB, error) {
	var err error
	var db *gorm.DB

	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
	})

	return &DB{Db: db}, err
}

func AutoMigrate(db *DB, models ...interface{}) {
	db.Db.AutoMigrate(models...)
}
