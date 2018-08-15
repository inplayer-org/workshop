package application

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"encoding/json"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))

}

func (a *App) initializeRoutes() {
/*	a.Router.HandleFunc("/TVOJA", a.GetEmployers).Methods("GET")
	a.Router.HandleFunc("/TVOJA", a.CreateEmployers).Methods("POST")
	a.Router.HandleFunc("/TVOJA/{id:[0-9]+}", a.GetEmployer).Methods("GET")
	a.Router.HandleFunc("/TVOJA/{id:[0-9]+}", a.UpdateEmployer).Methods("PUT")
	a.Router.HandleFunc("/TVOJA/{id:[0-9]+}", a.DeleteEmployer).Methods("DELETE") */
}

func (a *App) GetEmployers(w http.ResponseWriter, r *http.Request) {


	Employers, err := GetEmployers(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, Employers)
}

func (a *App) CreateEmployers(w http.ResponseWriter, r *http.Request) {
//	var e Employers (CHECK ZA TVOJTA STRUCT) <<<<<<<<<<<<<<<<<<<<<<<<
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := u.CreateEmployers(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, e)
}

func (a *App) GetEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employer ID")
		return
	}

//	e := employer{ID: id} (CHECK ZA TVOJTA STRUCT)
	if err := e.GetEmployer(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Employer not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, e)
}
func (a *App) UpdateEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

//	var e employer(CHECK ZA TVOJTA STRUCT) <<<<<<<<<<<<<<<<<<<<<<<<<<<<
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	e.ID = id

	if err := e.UpdateEmployer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, e)
}

func (a *App) DeleteEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

//	e := employer{ID: id} (CHECK ZA TVOJTA STRUCT) <<<<<<<<<<<<<<<<<<<<<<<<<<<
	if err := e.DeleteEmployer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
response, _ := json.Marshal(payload)

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(code)
w.Write(response)
}
