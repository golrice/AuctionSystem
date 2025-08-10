package route

import (
	wsApplication "auctionsystem/api/route/ws/application"
	"auctionsystem/api/route/ws/infra/mq"
	"auctionsystem/internal/auction/application"
	"auctionsystem/internal/auction/infra/cache"
	"auctionsystem/internal/auction/infra/mixeddb"
	"auctionsystem/internal/auction/infra/persistence"

	"auctionsystem/api/route/ws/interface/ws"
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewWSRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	hubService := wsApplication.NewAuctionHubService(db.Redis)

	cache := cache.NewAuctionCacheImpl(db.Redis)
	persRepo := persistence.NewAuctionPersistencyImpl(db.Db)
	auctionService := application.NewAuctionService(mixeddb.NewAuctionFullRepository(cache, persRepo), timeout, mq.NewRedisRepository(db.Redis))

	go hubService.RunHub()

	wsController := ws.NewAuctionWSHandler(hubService, auctionService)

	group.GET("", wsController.ServeWS)
}
