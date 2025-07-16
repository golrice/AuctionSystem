package user

import (
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetUser(db *gorm.DB, user *User) (*User, error) {
	var result User
	err := db.Where(user).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateUser(db *gorm.DB, user *User) error {
	return db.Model(&User{}).Updates(user).Error
}

func DeleteUser(db *gorm.DB, id int) error {
	return db.Delete(&User{}, id).Error
}
