package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	HTTP struct {
		Port          string `env:"HTTP_PORT"`
		MainServerURL string `env:"MAIN_SERVER_URL"`
	}
	Database struct {
		DSN string `env:"DSN"`
	}
}

func InitializeConfig() (*Config, error) {
	conf := &Config{}

	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}

	err = dotenv.NewDecoder(file).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
