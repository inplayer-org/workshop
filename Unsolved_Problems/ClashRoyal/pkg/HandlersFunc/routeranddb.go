package HandlersFunc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"net/http"
)

//App
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize creates var App
func (a *App) Initialize(db *sql.DB,router *mux.Router) {

	a.DB=db

	a.Router=router

	a.initializeRoutes()
	a.Run(":3303")
}

//Run starts server
func (a *App) Run(addr string) {

	log.Fatal(http.ListenAndServe(addr, a.Router))

}