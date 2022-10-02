package cmd

import (
	"intermediate_server/app"
	"intermediate_server/config"
	"log"
)

func Execute() {
	conf, err := config.InitializeConfig()
	if err != nil {
		panic(err)
	}

	server, err := app.Init(conf)
	if err != nil {
		panic(err)
	}

	log.Println(server.Start(conf))
}
