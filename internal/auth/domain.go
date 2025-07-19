package auth

import (
	"auctionsystem/internal/user"
	"time"
)

type tokenService struct {
}

type TokenService interface {
	GenerateToken(userId uint, role user.Role, accessExpireDelta time.Duration, refreshExpireDelta time.Duration, accessSecret string, refreshSecret string) *FullTokenSchema
	RefreshToken(userId uint, role user.Role, accessTokenSecret string, refreshToken string, refreshTokenSecret string, accessTokenExpiryHour int, refreshTokenExpiryHour int) (*RefreshTokenResponseSchema, error)
	ValidateToken(token string, secret string) (*user.Claims, error)
}

type authService struct {
	ts             TokenService
	repo           user.UserRepository
	contextTimeout time.Duration
}

type AuthService interface {
	Login(request *LoginRequestSchema, accessSecret string, refreshSecret string) (*LoginResponseSchema, error)
	Signup(signupSchema *SignupRequestSchema) (*SignupResponseSchema, error)
}
