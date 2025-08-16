package middleware

import (
	"auctionsystem/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		// 高亮显示
		logger.Logger.Printf("[%s] %s %s %d %v",
			clientIP, method, path, status, latency)
	}
}
