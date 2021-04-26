package auto

import (
"github.com/prasadadireddi/scytaleapi/api/models"
)

var workload = []models.Workload{
	models.Workload{SpiffeID: "test", Selectors: []string{"Python:Java", "Java:Python"}},
	models.Workload{SpiffeID: "test1", Selectors: []string{"Python1:Java1", "Java1:Python1"}},
}

