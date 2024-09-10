package tasks

import (
	"secure-api/models"
)

func (service *Service) CreateTask(task *models.Tasks) error {
	return service.DB.
		Create(task).
		Error
}
