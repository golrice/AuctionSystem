package user

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserModel) TableName() string {
	return "users"
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}
