package application

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/employers", a.GetEmployers).Methods("GET")
	a.Router.HandleFunc("/employer", a.CreateEmployers).Methods("POST")
	a.Router.HandleFunc("/employer/{id:[0-9]+}", a.GetEmployer).Methods("GET")
	a.Router.HandleFunc("/employer/{id:[0-9]+}", a.UpdateEmployer).Methods("PUT")
	a.Router.HandleFunc("/employer/{id:[0-9]+}", a.DeleteEmployer).Methods("DELETE")

	a.Router.HandleFunc("/equipments", a.GetEquipments).Methods("GET")
	a.Router.HandleFunc("/equipment/{id:[0-9]+}", a.GetEquipment).Methods("GET")
	a.Router.HandleFunc("/equipment/{id:[0-9]+}", a.UpdateEquipment).Methods("PUT")
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



