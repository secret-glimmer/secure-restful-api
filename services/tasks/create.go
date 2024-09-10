package tasks

import (
	"secure-api/models"
)

func (service *Service) CreateTask(task *models.Task) error {
	return service.DB.
		Create(task).
		Error
}
