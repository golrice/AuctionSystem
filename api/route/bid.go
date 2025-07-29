package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/bid"
	"time"

	"github.com/gin-gonic/gin"
)

func NewBidRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	bidService := bid.NewService(bid.NewMemoryRepository(db.Redis), bid.NewPersistentRepository(db.Db), timeout)

	// 创建一次出价
	group.POST("/create", func(ctx *gin.Context) {
		var request bid.CreateRequestSchema
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(err)
			return
		}

		response, err := bidService.Create(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 查询出价
	group.GET("/list", func(ctx *gin.Context) {
		var request bid.ListRequestSchema
		if err := ctx.ShouldBindQuery(&request); err != nil {
			ctx.Error(err)
			return
		}

		response, err := bidService.List(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 删除出价
	group.DELETE("/delete", func(ctx *gin.Context) {
		var request bid.DeleteRequestSchema
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(err)
			return
		}

		response, err := bidService.Delete(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 更新出价
	group.PUT("/update", func(ctx *gin.Context) {
		var request bid.UpdateRequestSchema
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(err)
			return
		}

		response, err := bidService.Update(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 获取出价详情
	group.GET("/detail", func(ctx *gin.Context) {
		var request bid.DetailRequestSchema
		if err := ctx.ShouldBindQuery(&request); err != nil {
			ctx.Error(err)
			return
		}

		response, err := bidService.Detail(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})
}
