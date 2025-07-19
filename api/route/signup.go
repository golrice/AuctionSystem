package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	authService := auth.NewAuthService(user.NewUserRepository(db.Db), timeout)

	group.POST("/signup", func(ctx *gin.Context) {
		var signupSchema auth.SignupRequestSchema

		if err := ctx.ShouldBindJSON(&signupSchema); err != nil {
			ctx.Error(errors.New("invalid request"))
			return
		}

		response, err := authService.Signup(&signupSchema)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})
}
