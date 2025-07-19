package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/refresh_token", func(ctx *gin.Context) {
		var refreshTokenSchema auth.RefreshTokenRequestSchema

		if err := ctx.ShouldBindJSON(&refreshTokenSchema); err != nil {
			ctx.Error(err)
			return
		}

		userId := ctx.GetUint("user_id")
		token, err := auth.RefreshToken(userId, env.AccessTokenSecret, refreshTokenSchema.RefreshToken, env.RefreshTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, token)
	})
}
