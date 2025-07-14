package route

import (
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "signup",
		})
	})
}
