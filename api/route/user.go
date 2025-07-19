package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUserRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	userService := user.NewUserService(user.NewUserRepository(db.Db), timeout)

	// 查询
	group.GET("", func(ctx *gin.Context) {
		user_id := ctx.GetUint("user_id")

		userSchema := user.GetRequestSchema{
			ID: user_id,
		}
		response, err := userService.Get(userSchema)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 更新
	group.PUT("", func(ctx *gin.Context) {
		user_id := ctx.GetUint("user_id")

		userSchema := user.UpdateRequestSchema{
			ID: user_id,
		}
		response, err := userService.Update(userSchema)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 删除
	group.DELETE("", func(ctx *gin.Context) {
		user_id := ctx.GetUint("user_id")

		userSchema := user.DeleteRequestSchema{
			ID: user_id,
		}
		response, err := userService.Delete(userSchema)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})
}
