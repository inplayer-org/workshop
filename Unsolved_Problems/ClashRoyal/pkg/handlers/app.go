package handlers

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
)

//App
type App struct {
	Router *mux.Router
	DB     *sql.DB
	Client _interface.ClientInterface
}

//Initialize creates var App
func (a *App) Initialize(db *sql.DB, router *mux.Router) {

	a.DB = db

	a.Router = router

	a.Client = _interface.NewClient()

	a.Client = _interface.NewClient()

	a.initializeRoutes()

	a.Run(":3303")

}

//Run starts server
func (a *App) Run(addr string) {

	log.Fatal(http.ListenAndServe(addr, a.Router))

}
