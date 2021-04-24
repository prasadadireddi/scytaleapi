package database

import (
	"github.com/jinzhu/gorm"
	"github.com/prasadadireddi/scytaleapi/config"
)

// Connect to the DATABASE
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
