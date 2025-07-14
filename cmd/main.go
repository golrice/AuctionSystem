package main

import (
	"auctionsystem/api/route"
	"auctionsystem/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.Close()

	base := gin.Default()

	timeout := time.Second * time.Duration(app.Env.ContextTimeout)
	route.Setup(app.Env, timeout, app.Db, base)

	base.Run(app.Env.ServerAddress)
}
