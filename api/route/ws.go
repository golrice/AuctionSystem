package route

import (
	wsApplication "auctionsystem/api/route/ws/application"
	"auctionsystem/api/route/ws/infra/mq"

	auctionApplication "auctionsystem/internal/auction/application"
	auctionPersistence "auctionsystem/internal/auction/infra/persistence"

	"auctionsystem/api/route/ws/interface/ws"
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewWSRoute(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, group *gin.RouterGroup) {
	hubService := wsApplication.NewAuctionHubService(db.Redis)
	auctionService := auctionApplication.NewAuctionService(auctionPersistence.NewAuctionRepositoryImpl(db.Db), timeout, mq.NewRedisRepository(db.Redis))

	go hubService.RunHub()

	wsController := ws.NewAuctionWSHandler(hubService, auctionService)

	group.GET("", wsController.ServeWS)
}
