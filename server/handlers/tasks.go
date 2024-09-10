package handlers

import (
	"encoding/json"
	"secure-api/models"
	"secure-api/requests"
	"secure-api/responses"
	s "secure-api/server"
	"secure-api/services/tasks"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HandlerTasks struct {
	Server *s.Server
}

func NewHandlerTasks(server *s.Server) *HandlerTasks {
	return &HandlerTasks{
		Server: server,
	}
}

// Refresh godoc
// @Summary Create task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param params body requests.RequestTask true "Task Request"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /tasks [post]
func (h *HandlerTasks) CreateTask(c *fiber.Ctx) error {
	request := requests.RequestTask{}

	err := json.Unmarshal(c.Body(), &request)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusBadRequest, "Task request data is invalid.")
	}

	now := time.Now()

	task := models.Tasks{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		DeletedAt:   nil,
	}

	service := tasks.NewService(h.Server.DB)
	err = service.CreateTask(&task)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create task.")
	}

	return responses.MessageResponse(c, fiber.StatusOK, "Task is successfully created.")
}
