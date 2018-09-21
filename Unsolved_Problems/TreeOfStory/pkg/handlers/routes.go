package handlers

func (a *App) initializeRoutes() {



	a.Router.HandleFunc("/", a.Home).Methods("GET")

}
