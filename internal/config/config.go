package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Config is the application configuration
	Config struct {
		ServerPort string
	}
)

// LoadConfig loads the application configuration from the .env file
func LoadConfig() (config *Config, err error) {
	if err = godotenv.Load(os.ExpandEnv(".env")); err != nil {
		return nil, err
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	config = &Config{
		ServerPort: serverPort,
	}

	return config, nil
}
