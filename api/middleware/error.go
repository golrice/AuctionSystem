package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error: %v", err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
				})
			}
		}()
		ctx.Next()
	}
}
