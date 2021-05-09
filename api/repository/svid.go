package repository

import (
_ "github.com/prasadadireddi/scytaleapi/api/models"
)

// PostRepository is the interface Post CRUD
type SvidRepository interface {
	ValidateSpiffeID(string) (int, error)
}
