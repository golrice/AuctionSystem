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

		token, err := auth.RefreshToken(refreshTokenSchema.RefreshToken, env)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, token)
	})
}
