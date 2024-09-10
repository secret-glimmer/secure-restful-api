package tasks

import (
	"secure-api/models"

	"github.com/google/uuid"
)

func (service *Service) ReadTaskByID(id uuid.UUID, task *models.Tasks) error {
	return service.DB.
		Find(task, id).
		Error
}

func (service *Service) ReadAllTask(tasks *[]models.Tasks) error {
	return service.DB.
		Find(tasks).
		Error
}
