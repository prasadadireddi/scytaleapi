
package controllers

import (
	 "encoding/json"
	 "net/http"
	 _ "github.com/gorilla/mux"
         "github.com/prasadadireddi/scytaleapi/api/models"
	 "github.com/prasadadireddi/scytaleapi/api/repository"
	 "github.com/prasadadireddi/scytaleapi/api/repository/crud"
	 "github.com/prasadadireddi/scytaleapi/api/responses"
)

func ValidateSpiffeID(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//spiffeid := vars["spiffeid"]
        workload := models.Workload{}
	err := json.NewDecoder(r.Body).Decode(&workload)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
        spiffeid := workload.SpiffeID

	repo := crud.NewRepositorySvidCRUD()

	func(SvidRepository repository.SvidRepository) {
		ret, err := SvidRepository.ValidateSpiffeID(spiffeid)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, ret)
	}(repo)
}
