package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
				})
			}
		}()
		ctx.Next()

		if len(ctx.Errors) > 0 {
			lastErr := ctx.Errors.Last()
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": lastErr.Error(),
			})
			return
		}
	}
}
