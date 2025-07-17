package middleware

import (
	"auctionsystem/bootstrap"
	"auctionsystem/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(env *bootstrap.Env) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format"})
			return
		}

		claims, err := auth.ValidateToken(tokenString, env.AccessTokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
}
