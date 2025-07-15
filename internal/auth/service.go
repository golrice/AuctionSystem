package auth

import "time"

func GenerateAccessToken(userId uint, expireDelta time.Duration) *AccessTokenSchema {
	return &AccessTokenSchema{
		AccessToken:           "access_token",
		AccessTokenType:       "Bearer",
		AccessTokenExpireTime: time.Now().Add(expireDelta),
	}
}

func GenerateRefreshToken(userId uint, expireDelta time.Duration) *RefreshTokenSchema {
	return &RefreshTokenSchema{
		RefreshToken:           "refresh_token",
		RefreshTokenType:       "Bearer",
		RefreshTokenExpireTime: time.Now().Add(expireDelta),
	}
}

func GenerateToken(userId uint, accessExpireDelta time.Duration, refreshExpireDelta time.Duration) *FullTokenSchema {
	access_token := GenerateAccessToken(userId, accessExpireDelta)
	refresh_token := GenerateRefreshToken(userId, refreshExpireDelta)

	return &FullTokenSchema{
		AccessTokenSchema:  *access_token,
		RefreshTokenSchema: *refresh_token,
	}
}
