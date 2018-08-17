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
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/employerinfo"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
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
	a.Router.HandleFunc("/employers", a.GetEmployers).Methods("GET")
	a.Router.HandleFunc("/employer", a.CreateEmployers).Methods("POST")
	a.Router.HandleFunc("/employer/{id:[0-9]+}", a.GetEmployer).Methods("GET")
	a.Router.HandleFunc("/employer/{id:[0-9]+}", a.UpdateEmployer).Methods("PUT")
	a.Router.HandleFunc("/employer/{id:[0-9]+}", a.DeleteEmployer).Methods("DELETE")

	a.Router.HandleFunc("/equipments", a.GetEquipments).Methods("GET")
	a.Router.HandleFunc("/equipment/{id:[0-9]+}", a.GetEquipment).Methods("GET")
	a.Router.HandleFunc("/equipment/{id:[0-9]+}", a.UpdateEquipment).Methods("PUT")
	//a.Router.HandleFunc("/equipment", a.CreateEquipment).Methods("POST")
	a.Router.HandleFunc("/equipment/{id:[0-9]+}", a.DeleteEquipment).Methods("DELETE")

	a.Router.HandleFunc("/positions", a.GetPositions).Methods("GET")
	a.Router.HandleFunc("/position/{name:[a-z]+}", a.GetPosition).Methods("GET")
	a.Router.HandleFunc("/position/{name:[a-z]+}", a.UpdatePosition).Methods("PUT")
	a.Router.HandleFunc("/position", a.CreatePosition).Methods("POST")
	a.Router.HandleFunc("/position/{name:[a-z]+}", a.DeletePosition).Methods("DELETE")

	a.Router.HandleFunc("/contracts", a.GetContracts).Methods("GET")
	a.Router.HandleFunc("/contract/{id:[0-9]+}", a.GetContract).Methods("GET")
	a.Router.HandleFunc("/contract/{id:[0-9]+}", a.UpdateContract).Methods("PUT")
	a.Router.HandleFunc("/contract", a.CreateContract).Methods("POST")
	a.Router.HandleFunc("/contract/{id:[0-9]+}", a.DeleteContract).Methods("DELETE")
}
func (a *App) GetEmployers(w http.ResponseWriter, r *http.Request) {


	employers, err := employerinfo.GetAllEmployers(a.DB)
	e:=a.DB.Ping()

	if e!=nil{
		respondWithError(w, http.StatusNotFound, "you have wron username,password or database name")
		return
	}

	if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, employers)
}

func (a *App) CreateEmployers(w http.ResponseWriter, r *http.Request) {
	var e employerinfo.EmployerInfo

	//id dobiva vrednos tuka

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := e.Create(a.DB); err != nil {

		switch err {
		case errorhandle.Err:
			respondWithError(w, http.StatusConflict, err.Error())
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Position not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
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

	e := employerinfo.EmployerInfo{ID: id}
	if err := e.Get(a.DB); err != nil {

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

	var e employerinfo.EmployerInfo
	e.ID = id
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()


	if err := e.Update(a.DB); err != nil {

			switch err {
			case errorhandle.Err:
				respondWithError(w, http.StatusConflict, err.Error())
			case sql.ErrNoRows:
				respondWithError(w, http.StatusNotFound, "Position not found")
			default:
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		respondWithJSON(w, http.StatusCreated, e)
	}

func (a *App) DeleteEmployer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

	e := employerinfo.EmployerInfo{ID: id}
	if err := e.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) GetContracts(w http.ResponseWriter, r *http.Request) {


	cs, err := employerinfo.GetAllContracts(a.DB)
	fmt.Println(cs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}



	respondWithJSON(w, http.StatusOK, cs)
}
func (a *App) GetContract(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract ID")
		return
	}

	c := employerinfo.Contract{ContractNumber: id}
	if err := c.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Contract not found")

		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}
func (a *App) UpdateContract(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract ID")
		return
	}

	var c employerinfo.Contract
	c.EmployerID = id
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	// c.EmployerID = id

	if err := c.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}
func (a *App) CreateContract(w http.ResponseWriter, r *http.Request) {
	var c employerinfo.Contract
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {



		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()


	if err := c.Create(a.DB); err != nil {

		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}


	respondWithJSON(w, http.StatusCreated, c)
}
func (a *App) DeleteContract(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract ID")
		return
	}

	c := employerinfo.Contract{ContractNumber: id}
	if err := c.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


func (a *App) GetPositions(w http.ResponseWriter, r *http.Request) {


	positions, err := employerinfo.GetAllPositions(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, positions)
}

func (a *App) GetPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, err := vars["name"]

	if err != true {
		respondWithError(w, http.StatusBadRequest, "Invalid Position ")
		return
	}

	p := employerinfo.Position{Name: name}
	//fmt.Println(p.Name)
	if err := p.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "position not found")

		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}
func (a *App) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, err := vars["name"]
	if err != true {
	respondWithError(w, http.StatusBadRequest, "Invalid Position ID")
		return
}

	var p employerinfo.Position
	p.Name = name
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	// p.EmployerID = id

	if err := p.Update(a.DB); err != nil {

		switch err {
		case errorhandle.Err:
			respondWithError(w, http.StatusConflict, err.Error())

		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}
func (a *App) CreatePosition(w http.ResponseWriter, r *http.Request) {
	var p employerinfo.Position
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.Create(a.DB); err != nil {

		switch err {
		case errorhandle.Err:
			respondWithError(w, http.StatusConflict, err.Error())

		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) DeletePosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, err := vars["name"]
	if err != true {
		respondWithError(w, http.StatusBadRequest, "Invalid Position ID")
		return
	}

	p := employerinfo.Position{Name: name}
	if err := p.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}



func (a *App) GetEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Equipment ID")
		return
	}

	    e := employerinfo.Equipment{EmployerID: id}
	if err := e.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "equipment not found")

		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, e)
}
func (a *App) UpdateEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Equipment ID")
		return
	}

	 var e employerinfo.Equipment
	e.EmployerID = id
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	// e.EmployerID = id

	if err := e.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, e)
}

func (a *App) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Employer ID")
		return
	}

	 e := employerinfo.Equipment{EmployerID: id}
	if err := e.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) GetEquipments(w http.ResponseWriter, r *http.Request) {


	equipments, err := employerinfo.GetAllEquipments(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, equipments)
}

/*func (a *App) CreateEquipment(w http.ResponseWriter, r *http.Request) {
	var e employerinfo.Equipment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := e.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, e)
}*/




func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

