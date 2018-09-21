package handlers

import (
	"net/http"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("test homepage"))

}