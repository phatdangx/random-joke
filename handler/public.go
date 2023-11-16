package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UseCase interface {
	FetchRandomJoke(ctx echo.Context) (string, error)
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
	joke, err := h.usecase.FetchRandomJoke(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, joke)
}