package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	group.POST("/login", LoginHandler(env, db, timeout))
}

// LoginHandler 登录接口
// @Summary 登录
// @Description 登录获取访问令牌和刷新令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body auth.LoginRequestSchema true "登录请求体"
// @Success 200 {object} auth.LoginResponseSchema
// @Failure 400 {object} kernal.ErrorResult
// @Router /auth/login [post]
func LoginHandler(env *bootstrap.Env, db *bootstrap.DB, timeout time.Duration) gin.HandlerFunc {
	authService := auth.NewAuthService(user.NewUserRepository(db.Db), timeout)

	return func(ctx *gin.Context) {
		var loginSchema auth.LoginRequestSchema

		if err := ctx.ShouldBindJSON(&loginSchema); err != nil {
			ctx.Error(err)
			return
		}

		response, err := authService.Login(&loginSchema, env.AccessTokenSecret, env.RefreshTokenSecret)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	}
}
