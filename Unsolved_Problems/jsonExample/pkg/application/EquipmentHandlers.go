package application

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/employerinfo"
	"database/sql"
	"encoding/json"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
)

func (a *App) GetEquipment(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "ID shpuld be int")
		return
	}

	e := employerinfo.Equipment{EmployerID: id}
	if err := e.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "equipment not found")

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, e)
}

func (a *App) UpdateEquipment(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "ID should be int")
		return
	}

	var e employerinfo.Equipment
	e.EmployerID = id
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()
	// e.EmployerID = id

	if err := e.Update(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "equipment doesnt exist create it")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, e)
}

func (a *App) DeleteEquipment(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

	e := employerinfo.Equipment{EmployerID: id}
	if err := e.Delete(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "equipment not found or already deleted")

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) GetEquipments(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	equipments, err := employerinfo.GetAllEquipments(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "no equipment found")

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, equipments)
}