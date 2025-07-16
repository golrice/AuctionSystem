package route

import (
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, rootRoute *gin.Engine) {
	// publicRoute := rootRoute.Group("")

	authRoute := rootRoute.Group("/auth")
	NewSignupRoute(env, timeout, db, authRoute)
	NewLoginRoute(env, timeout, db, authRoute)
	NewRefreshTokenRoute(env, timeout, db, authRoute)

	api := rootRoute.Group("/api")

	_ = api.Group("/v1")
}
