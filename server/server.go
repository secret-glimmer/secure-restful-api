package server

import (
	"secure-api/config"
	"secure-api/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
	DB     *gorm.DB
}

func NewServer(config *config.Config) (*Server, error) {
	app := fiber.New()

	db, err := db.Init(config)
	if err != nil {
		return nil, err
	}

	return &Server{
		App:    app,
		Config: config,
		DB:     db,
	}, nil
}

func (server *Server) Listen() {
	server.App.Listen(":" + server.Config.ServerPort)
}
