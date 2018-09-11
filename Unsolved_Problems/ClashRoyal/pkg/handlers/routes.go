package handlers


func (a *App) initializeRoutes() {



	a.Router.HandleFunc("/clans/{name}", a.GetClanByName).Methods("GET")
	a.Router.HandleFunc("/clans/{name}/{tag}", a.GetClanByTag).Methods("GET")
	a.Router.HandleFunc("/clans/{name}/{tag}/update",a.UpdateClan).Methods("GET")

	a.Router.HandleFunc("/players/{name}", a.GetPlayerByName).Methods("GET")
	a.Router.HandleFunc("/players/{name}/{tag}", a.GetPlayerByTag).Methods("GET")

	a.Router.HandleFunc("/players/{name}/{tag}/update",a.UpdatePlayer).Methods("GET")
	a.Router.HandleFunc("/",a.Home).Methods("GET")

	a.Router.HandleFunc("/locations",a.GetLocations).Methods("GET")
	a.Router.HandleFunc("/locations/{name}", a.GetLocationByName).Methods("GET")

	a.Router.HandleFunc("/search",a.Search).Methods("GET")
}