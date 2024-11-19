package types

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/afero"
)

var AppFs = afero.NewOsFs()

type Config struct {
	Version      string       `json:"version"`
	Repositories []Repository `json:"repositories"`
}

type Repository struct {
	Location string `json:"location"`
}

func NewConfig() *Config {
	return &Config{
		Version:      "1.0",
		Repositories: []Repository{},
	}
}

func ReadConfig(configFilePath string) (*Config, error) {
	file, err := afero.ReadFile(AppFs, configFilePath)
	if err != nil {
		fmt.Println("Error reading config file")
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error parsing config file")
		return nil, err
	}

	return &config, nil

}

func WriteConfig(config *Config, configFilePath string) {
	configJson, err := json.Marshal(&config)
	if err != nil {
		fmt.Println("Error creating config file")
		return
	}

	err = afero.WriteFile(AppFs, configFilePath, configJson, 0644)
	if err != nil {
		fmt.Println("Error writing config file")
		return
	}
}
