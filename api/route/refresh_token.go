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

func NewRefreshTokenRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	tokenService := auth.NewTokenService()

	group.POST("/refresh_token", RefreshTokenHandler(env, tokenService))
}

// RefreshTokenHandler 刷新令牌接口
// @Summary Refresh token
// @Description 使用刷新令牌刷新访问令牌
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body auth.RefreshTokenRequestSchema true "刷新令牌请求体"
// @Success 200 {object} kernal.SuccessResult{data=auth.RefreshTokenResponseSchema}
// @Failure 400 {object} kernal.ErrorResult
// @Router /auth/refresh_token [post]
func RefreshTokenHandler(env *bootstrap.Env, tokenService auth.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var refreshTokenSchema auth.RefreshTokenRequestSchema

		if err := ctx.ShouldBindJSON(&refreshTokenSchema); err != nil {
			ctx.JSON(http.StatusBadRequest, kernal.ErrorResult{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			})
			return
		}

		userId := ctx.GetUint("user_id")
		role := ctx.GetString("role")
		token, err := tokenService.RefreshToken(userId, user.Role(role), env.AccessTokenSecret, refreshTokenSchema.RefreshToken, env.RefreshTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, kernal.ErrorResult{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, kernal.NewSuccessResult(token))
	}
}
