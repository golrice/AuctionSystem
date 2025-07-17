package auth

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/common"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

func generateAccessToken(userId uint, expireDelta time.Duration, secret string) *AccessTokenSchema {
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

func generateRefreshToken(userId uint, expireDelta time.Duration, secret string) *RefreshTokenSchema {
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
	access_token := generateAccessToken(userId, accessExpireDelta, accessSecret)
	refresh_token := generateRefreshToken(userId, refreshExpireDelta, refreshSecret)

	return &FullTokenSchema{
		AccessTokenSchema:  *access_token,
		RefreshTokenSchema: *refresh_token,
	}
}

func RefreshToken(refreshToken string, env *bootstrap.Env) (*RefreshTokenResponseSchema, error) {
	claims, err := ValidateToken(refreshToken, env.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &RefreshTokenResponseSchema{
		FullTokenSchema: *GenerateToken(claims.UserId, time.Duration(env.AccessTokenExpiryHour), time.Duration(env.RefreshTokenExpiryHour), env.AccessTokenSecret, env.RefreshTokenSecret),
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func ValidateToken(tokenString string, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token claims")
}
