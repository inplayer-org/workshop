package application

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/employerinfo"
	"fmt"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
)

func (a *App) GetContracts(w http.ResponseWriter, r *http.Request) {
	errorhandle.CheckDB(a.DB,w)

	cs, err := employerinfo.GetAllContracts(a.DB)
	fmt.Println(cs)
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}



	errorhandle.RespondWithJSON(w, http.StatusOK, cs)
}
func (a *App) GetContract(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Contract ID")
		return
	}

	c := employerinfo.Contract{ContractNumber: id}
	if err := c.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "Contract not found")

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, c)
}
func (a *App) UpdateContract(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Contract ID")
		return
	}

	var c employerinfo.Contract
	c.EmployerID = id
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	// c.EmployerID = id

	if err := c.Update(a.DB); err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, c)
}
func (a *App) CreateContract(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	var c employerinfo.Contract
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {



		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()


	if err := c.Create(a.DB); err != nil {

		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}


	errorhandle.RespondWithJSON(w, http.StatusCreated, c)
}
func (a *App) DeleteContract(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Contract ID")
		return
	}

	c := employerinfo.Contract{ContractNumber: id}
	if err := c.Delete(a.DB); err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

