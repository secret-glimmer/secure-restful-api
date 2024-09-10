package routes

import (
	"secure-api/server"
	"secure-api/server/handlers"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

func ConfigureRoutes(server *server.Server) {
	server.App.Get("/docs/*", swagger.WrapHandler)

	groupTasks := server.App.Group("/tasks")
	GroupTasks(server, groupTasks)
}

func GroupTasks(server *server.Server, group fiber.Router) {
	handler := handlers.NewHandlerTasks(server)

	group.Post("", handler.CreateTask)
	group.Get("", handler.ListTasks)
	group.Get("/:id", handler.GetTask)
	group.Put("/:id", handler.UpdateTask)
	group.Delete("/:id", handler.DeleteTask)
}
