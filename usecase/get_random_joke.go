package usecase

import (
	"random-joke/model"

	"github.com/labstack/echo/v4"
)

type NameRepository interface {
	GetRandomName() (name *model.Name, err error)
}

type JokeRepository interface {
	GetRandomJokeBaseOnName(name model.Name) (joke model.RandomJoke, err error)
}

type UseCase struct {
	name NameRepository
	joke JokeRepository
}

func NewJokeUseCase(nameRepo NameRepository, jokeRepo JokeRepository) *UseCase {
	return &UseCase{
		name: nameRepo,
		joke: jokeRepo,
	}
}

func (uc *UseCase) FetchRandomJoke(c echo.Context) (string, error) {
	name, err := uc.name.GetRandomName()
	if err != nil {
		return "", err
	}

	joke, err := uc.joke.GetRandomJokeBaseOnName(name)
	if err != nil {
		return "", err
	}

	return joke.Value.Joke, nil
}
