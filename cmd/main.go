package main

import (
	"log"
	"secure-api/config"
	"secure-api/server"
)

func main() {
	config := config.NewConfig()
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatalln(err)
	}

	server.Listen()
}
