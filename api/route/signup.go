package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/signup", func(ctx *gin.Context) {
		var signupSchema auth.SignupSchema

		if err := ctx.ShouldBindJSON(&signupSchema); err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid signup schema",
			})
			return
		}

		response, err := user.Signup(db.Db, &signupSchema)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid signup schema",
			})
			return
		}

		ctx.JSON(200, response)
	})
}
