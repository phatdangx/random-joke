package app

import (
	"context"
	"os"
	"os/signal"
	"random-joke/delivery/http"
	"random-joke/handler"
	"random-joke/repository/external"
	"random-joke/usecase"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	nethttp "net/http"
)

func Start() {
	e := echo.New()

	// init http client with timeout 3 seconds
	client := &nethttp.Client{
		Timeout: time.Second * 3,
	}

	// register external repository
	nameRepo := external.NewNameService(client)
	jokeRepo := external.NewRandomJokeService(client)

	// register usecase
	usecase := usecase.NewJokeUseCase(nameRepo, jokeRepo)

	// register handler
	publicHandler := handler.NewPublicHandler(usecase)

	// register routes
	http.NewPublicRoutes(e, publicHandler)

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
