package repository

import (
"github.com/prasadadireddi/scytaleapi/api/models"
)

// PostRepository is the interface Post CRUD
type WorkloadRepository interface {
	Save(models.Workload) (models.Workload, error)
	FindAll() ([]models.Workload, error)
	FindBySelector(string) ([]models.Workload, error)
	UpdateWorkload(string, models.Workload) (models.Workload, error)
	UpdateSelector(string, string) (models.Workload, error)
	DeleteWorkload(string) (models.Workload, error)
	DeleteSelector(string, string) (models.Workload, error)
}
