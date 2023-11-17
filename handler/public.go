package handler

import (
	"context"
	"net/http"
	"random-joke/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

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
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, joke)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrorRandomeService:
		return http.StatusServiceUnavailable
	case model.ErrorJokeService:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
