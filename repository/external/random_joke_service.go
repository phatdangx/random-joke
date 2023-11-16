package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"random-joke/model"

	"github.com/labstack/gommon/log"
)

type RandomJokeService struct {
	client *http.Client
}

func NewRandomJokeService(client *http.Client) *RandomJokeService {
	return &RandomJokeService{
		client: client,
	}
}

func (rjs *RandomJokeService) GetRandomJokeBaseOnName(name model.Name) (joke *model.RandomJoke, err error) {
	// Prepare URL
	url := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", name.FirstName, name.LastName)

	// Send GET req
	resp, err := rjs.client.Get(url)
	if err != nil {
		log.Errorf("get random joke error %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read random joke body error: %v", err)
	}

	// Unmarshal the JSON data into struct
	err = json.Unmarshal(body, joke)
	if err != nil {
		log.Errorf("unmarshall random joke error: %v", err)
	}

	return joke, nil
}
