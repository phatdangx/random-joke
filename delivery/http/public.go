package http

import (
	handler "random-joke/handler"

	"github.com/labstack/echo/v4"
)

func NewPublicRoutes(e *echo.Echo, h *handler.PublicHandler) {
	e.GET("/", h.GetRandomJoke)
}
