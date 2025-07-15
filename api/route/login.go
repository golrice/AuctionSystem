package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/login", func(ctx *gin.Context) {
		// 获取 login 参数, 类型为 LoginSchema
		var loginSchema auth.LoginSchema

		if err := ctx.ShouldBindJSON(&loginSchema); err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid login schema",
			})
			return
		}

		response, err := user.Login(db.Db, &loginSchema)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid login schema",
			})
			return
		}

		ctx.JSON(200, response)
	})
}
