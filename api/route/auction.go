package route

import (
	"auctionsystem/api/route/ws/infra/mq"
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auction/application"
	"auctionsystem/internal/auction/infra/cache"
	"auctionsystem/internal/auction/infra/mixeddb"
	"auctionsystem/internal/auction/infra/persistence"
	"auctionsystem/internal/auction/interface/rest"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuctionRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	cache := cache.NewAuctionCacheImpl(db.Redis)
	persRepo := persistence.NewAuctionPersistencyImpl(db.Db)
	auctionService := application.NewAuctionService(mixeddb.NewAuctionFullRepository(cache, persRepo), timeout, mq.NewRedisRepository(db.Redis))

	auctionController := rest.NewAuctionHandler(*auctionService)

	group.POST("", auctionController.CreateAuction)
	group.GET("/latest", auctionController.ListLatestAuctions)
	group.POST("/bid", auctionController.CreateBid)
	group.GET("/bid/highest", auctionController.GetHighestBid)
}
