package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"random-joke/mocks"
	"random-joke/model"
	"random-joke/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchRandomJoke(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockNameRepo := new(mocks.NameRepository)
		mockJokeRepo := new(mocks.JokeRepository)
		mockName := model.Name{
			FirstName: "John",
			LastName:  "Doe",
		}
		mockJoke := model.RandomJoke{
			Type: "success",
			Value: struct {
				Categories []string `json:"categories"`
				ID         int      `json:"id"`
				Joke       string   `json:"joke"`
			}{
				Categories: []string{"animal", "dad"},
				ID:         123,
				Joke:       "John Doe is funny",
			},
		}
		mockNameRepo.On("GetRandomName").Return(&mockName, nil)
		mockJokeRepo.On("GetRandomJokeBaseOnName", mock.AnythingOfType("*model.Name")).Return(&mockJoke, nil)
		uc := usecase.NewJokeUseCase(
			mockNameRepo,
			mockJokeRepo,
		)

		joke, err := uc.FetchRandomJoke(context.TODO())
		assert.NoError(t, err)
		assert.Contains(
			t, joke,
			fmt.Sprintf("%s %s", mockName.FirstName, mockName.LastName),
			"Random joke service does not contian name from name service")
		mockNameRepo.AssertExpectations(t)
		mockJokeRepo.AssertExpectations(t)
	})
	t.Run("name-repo-failure", func(t *testing.T) {
		mockNameRepo := new(mocks.NameRepository)
		mockJokeRepo := new(mocks.JokeRepository)
		mockNameRepo.On("GetRandomName").Return(nil, errors.New("Unexpexted Name Service Error"))
		uc := usecase.NewJokeUseCase(
			mockNameRepo,
			mockJokeRepo,
		)

		joke, err := uc.FetchRandomJoke(context.TODO())
		assert.Error(t, err)
		assert.Equal(t, "", joke)
		mockNameRepo.AssertExpectations(t)
		mockJokeRepo.AssertNotCalled(t, "GetRandomJokeBaseOnName", mock.Anything)
	})
	t.Run("joke-repo-failure", func(t *testing.T) {
		mockNameRepo := new(mocks.NameRepository)
		mockJokeRepo := new(mocks.JokeRepository)

		mockName := model.Name{
			FirstName: "John",
			LastName:  "Doe",
		}
		mockNameRepo.On("GetRandomName").Return(&mockName, nil)
		mockJokeRepo.On("GetRandomJokeBaseOnName", mock.AnythingOfType("*model.Name")).Return(nil, errors.New("Unexpexted Joke Service Error"))
		uc := usecase.NewJokeUseCase(
			mockNameRepo,
			mockJokeRepo,
		)

		joke, err := uc.FetchRandomJoke(context.TODO())
		assert.Error(t, err)
		assert.Equal(t, "", joke)
		mockNameRepo.AssertExpectations(t)
		mockJokeRepo.AssertExpectations(t)
	})
}
