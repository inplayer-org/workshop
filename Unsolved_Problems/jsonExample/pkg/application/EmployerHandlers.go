package application

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/employerinfo"
	"encoding/json"
)

func (a *App) GetEmployers(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	employers, err := employerinfo.GetAllEmployers(a.DB)

	if err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, employers)
}

func (a *App) CreateEmployers(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	var e employerinfo.EmployerInfo

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	if err := e.Create(a.DB); err != nil {

		switch err {
		case errorhandle.Err:
			errorhandle.RespondWithError(w, http.StatusConflict, err.Error())
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "Position not found")
		case errorhandle.ErrBadFormat:
			errorhandle.RespondWithError(w, http.StatusBadRequest, err.Error())
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusCreated, e)
}

func (a *App) GetEmployer(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

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

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "ID should be intteger")
		return
	}

	e:=employerinfo.EmployerInfo{ID:id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
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

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "ID should be integer")
		return
	}

	e := employerinfo.EmployerInfo{ID: id}
	if err := e.Delete(a.DB); err != nil {
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
