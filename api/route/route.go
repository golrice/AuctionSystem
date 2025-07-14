package route

import (
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, rootRoute *gin.Engine) {
	publicRoute := rootRoute.Group("")
	NewSignupRoute(env, timeout, db, publicRoute)
	NewLoginRoute(env, timeout, db, publicRoute)
	NewRefreshTokenRoute(env, timeout, db, publicRoute)

	api := rootRoute.Group("/api")

	_ = api.Group("/v1")
}
