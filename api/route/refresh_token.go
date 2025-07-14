package route

import (
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/refresh_token", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "refresh_token",
		})
	})
}
