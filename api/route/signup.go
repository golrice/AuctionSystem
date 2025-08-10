package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"auctionsystem/pkg/kernal"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/signup", SignupHandler(db, timeout))
}

// SignupHandler 注册接口
// @Summary 注册
// @Description 注册新用户
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body auth.SignupRequestSchema true "注册请求体"
// @Success 200 {object} auth.SignupResponseSchema
// @Failure 400 {object} kernal.ErrorResult
// @Router /auth/signup [post]
func SignupHandler(db *bootstrap.DB, timeout time.Duration) gin.HandlerFunc {
	authService := auth.NewAuthService(user.NewUserRepository(db.Db), timeout)

	return func(ctx *gin.Context) {
		var signupSchema auth.SignupRequestSchema

		if err := ctx.ShouldBindJSON(&signupSchema); err != nil {
			ctx.JSON(http.StatusOK, kernal.ErrorResult{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			})
			return
		}

		response, err := authService.Signup(&signupSchema)
		if err != nil {
			ctx.JSON(http.StatusOK, kernal.ErrorResult{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
