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

func TestGetRandomName(t *testing.T) {
	// Mock response
	mockName := &model.Name{FirstName: "John", LastName: "Doe"}
	mockNameJSON, _ := json.Marshal(mockName)

	// Create a mock server
	mockServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(mockNameJSON)
		}))
	defer mockServer.Close()

	// Mock the external service URL with the mock server URL
	config.Config.ExternalService.RandomName = mockServer.URL

	// Create a NameService with a real http client
	ns := external.NewNameService(http.DefaultClient)

	// Call the method
	name, err := ns.GetRandomName()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, name)
	assert.Equal(t, mockName.FirstName, name.FirstName)
	assert.Equal(t, mockName.LastName, name.LastName)
}
