package workers

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
)

type Worker interface {
	FinishUpdate(db *sql.DB,client _interface.ClientInterface)string
}

func StartWorker(db *sql.DB,InfoChan <- chan Worker,done chan <- string){


	client := _interface.NewClient()

	for information :=  range InfoChan {

		done<-information.FinishUpdate(db,client)

	}
}