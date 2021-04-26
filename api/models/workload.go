package models

import (
"errors"
"github.com/lib/pq"

)

type SpiffeSorter []Workload

func (a SpiffeSorter) Len() int           { return len(a) }
func (a SpiffeSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SpiffeSorter) Less(i, j int) bool { return a[i].SpiffeID < a[j].SpiffeID }

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
