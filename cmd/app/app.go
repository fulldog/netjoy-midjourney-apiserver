package app

import (
	"github.com/hongliang5316/midjourney-apiserver/internal/application"
	"github.com/hongliang5316/midjourney-apiserver/internal/config"
	"log"
)

func Run(cfg *config.Config) {
	app := application.NewCfg(cfg)
	if err := app.Run(); err != nil {
		log.Fatalf("Call app.Run failed, err: %+v", err)
	}
}
