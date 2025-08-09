package main

import (
	"auctionsystem/api/route"
	"auctionsystem/bootstrap"
	"time"

	_ "auctionsystem/docs"

	"github.com/gin-gonic/gin"
)

// @title Auction System API
// @version 1.0
// @description Auction System API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := bootstrap.App()
	defer app.Close()

	base := gin.New()

	timeout := time.Second * time.Duration(app.Env.ContextTimeout)

	route.Setup(app.Env, timeout, app.Db, base)

	base.Run(app.Env.ServerAddress)
}
