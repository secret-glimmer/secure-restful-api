package main

import (
	"log"
	"secure-api/config"
	"secure-api/db"
	"secure-api/models"
)

func main() {
	config := config.NewConfig()
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	db := db.Init(config)
	db.AutoMigrate(&models.Tasks{})
}
