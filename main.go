package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/chatex-com/process-manager"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/MarySmirnova/pereval/internal"
	"github.com/MarySmirnova/pereval/internal/config"
)

var cfg config.Application

func init() {
	_ = godotenv.Load(".env")
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	process.SetLogger(NewPMLogger())
}

func main() {
	app, err := internal.NewApplication(cfg)
	if err != nil {
		panic(err)
	}

	app.Run()
}
