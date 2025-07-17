package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
}
