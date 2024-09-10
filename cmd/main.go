package main

import (
	"log"
	"secure-api/config"
	"secure-api/server"
	"secure-api/server/routes"
)

// @Title Secure REST API
// @Version 1.0
// @description This is a Secure REST API written in GO.
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

	routes.ConfigureRoutes(server)
	server.Listen()
}
