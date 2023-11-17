package model_test

import (
	"encoding/json"
	"random-joke/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameMarshaling(t *testing.T) {
	name := model.Name{
		FirstName: "John",
		LastName:  "Doe",
	}

	bytes, err := json.Marshal(name)
	assert.NoError(t, err, "Marshaling should not produce an error")

	var marshaled string = string(bytes)
	assert.Contains(t, marshaled, "John", "Marshaled string should contain first name")
	assert.Contains(t, marshaled, "Doe", "Marshaled string should contain last name")
}

func TestNameUnmarshaling(t *testing.T) {
	jsonString := `{"first_name":"Jane","last_name":"Doe"}`

	var name model.Name
	err := json.Unmarshal([]byte(jsonString), &name)
	assert.NoError(t, err, "Unmarshaling should not produce an error")

	assert.Equal(t, "Jane", name.FirstName, "FirstName should be correctly unmarshaled")
	assert.Equal(t, "Doe", name.LastName, "LastName should be correctly unmarshaled")
}
