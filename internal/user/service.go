package user

import (
	"auctionsystem/internal/auth"
	"auctionsystem/internal/common"
	"time"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, loginSchema *auth.LoginSchema) (auth.LoginResponseSchema, error) {
	var user User
	err := db.Where("name = ? AND password = ?", loginSchema.Name, loginSchema.Password).First(&user).Error
	if err != nil {
		return auth.LoginResponseSchema{
			ResponseSchema: common.ResponseSchema{
				Code: 1,
				Msg:  "invalid login schema",
			},
		}, err
	}

	token := auth.GenerateToken(user.ID, time.Hour*24, time.Hour*24*7)
	return auth.LoginResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		FullTokenSchema: *token,
	}, nil
}

func Signup(db *gorm.DB, signupSchema *auth.SignupSchema) (auth.SignupResponseSchema, error) {
	err := CreateUser(db, &User{
		Name:     signupSchema.Name,
		Password: signupSchema.Password,
	})
	if err != nil {
		return auth.SignupResponseSchema{
			ResponseSchema: common.ResponseSchema{
				Code: 1,
				Msg:  "invalid signup schema",
			},
		}, err
	}

	return auth.SignupResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}
