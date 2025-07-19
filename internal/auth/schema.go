package auth

import (
	"auctionsystem/internal/common"
	"time"
)

type SignupRequestSchema struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type SignupResponseSchema struct {
	common.ResponseSchema
}

type LoginRequestSchema struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccessTokenSchema struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenType       string    `json:"access_token_type"`
	AccessTokenExpireTime time.Time `json:"access_token_expire_time"`
}

type RefreshTokenSchema struct {
	RefreshToken           string    `json:"refresh_token"`
	RefreshTokenType       string    `json:"refresh_token_type"`
	RefreshTokenExpireTime time.Time `json:"refresh_token_expire_time"`
}

type FullTokenSchema struct {
	AccessTokenSchema
	RefreshTokenSchema
}

type LoginResponseSchema struct {
	common.ResponseSchema
	FullTokenSchema
}

type RefreshTokenRequestSchema struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponseSchema struct {
	common.ResponseSchema
	FullTokenSchema
}
