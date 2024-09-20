package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config holds the configuration values for the application.
type Config struct {
	Port string `json:"port"`
	URL  string `json:"url"`
}

// LoadConfig loads the configuration from the provided JSON file.
func LoadConfig() (*Config, error) {
	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = "config.json"
		log.Printf("CONFIG_FILE environment variable not set. Using default: %s", configFile)
	} else {
		log.Printf("Using CONFIG_FILE: %s", configFile)
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
