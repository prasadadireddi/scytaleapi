package crud

import (

"github.com/jinzhu/gorm"
"github.com/prasadadireddi/scytaleapi/api/models"
"github.com/prasadadireddi/scytaleapi/api/utils/channels"
)

// RepositoryPostsCRUD is the struct for the Post CRUD
type RepositoryWorkloadsCRUD struct {
	db *gorm.DB
}

// NewRepositoryPostsCRUD returns a new repository with DB connection
func NewRepositoryWorkloadsCRUD(db *gorm.DB) *RepositoryWorkloadsCRUD {
	return &RepositoryWorkloadsCRUD{db}
}

// Save returns a new post created or an error
func (r *RepositoryWorkloadsCRUD) Save(workload models.Workload) (models.Workload, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Workload{}).Create(&workload).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return workload, nil
	}
	return models.Workload{}, err
}

func (r *RepositoryWorkloadsCRUD) FindAll() ([]models.Workload, error) {
	var err error
	workloads := []models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Find(&workloads).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)
	if channels.OK(done) {
		return workloads, nil
	}
	return nil, err
}


func (r *RepositoryWorkloadsCRUD) Update(spiffeid string, workload models.Workload) (models.Workload, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Workload{}).Where("spiffeid = ?", spiffeid).Take(&models.Workload{}).UpdateColumns(
			map[string]interface{}{
				"spiffeid": workload.SpiffeID,
				"selectors":  workload.Selectors,
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		return workload, nil
	}
	return models.Workload{}, rs.Error
}


func (r *RepositoryWorkloadsCRUD) Delete(spiffeid string) (models.Workload, error) {
	var rs *gorm.DB
	workload := models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Workload{}).Where("spiffeid = ?", spiffeid).Take(&models.Workload{}).Delete(&models.Workload{})
		ch <- true
	}(done)

	if channels.OK(done) {
		return workload, nil
	}
	return models.Workload{}, rs.Error
}