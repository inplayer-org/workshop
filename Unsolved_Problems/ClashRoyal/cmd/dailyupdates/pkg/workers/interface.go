package workers

import (
	"database/sql"
)

type Worker interface {
	FinishUpdate(db *sql.DB)string
}

func StartWorker(db *sql.DB,InfoChan <- chan Worker,done chan <- string){


	for information :=  range InfoChan {

		done<-information.FinishUpdate(db)

	}
}