package application

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/employerinfo"
	"database/sql"
	"encoding/json"
)

func (a *App) GetPositions(w http.ResponseWriter, r *http.Request) {
	errorhandle.CheckDB(a.DB,w)

	positions, err := employerinfo.GetAllPositions(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "no positions found")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, positions)
}

func (a *App) GetPosition(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	name, err := vars["name"]
	if err != true {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "enter name for the position you want to get info")
		return
	}

	p := employerinfo.Position{Name: name}
	if err := p.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "position not found")

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) UpdatePosition(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	name, err := vars["name"]
	if err != true {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "enter name for the position you want to get info")
		return
	}

	p:=employerinfo.Position{Name:name}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	if err := p.Update(a.DB); err != nil {
		switch err {
		case errorhandle.Err:
			errorhandle.RespondWithError(w, http.StatusConflict, err.Error())

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusCreated, p)
}

func (a *App) CreatePosition(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	var p employerinfo.Position
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	if err := p.Create(a.DB); err != nil {
		switch err {
		case errorhandle.Err:
			errorhandle.RespondWithError(w, http.StatusConflict, err.Error())

		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusCreated, p)
}

func (a *App) DeletePosition(w http.ResponseWriter, r *http.Request) {

	errorhandle.CheckDB(a.DB,w)

	vars := mux.Vars(r)
	name, err := vars["name"]
	if err != true {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid Position name")
		return
	}

	p := employerinfo.Position{Name: name}
	if err := p.Delete(a.DB); err != nil {
		switch err {
			case sql.ErrNoRows:
		errorhandle.RespondWithError(w, http.StatusNotFound, "position does not exist or it is already deleted")
			default:
		errorhandle.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


