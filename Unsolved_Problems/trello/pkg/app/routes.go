package app

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/search", a.Search).Methods("GET")

	a.Router.HandleFunc("/members/{tag}", a.GetMemberByUsername).Methods("GET")


	a.Router.HandleFunc("/board",a.GetBoards).Methods("GET")
	a.Router.HandleFunc("/board/{id}",a.GetBoardByID).Methods("GET")

}
