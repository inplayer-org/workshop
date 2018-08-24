package HandlersFunc


func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/clans/{name:[a-z]+}", a.GetClanByName).Methods("GET")
	a.Router.HandleFunc("/players/{name:[a-z]+}", a.GetPlayerByName).Methods("GET")

}
