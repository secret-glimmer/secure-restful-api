package tasks

import (
	"secure-api/models"
)

func (service *Service) UpdateTask(task *models.Task) error {
	return service.DB.
		Save(task).
		Error
}
