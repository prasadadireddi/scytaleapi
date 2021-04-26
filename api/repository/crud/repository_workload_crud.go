package crud

import (
"fmt"
"sort"
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

func (r *RepositoryWorkloadsCRUD) FindAllSorted() ([]models.Workload, error) {
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
		sort.Sort(models.SpiffeSorter(workloads))
		ch <- true
	}(done)
	if channels.OK(done) {
		return workloads, nil
	}
	return nil, err
}

func (r *RepositoryWorkloadsCRUD) FindBySelector(selector string) ([]models.Workload, error) {
	var err error
	workloads := []models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Workload{}).Where("? = ANY(selectors)", selector).Find(&workloads).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return workloads, nil
	}
	return []models.Workload{}, err
}

func (r *RepositoryWorkloadsCRUD) UpdateWorkload(spiffeid string, workload models.Workload) (models.Workload, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Workload{}).Where("spiffe_id = ?", spiffeid).Update(
			map[string]interface{}{
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

func (r *RepositoryWorkloadsCRUD) UpdateSelector(spiffeid string, selector string) (models.Workload, error) {
	var rs *gorm.DB
	var err error
	workload := models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Workload{}).Where("spiffe_id = ?", spiffeid).Find(&workload).Error
		workload.Selectors = append(workload.Selectors, selector)
		rs = r.db.Debug().Model(&models.Workload{}).Where("spiffe_id = ?", spiffeid).Update(
			map[string]interface{}{
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

func (r *RepositoryWorkloadsCRUD) DeleteWorkload(spiffeid string) (models.Workload, error) {
	var rs *gorm.DB
	workload := models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Workload{}).Where("spiffe_id = ?", spiffeid).Take(&models.Workload{}).Delete(&models.Workload{})
		ch <- true
	}(done)

	if channels.OK(done) {
		return workload, nil
	}
	return models.Workload{}, rs.Error
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func (r *RepositoryWorkloadsCRUD) DeleteSelector(spiffeid string, selector string) (models.Workload, error) {
	var rs *gorm.DB
	var err error
	workload := models.Workload{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Workload{}).Where("spiffe_id = ?", spiffeid).Find(&workload).Error
		// fmt.Println(workload.Selectors)
		index := 0
	    for _, i := range workload.Selectors {
	        if i != selector {
	            fmt.Println(i)
	            index++
	        } else {
	           workload.Selectors = RemoveIndex(workload.Selectors, index)
	        }
	    }

	    rs = r.db.Debug().Model(&models.Workload{}).Where("spiffe_id = ?", spiffeid).Update(
			map[string]interface{}{
				"selectors":  workload.Selectors,
			},
		)
	    // fmt.Println(workload.Selectors)
		ch <- true
	}(done)

	if channels.OK(done) {
		return workload, nil
	}
	return models.Workload{}, rs.Error
}