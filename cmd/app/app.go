package app

import (
	"encoding/json"
	"github.com/hongliang5316/midjourney-apiserver/internal/application"
	"github.com/hongliang5316/midjourney-apiserver/internal/config"
)

func Run(js []byte) error {
	var cfg config.Config
	err := json.Unmarshal(js, &cfg)
	if err != nil {
		return err
	}
	app := application.NewCfg(cfg)
	if err := app.Run(); err != nil {
		return err
	}
	return nil
}
