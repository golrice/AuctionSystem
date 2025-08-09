package route

import (
	"auctionsystem/api/middleware"
	"auctionsystem/bootstrap"
	"auctionsystem/internal/user"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, rootRoute *gin.Engine) {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// swagger 设置
	rootRoute.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootRoute.Use(gin.Recovery())

	if gin.Mode() != gin.ReleaseMode {
		rootRoute.Use(gin.Logger())
	}
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
}
