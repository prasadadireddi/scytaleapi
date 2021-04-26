package repository

import (
"github.com/prasadadireddi/scytaleapi/api/models"
)

// PostRepository is the interface Post CRUD
type WorkloadRepository interface {
	Save(models.Workload) (models.Workload, error)
	FindAll() ([]models.Workload, error)
	FindBySelector(string) ([]models.Workload, error)
	Update(string, models.Workload) (models.Workload, error)
	Delete(string) (models.Workload, error)
}
