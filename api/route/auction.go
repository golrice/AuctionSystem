package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auction"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuctionRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	auctionService := auction.NewAuctionService(auction.NewAuctionRepository(db.Db), timeout)

	// 创建某次拍卖
	group.POST("", func(ctx *gin.Context) {
		var request auction.CreateRequestSchema

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(errors.New("invalid request"))
			return
		}

		request.UserID = ctx.GetUint("user_id")
		response, err := auctionService.Create(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 获取某次拍卖
	group.GET("", func(ctx *gin.Context) {
		var request auction.GetRequestSchema

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(errors.New("invalid request"))
			return
		}

		response, err := auctionService.Get(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 更新拍卖
	group.PUT("", func(ctx *gin.Context) {
		var request auction.UpdateRequestSchema

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(errors.New("invalid request"))
			return
		}

		response, err := auctionService.Update(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})

	// 删除拍卖
	group.DELETE("", func(ctx *gin.Context) {
		var request auction.DeleteRequestSchema

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.Error(errors.New("invalid request"))
			return
		}

		response, err := auctionService.Delete(&request)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, response)
	})
}
