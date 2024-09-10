package tasks

import (
	"secure-api/models"
)

func (service *Service) UpdateTask(task *models.Tasks) error {
	return service.DB.
		Save(task).
		Error
}
