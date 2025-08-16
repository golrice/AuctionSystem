package main

import (
	"auctionsystem/api/route"
	"auctionsystem/bootstrap"
	"auctionsystem/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "auctionsystem/docs"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func StartPprof() {
	go func() {
		logger.Logger.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func main() {
	StartPprof()

	app := bootstrap.App()
	defer app.Close()

	base := gin.New()

	timeout := time.Second * time.Duration(app.Env.ContextTimeout)

	route.Setup(app.Env, timeout, app.Db, base)

	httpServer := &http.Server{
		Addr:    app.Env.ServerAddress,
		Handler: base,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Logger.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server Shutdown:", err)
	}
	logger.Logger.Println("Server exiting")
}
