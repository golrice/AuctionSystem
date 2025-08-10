package auth

import (
	"auctionsystem/internal/common"
	"auctionsystem/internal/user"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func NewTokenService() TokenService {
	return &tokenService{}
}

func NewAuthService(repo user.UserRepository, contextTimeout time.Duration) AuthService {
	return &authService{ts: NewTokenService(), repo: repo, contextTimeout: contextTimeout}
}

func (s *tokenService) generateAccessToken(userId uint, role user.Role, expireDelta time.Duration, secret string) *AccessTokenSchema {
	userCliams := user.Claims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDelta)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userCliams)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil
	}

	return &AccessTokenSchema{
		AccessToken:           tokenString,
		AccessTokenType:       "Bearer",
		AccessTokenExpireTime: time.Now().Add(expireDelta),
	}
}

func (s *tokenService) generateRefreshToken(userId uint, role user.Role, expireDelta time.Duration, secret string) *RefreshTokenSchema {
	userClaims := user.Claims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDelta)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil
	}

	return &RefreshTokenSchema{
		RefreshToken:           tokenString,
		RefreshTokenType:       "Bearer",
		RefreshTokenExpireTime: time.Now().Add(expireDelta),
	}
}

func (s *tokenService) GenerateToken(userId uint, role user.Role, accessExpireDelta time.Duration, refreshExpireDelta time.Duration, accessSecret string, refreshSecret string) *FullTokenSchema {
	access_token := s.generateAccessToken(userId, role, accessExpireDelta, accessSecret)
	refresh_token := s.generateRefreshToken(userId, role, refreshExpireDelta, refreshSecret)

	return &FullTokenSchema{
		AccessTokenSchema:  *access_token,
		RefreshTokenSchema: *refresh_token,
	}
}

func (s *tokenService) RefreshToken(userId uint, role user.Role, accessTokenSecret string, refreshToken string, refreshTokenSecret string, accessTokenExpiryHour int, refreshTokenExpiryHour int) (*RefreshTokenResponseSchema, error) {
	claims, err := s.ValidateToken(refreshToken, refreshTokenSecret)
	if err != nil {
		return nil, err
	}
	if claims.UserId != userId {
		return nil, errors.New("user id not match")
	}

	return &RefreshTokenResponseSchema{
		FullTokenSchema: *s.GenerateToken(claims.UserId, role, time.Duration(accessTokenExpiryHour), time.Duration(refreshTokenExpiryHour), accessTokenSecret, refreshTokenSecret),
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (s *tokenService) ValidateToken(tokenString string, secret string) (*user.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &user.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err // 签名错误/过期/格式问题
	}

	// 检查类型断言和有效性
	if claims, ok := token.Claims.(*user.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token claims")
}

func (s *authService) Login(loginSchema *LoginRequestSchema, accessTokenSecret string, refreshTokenSecret string) (*LoginResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model, err := s.repo.Get(ctx, &user.Model{Name: loginSchema.Name})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(loginSchema.Password))
	if err != nil {
		return nil, errors.New("password incorrect")
	}

	token := s.ts.GenerateToken(model.ID, model.Role, time.Hour*24, time.Hour*24*7, accessTokenSecret, refreshTokenSecret)
	return &LoginResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		FullTokenSchema: *token,
	}, nil
}

func (s *authService) Signup(signupSchema *SignupRequestSchema) (*SignupResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	// 检查是否存在过user
	if _, err := s.repo.Get(ctx, &user.Model{Name: signupSchema.Name}); err == nil {
		return nil, errors.New("user already exists")
	}
	// 检查email合法性
	if err := s.validateEmail(signupSchema.Email); err != nil {
		return nil, err
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(signupSchema.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("create user failed")
	}

	if err = s.repo.Create(ctx, &user.Model{
		Name:     signupSchema.Name,
		Password: string(encryptPassword),
		Email:    signupSchema.Email,
	}); err != nil {
		return nil, errors.New("create user failed")
	}

	return &SignupResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (s *authService) validateEmail(email string) error {
	// 简单检查email合法性
	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	return nil
}
