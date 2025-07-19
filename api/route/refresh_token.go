package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	tokenService := auth.NewTokenService()

	group.POST("/refresh_token", func(ctx *gin.Context) {
		var refreshTokenSchema auth.RefreshTokenRequestSchema

		if err := ctx.ShouldBindJSON(&refreshTokenSchema); err != nil {
			ctx.Error(err)
			return
		}

		userId := ctx.GetUint("user_id")
		role := ctx.GetString("role")
		token, err := tokenService.RefreshToken(userId, user.Role(role), env.AccessTokenSecret, refreshTokenSchema.RefreshToken, env.RefreshTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, token)
	})
}
