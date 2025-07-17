package route

import (
	"auctionsystem/api/middleware"
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, rootRoute *gin.Engine) {
	rootRoute.Use(gin.Logger())
	rootRoute.Use(middleware.CORSMiddleware())
	rootRoute.Use(middleware.ErrorHandleMiddle())

	publicRoute := rootRoute.Group("/")
	NewSignupRoute(env, timeout, db, publicRoute)
	NewLoginRoute(env, timeout, db, publicRoute)

	authRoute := rootRoute.Group("/auth")
	authRoute.Use(middleware.AuthMiddleware(env))
	NewRefreshTokenRoute(env, timeout, db, authRoute)

	api := rootRoute.Group("/api")

	_ = api.Group("/v1")
}
