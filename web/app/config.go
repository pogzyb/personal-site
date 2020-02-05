package app

import (
	"fmt"
	"os"
)

type Config struct {
	Environment	string
	SecretKey 	[]byte
}

func InitConfig() (*Config, error) {
	config := &Config{
		Environment: os.Getenv("APP_ENV"),
		SecretKey: []byte(os.Getenv("APP_SECRET")),
	}
	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("need to set $APP_SECRET environment variable")
	}
	return config, nil
}