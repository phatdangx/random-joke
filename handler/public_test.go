package handler_test

import (
	"net/http"
	"net/http/httptest"
	"random-joke/handler"
	"random-joke/mocks"
	"random-joke/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRandomJoke(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUseCase := new(mocks.UseCase)
		h := handler.NewPublicHandler(mockUseCase)
		e := echo.New()
		// Set expectations
		mockUseCase.On(
			"FetchRandomJoke",
			mock.AnythingOfType("*context.emptyCtx")).Return("Funny joke", nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assert no error on handler
		assert.NoError(t, h.GetRandomJoke(c))

		// Assert correct status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Assert response body
		assert.Contains(t, rec.Body.String(), "Funny joke")

		mockUseCase.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		mockUseCase := new(mocks.UseCase)
		h := handler.NewPublicHandler(mockUseCase)
		e := echo.New()
		// Set expectations
		mockUseCase.On(
			"FetchRandomJoke",
			mock.AnythingOfType("*context.emptyCtx")).Return("", model.ErrorRandomeService)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		_ = h.GetRandomJoke(c)

		// Assert correct status code for error
		assert.Equal(t, http.StatusServiceUnavailable, rec.Code)

		mockUseCase.AssertExpectations(t)
	})
}
