package external_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"random-joke/config"
	"random-joke/model"
	"random-joke/repository/external"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomJokeBaseOnName(t *testing.T) {
	// Create a sample joke response
	sampleJoke := model.RandomJoke{
		Type: "success",
		Value: struct {
			Categories []string `json:"categories"`
			ID         int      `json:"id"`
			Joke       string   `json:"joke"`
		}{
			Categories: []string{"animal", "dad"},
			ID:         123,
			Joke:       "Why did the chicken cross the road? To get to the other side!",
		},
	}

	jokeBytes, _ := json.Marshal(sampleJoke)

	// Setup a mock server
	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(jokeBytes)
		}))
	defer mockServer.Close()

	// Mock the config with the mock server URL
	config.Config.ExternalService.RandomJoke = mockServer.URL

	// Create a RandomJokeService with a real http client
	rjs := external.NewRandomJokeService(http.DefaultClient)

	// Prepare test name
	testName := &model.Name{FirstName: "John", LastName: "Doe"}

	// Test the method
	joke, err := rjs.GetRandomJokeBaseOnName(testName)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, joke)
	assert.Equal(t, sampleJoke.Value.Joke, joke.Value.Joke)
}
