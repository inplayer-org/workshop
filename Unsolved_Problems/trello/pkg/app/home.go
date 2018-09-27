package app

import (
	"net/http"
	"fmt"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"home page")
}