package route

import (
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "login",
		})
	})
}
