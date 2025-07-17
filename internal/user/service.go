package user

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/common"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, loginSchema *auth.LoginRequestSchema, env *bootstrap.Env) (*auth.LoginResponseSchema, error) {
	user, err := GetUser(db, &User{Name: loginSchema.Name})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginSchema.Password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	token := auth.GenerateToken(user.ID, time.Hour*24, time.Hour*24*7, env.AccessTokenSecret, env.RefreshTokenSecret)
	return &auth.LoginResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		FullTokenSchema: *token,
	}, nil
}

func Signup(db *gorm.DB, signupSchema *auth.SignupRequestSchema) (*auth.SignupResponseSchema, error) {
	// 检查是否存在过user
	if _, err := GetUser(db, &User{Name: signupSchema.Name}); err == nil {
		return nil, errors.New("user already exists")
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(signupSchema.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if err = CreateUser(db, &User{
		Name:     signupSchema.Name,
		Password: string(encryptPassword),
	}); err != nil {
		return nil, errors.New("create user failed")
	}

	return &auth.SignupResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}
