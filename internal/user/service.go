package user

import (
	"auctionsystem/internal/auth"
	"auctionsystem/internal/common"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo           UserRepository
	contextTimeout time.Duration
}

func NewUserService(repo UserRepository, contextTimeout time.Duration) UserService {
	return &userService{repo: repo, contextTimeout: contextTimeout}
}

func (s *userService) Login(loginSchema *auth.LoginRequestSchema, accessTokenSecret string, refreshTokenSecret string) (*auth.LoginResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model, err := s.repo.Get(ctx, &UserModel{Name: loginSchema.Name})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(loginSchema.Password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	token := auth.GenerateToken(model.ID, time.Hour*24, time.Hour*24*7, accessTokenSecret, refreshTokenSecret)
	return &auth.LoginResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		FullTokenSchema: *token,
	}, nil
}

func (s *userService) Signup(signupSchema *auth.SignupRequestSchema) (*auth.SignupResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	// 检查是否存在过user
	if _, err := s.repo.Get(ctx, &UserModel{Name: signupSchema.Name}); err == nil {
		return nil, errors.New("user already exists")
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(signupSchema.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if err = s.repo.Create(ctx, &UserModel{
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
