package route

import (
	"auctionsystem/api/middleware"
	"auctionsystem/bootstrap"
	"auctionsystem/internal/user"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, rootRoute *gin.Engine) {
	rootRoute.Use(middleware.ErrorHandleMiddleware())
	rootRoute.Use(gin.Logger())
	rootRoute.Use(middleware.CORSMiddleware())

	// 公共路由
	publicRoute := rootRoute.Group("/")
	NewSignupRoute(env, timeout, db, publicRoute)
	NewLoginRoute(env, timeout, db, publicRoute)

	// 管理员路由
	adminRoute := rootRoute.Group("/admin")
	adminRoute.Use(middleware.JWTMiddleware(env))
	adminRoute.Use(middleware.RequireRole(user.RoleAdmin))

	// 认证路由
	authRoute := rootRoute.Group("/auth")
	authRoute.Use(middleware.JWTMiddleware(env))
	authRoute.Use(middleware.RequireRole(user.RoleAdmin, user.RoleUser))
	NewRefreshTokenRoute(env, timeout, db, authRoute)

	// api路由
	apiRoute := rootRoute.Group("/api")
	apiRoute.Use(middleware.JWTMiddleware(env))
	apiRoute.Use(middleware.RequireRole(user.RoleAdmin, user.RoleUser))

	// 用户api
	NewUserRoute(env, timeout, db, apiRoute.Group("/user"))
	// 拍卖api
	NewAuctionRoute(env, timeout, db, apiRoute.Group("/auction"))
	// 出价api
	NewBidRoute(env, timeout, db, apiRoute.Group("/bid"))
}
