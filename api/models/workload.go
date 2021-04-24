package models

import (
"errors"
"github.com/lib/pq"

)

// Post model
type Workload struct {
	SpiffeID             string         `gorm:"primary_id json:"spiffeid"`
	Selectors            pq.StringArray `gorm:"type:varchar(60)[]" json:"selectors,omitempty"`

}


// Validate validates the inputs
func (w *Workload) Validate() error {
	if w.SpiffeID == "" {
		return errors.New("SpiffeID is required")
	}

	//if w.Selectors == "" {
	//	return errors.New("SpiffeID is required")
	//}


	return nil
}
