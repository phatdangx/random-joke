package external

import (
	"encoding/json"
	"io"
	"net/http"
	"random-joke/config"
	"random-joke/model"

	"github.com/labstack/gommon/log"
)

type NameService struct {
	client *http.Client
}

func NewNameService(client *http.Client) *NameService {
	return &NameService{
		client: client,
	}
}

func (ns *NameService) GetRandomName() (name *model.Name, err error) {
	// Send GET req
	resp, err := ns.client.Get(config.Config.ExternalService.RandomName)
	if err != nil {
		log.Errorf("get random name %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read random name body error: %v", err)
		return nil, err
	}

	// Unmarshal the JSON data into struct
	err = json.Unmarshal(body, &name)
	if err != nil {
		log.Errorf("unmarshall random name error: %v", err)
		return nil, err
	}

	return name, nil
}
