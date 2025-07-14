package user

import (
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
