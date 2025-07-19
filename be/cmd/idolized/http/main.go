package main

import (
	"context"
	"github.com/RandySteven/Idolized/apps"
	"github.com/RandySteven/Idolized/configs"
	"log"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	configPath, err := configs.ParseFlags()
	if err != nil {
		log.Println(`failed to read config path`, err)
		return
	}

	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Println(`failed to execute config`, err)
		return
	}

	_, err = apps.NewApps(config)
	if err != nil {
		log.Println(`failed to start app`)
		return
	}
}
