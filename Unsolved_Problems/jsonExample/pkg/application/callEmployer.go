package application

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/employerinfo"
	"encoding/json"
	"fmt"
)

func (a *App) GetEmployers(w http.ResponseWriter, r *http.Request) {


	employers, err := employerinfo.GetAllEmployers(a.DB)
	e:=a.DB.Ping()

	if e!=nil{
		errorhandle.RespondWithError(w, http.StatusNotFound, "you have wrong username,password or database name")
		return
	}

	if err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, employers)
}

func (a *App) CreateEmployers(w http.ResponseWriter, r *http.Request) {
	var e employerinfo.EmployerInfo


	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		fmt.Println(err)
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := e.Create(a.DB); err != nil {

		switch err {
		case errorhandle.Err:
			errorhandle.RespondWithError(w, http.StatusConflict, err.Error())
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "Position not found")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusCreated, e)
}

func (a *App) GetEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid employer ID")
		return
	}

	e := employerinfo.EmployerInfo{ID: id}
	if err := e.Get(a.DB); err != nil {

		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "Employer not found")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, e)
}
func (a *App) UpdateEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

	var e employerinfo.EmployerInfo
	e.ID = id
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()


	if err := e.Update(a.DB); err != nil {

		switch err {
		case errorhandle.Err:
			errorhandle.RespondWithError(w, http.StatusConflict, err.Error())
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "Position not found")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusCreated, e)
}

func (a *App) DeleteEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

	e := employerinfo.EmployerInfo{ID: id}
	if err := e.Delete(a.DB); err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
