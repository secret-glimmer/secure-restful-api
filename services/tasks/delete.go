package tasks

import (
	"secure-api/models"

	"github.com/google/uuid"
)

func (service *Service) DeleteTask(id uuid.UUID, task *models.Task) error {
	return service.DB.
		Delete(task, id).
		Error
}
