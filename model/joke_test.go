package model_test

import (
	"encoding/json"
	"random-joke/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomJokeMarshaling(t *testing.T) {
	joke := model.RandomJoke{
		Type: "success",
		Value: struct {
			Categories []string `json:"categories"`
			ID         int      `json:"id"`
			Joke       string   `json:"joke"`
		}{
			Categories: []string{"tech", "general"},
			ID:         123,
			Joke:       "Why do programmers prefer dark mode? Because light attracts bugs.",
		},
	}

	bytes, err := json.Marshal(joke)
	assert.NoError(t, err, "Marshaling should not produce an error")

	var marshaled string = string(bytes)
	assert.Contains(t, marshaled, "tech", "Marshaled string should contain category")
	assert.Contains(t, marshaled, "Why do programmers prefer dark mode?", "Marshaled string should contain joke")
}

func TestRandomJokeUnmarshaling(t *testing.T) {
	jsonString := `{"type":"success","value":{"categories":["tech","general"],"id":123,"joke":"Why do programmers prefer dark mode? Because light attracts bugs."}}`

	var joke model.RandomJoke
	err := json.Unmarshal([]byte(jsonString), &joke)
	assert.NoError(t, err, "Unmarshaling should not produce an error")

	assert.Equal(t, "success", joke.Type, "Type should be correctly unmarshaled")
	assert.Contains(t, joke.Value.Categories, "tech", "Categories should contain 'tech'")
	assert.Equal(t, 123, joke.Value.ID, "ID should be correctly unmarshaled")
	assert.Equal(t, "Why do programmers prefer dark mode? Because light attracts bugs.", joke.Value.Joke, "Joke should be correctly unmarshaled")
}
