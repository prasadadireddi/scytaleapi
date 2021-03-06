
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	 _ "encoding/json"
	 "github.com/gorilla/mux"
	 "github.com/prasadadireddi/scytaleapi/api/database"
	  "github.com/prasadadireddi/scytaleapi/api/models"
	 "github.com/prasadadireddi/scytaleapi/api/repository"
	 "github.com/prasadadireddi/scytaleapi/api/repository/crud"
	 "github.com/prasadadireddi/scytaleapi/api/responses"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func GetWorkloads(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {
		workloads, err := workloadRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, workloads)
	}(repo)
}


func GetWorkloadsSorted(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {
		workloads, err := workloadRepository.FindAllSorted()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, workloads)
	}(repo)
}

func GetWorkloadsBySelector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	selector := vars["selector"]

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {

		workloads, err := workloadRepository.FindBySelector(selector)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, workloads)
	}(repo)
}


func CreateWorkload(w http.ResponseWriter, r *http.Request) {
	workload := models.Workload{}
	err := json.NewDecoder(r.Body).Decode(&workload)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {

		workload, err := workloadRepository.Save(workload)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, workload.SpiffeID))
		responses.JSON(w, http.StatusCreated, workload)
	}(repo)
}

func UpdateWorkload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["spiffeid"]

	workload := models.Workload{}
	err := json.NewDecoder(r.Body).Decode(&workload)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {

		workload, err := workloadRepository.UpdateWorkload(sid, workload)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, workload.SpiffeID))
		responses.JSON(w, http.StatusCreated, workload)
	}(repo)
}

func UpdateSelector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["spiffeid"]
	selector := vars["selector"]

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {
		
		workload, err := workloadRepository.UpdateSelector(sid, selector)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, workload.SpiffeID))
		responses.JSON(w, http.StatusCreated, workload)
	}(repo)
}

func DeleteWorkload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["spiffeid"]

	workload := models.Workload{}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {
		workload, err = workloadRepository.DeleteWorkload(sid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, workload.SpiffeID))
		responses.JSON(w, http.StatusCreated, workload)
	}(repo)
}

func DeleteSelector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["spiffeid"]
	selector := vars["selector"]

	workload := models.Workload{}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryWorkloadsCRUD(db)

	func(workloadRepository repository.WorkloadRepository) {
		workload, err = workloadRepository.DeleteSelector(sid, selector)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, workload.SpiffeID))
		responses.JSON(w, http.StatusCreated, workload)
	}(repo)
}