package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UseCase interface {
	FetchRandomJoke(ctx context.Context) (string, error)
}

type PublicHandler struct {
	usecase UseCase
}

func NewPublicHandler(usecase UseCase) *PublicHandler {
	return &PublicHandler{
		usecase: usecase,
	}
}

func (h *PublicHandler) GetRandomJoke(c echo.Context) error {
	ctx := c.Request().Context()
	joke, err := h.usecase.FetchRandomJoke(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, joke)
}
