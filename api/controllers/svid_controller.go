
package controllers

import (
	 _ "encoding/json"
	 "net/http"
	 "github.com/gorilla/mux"
	 "github.com/prasadadireddi/scytaleapi/api/repository"
	 "github.com/prasadadireddi/scytaleapi/api/repository/crud"
	 "github.com/prasadadireddi/scytaleapi/api/responses"
)

func ValidateSpiffeID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spiffeid := vars["spiffeid"]
	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// defer db.Close()

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
