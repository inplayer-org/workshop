package errorhandle

import (
	"database/sql"
	"net/http"
)

func CheckDB(db *sql.DB,w http.ResponseWriter) {
	er := db.Ping()

	if er != nil {
		RespondWithError(w, http.StatusNotFound, "you have wrong username,password or database name")
		return
	}
}