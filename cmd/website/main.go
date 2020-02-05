package main

import (
	"github.com/pogzyb/website/web/app"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	log.Info("Starting Application!")
	application := app.New()
	application.Start()
}
