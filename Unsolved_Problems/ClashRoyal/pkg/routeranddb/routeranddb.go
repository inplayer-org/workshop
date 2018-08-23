package routeranddb


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(db *sql.DB,router *mux.Router) {

	a.DB=db

	a.Router=router

//	a.initializeRoutes()
}

func (a *App) Run(addr string) {

	log.Fatal(http.ListenAndServe(addr, a.Router))

}

