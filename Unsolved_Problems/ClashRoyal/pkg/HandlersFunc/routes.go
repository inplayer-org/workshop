package HandlersFunc


func (a *App) initializeRoutes() {


	a.Router.HandleFunc("/clans", a.GetClans).Methods("GET")
	a.Router.HandleFunc("/clans/{name:[a-z]+}", a.GetClanByName).Methods("GET")

//	a.Router.HandleFunc("/players",a.GetPlayers).Methods("GET")
	a.Router.HandleFunc("/players/{name:[a-z]+}", a.GetPlayerByName).Methods("GET")

	a.Router.HandleFunc("/",a.Home).Methods("GET")
	a.Router.HandleFunc("/locations/{name:[a-z]+}", a.GetLocationByName).Methods("GET")
}