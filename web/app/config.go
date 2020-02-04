package app

import (
	"fmt"
	"os"
)

type Config struct {
	SecretKey []byte
}

func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: []byte(os.Getenv("APP_SECRET")),
	}
	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("need to set $APP_SECRET environment variable")
	}
	return config, nil
}