package main

import (
	"auctionsystem/api/middleware"
	"auctionsystem/api/route"
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.Close()

	base := gin.New()

	timeout := time.Second * time.Duration(app.Env.ContextTimeout)

	middleware.Setup(app.Env, timeout, app.Db, base)
	route.Setup(app.Env, timeout, app.Db, base)

	base.Run(app.Env.ServerAddress)
}
