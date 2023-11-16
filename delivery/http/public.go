package http

import (
	handler "random-joke/handler"

	"github.com/labstack/echo/v4"
)

// NewArticleHandler will initialize the articles/ resources endpoint
func NewPublicRoutes(e *echo.Echo, h *handler.PublicHandler) {
	e.GET("/", h.GetRandomJoke)
}
