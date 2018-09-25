package app

func(a *App) initializeRoutes(){
	a.Router.HandleFunc("/",a.Home).Methods("GET")
}
