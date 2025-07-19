package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	userService := user.NewUserService(user.NewUserRepository(db.Db), timeout)

	group.POST("/login", func(ctx *gin.Context) {
		var loginSchema auth.LoginRequestSchema

		if err := ctx.ShouldBindJSON(&loginSchema); err != nil {
			ctx.Error(err)
			return
		}

		response, err := userService.Login(&loginSchema, env.AccessTokenSecret, env.RefreshTokenSecret)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})
}
