package user

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin" // 管理员
	RoleUser  Role = "user"  // 用户
)

type Model struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;type:varchar(255);not null"`         // 用户名
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null"` // 密码

	Email string `json:"email" gorm:"column:email;type:varchar(255)"`            // 邮箱
	Role  Role   `json:"role" gorm:"column:role;default:user;type:varchar(255)"` // 角色 默认是user

	Balance int64 `json:"balance" gorm:"column:balance;default:0"` // 剩余额度
}

type Claims struct {
	UserId uint `json:"user_id"`
	Role   Role `json:"role"`
	jwt.RegisteredClaims
}

func (Model) TableName() string {
	return "users"
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Model{})
}
