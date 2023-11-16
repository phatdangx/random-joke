package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ConfigSchema struct {
	Port            string `mapstructure:"port"`
	ExternalService struct {
		RandomName string `mapstructure:"random_name"`
		RandomJoke string `mapstructure:"random_joke"`
	} `mapstructure:"external_service"`
}

var Config ConfigSchema

func NewSchema() *ConfigSchema {
	schema := new(ConfigSchema)
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")             // Look for config in current directory
	config.AddConfigPath("config/")       // Optionally look for config in the working directory.
	config.AddConfigPath("../config/")    // Look for config needed for tests.
	config.AddConfigPath("../")           // Look for config needed for tests.
	config.AddConfigPath("../../config/") // used for integration_test

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("config error %s", err))
	}

	err = config.Unmarshal(&schema)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("config error: %s", err))
	}

	return schema
}

func init() {
	Config = *NewSchema()
}
