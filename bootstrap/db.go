package bootstrap

import (
	"fmt"
	"sync"

	"auctionsystem/internal/user"

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

		autoMigrate(db)
	})

	return &DB{Db: db}, err
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}

	return nil
}
