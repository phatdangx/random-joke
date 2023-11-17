package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"random-joke/config"
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

func (rjs *RandomJokeService) GetRandomJokeBaseOnName(name *model.Name) (joke *model.RandomJoke, err error) {
	// Parameters
	params := url.Values{}
	params.Add("limitTo", "nerdy")
	params.Add("firstName", name.FirstName)
	params.Add("lastName", name.LastName) // Replace with actual last name

	// Construct the complete URL with parameters
	url := fmt.Sprintf("%s?%s", config.Config.ExternalService.RandomJoke, params.Encode())

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
	err = json.Unmarshal(body, &joke)
	if err != nil {
		log.Errorf("unmarshall random joke error: %v", err)
	}

	return joke, nil
}
