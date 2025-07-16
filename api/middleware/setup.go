package middleware

import (
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *bootstrap.DB, base *gin.Engine) {
	base.Use(gin.Logger())

	base.Use(gin.Recovery())
}
