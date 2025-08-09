package route

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auction/application"
	"auctionsystem/internal/auction/infra/persistence"
	"auctionsystem/internal/auction/interface/rest"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuctionRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	auctionService := application.NewAuctionService(persistence.NewAuctionRepositoryImpl(db.Db), timeout)
	auctionController := rest.NewAuctionHandler(*auctionService)

	group.POST("", auctionController.CreateAuction)
	group.GET("/latest", auctionController.ListLatestAuctions)
	group.POST("/bid", auctionController.CreateBid)
	group.GET("/bid/higest", auctionController.GetHigestBid)
}
