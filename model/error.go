package model

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal servcer error")
	// ErrorRandomeService will throw if the name service is crashed
	ErrorRandomeService = errors.New("random name service error")
	// ErrorJokeService will throw if the current action already exists
	ErrorJokeService = errors.New("random joke service error")
)
