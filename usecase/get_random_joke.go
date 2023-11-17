package usecase

import (
	"context"
	"random-joke/model"
)

type NameRepository interface {
	GetRandomName() (name *model.Name, err error)
}

type JokeRepository interface {
	GetRandomJokeBaseOnName(name *model.Name) (joke *model.RandomJoke, err error)
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

func (uc *UseCase) FetchRandomJoke(ctx context.Context) (string, error) {
	name, err := uc.name.GetRandomName()
	if err != nil {
		return "", model.ErrorRandomeService
	}

	joke, err := uc.joke.GetRandomJokeBaseOnName(name)
	if err != nil {
		return "", model.ErrorJokeService
	}

	return joke.Value.Joke, nil
}
