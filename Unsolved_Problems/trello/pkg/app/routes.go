package app

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/search", a.Search).Methods("GET")
	a.Router.HandleFunc("/members/{tag}", a.GetMemberByUsername).Methods("GET")

}
