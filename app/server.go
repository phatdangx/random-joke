package app

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Start() {
	e := echo.New()

	// Start server
	go func() {
		port := "8080"
		log.Info("start at port " + port)
		err := e.Start(":" + port)
		if err != nil {
			log.Error(err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
