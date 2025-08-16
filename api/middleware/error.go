package middleware

import (
	"auctionsystem/pkg/kernal"
	"auctionsystem/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Printf("Panic: %v", err)
				ctx.JSON(http.StatusInternalServerError, kernal.NewErrorResult(http.StatusInternalServerError, "Internal server error"))
			}
		}()
		ctx.Next()

		if len(ctx.Errors) > 0 {
			lastErr := ctx.Errors.Last()
			ctx.JSON(http.StatusBadRequest, kernal.NewErrorResult(http.StatusBadRequest, lastErr.Error()))
			return
		}
	}
}
