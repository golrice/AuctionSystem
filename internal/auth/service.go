package auth

import (
	"auctionsystem/internal/common"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId uint, expireDelta time.Duration, secret string) *AccessTokenSchema {
	userCliams := UserClaims{
		UserId: userId,
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

func GenerateRefreshToken(userId uint, expireDelta time.Duration, secret string) *RefreshTokenSchema {
	userClaims := UserClaims{
		UserId: userId,
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

func GenerateToken(userId uint, accessExpireDelta time.Duration, refreshExpireDelta time.Duration, accessSecret string, refreshSecret string) *FullTokenSchema {
	access_token := GenerateAccessToken(userId, accessExpireDelta, accessSecret)
	refresh_token := GenerateRefreshToken(userId, refreshExpireDelta, refreshSecret)

	return &FullTokenSchema{
		AccessTokenSchema:  *access_token,
		RefreshTokenSchema: *refresh_token,
	}
}

func RefreshToken(refreshToken string, refreshSecret string) (*RefreshTokenResponseSchema, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return &RefreshTokenResponseSchema{
			ResponseSchema: common.ResponseSchema{
				Code: 0,
				Msg:  "success",
			},
			FullTokenSchema: *GenerateToken(claims.UserId, time.Hour, time.Hour*24, refreshSecret, refreshSecret),
		}, nil
	}

	return nil, errors.New("invalid refresh token")
}
