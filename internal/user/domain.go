package user

import (
	"auctionsystem/internal/auth"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, model *UserModel) error
	Get(ctx context.Context, model *UserModel) (*UserModel, error)
	Update(ctx context.Context, model *UserModel) error
	Delete(ctx context.Context, model *UserModel) error
}

type UserService interface {
	Login(loginSchema *auth.LoginRequestSchema, accessTokenSecret string, refreshTokenSecret string) (*auth.LoginResponseSchema, error)
	Signup(signupSchema *auth.SignupRequestSchema) (*auth.SignupResponseSchema, error)
}
