package workers

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)

type LocationWorker struct{
	Loc structures.Locationsinfo
}

func NewLocationWorker(loc structures.Locationsinfo)Worker{

	return &LocationWorker{loc}

}

func(locationWorker *LocationWorker) FinishUpdate(db *sql.DB)string{

	if locationWorker.Loc.IsCountry {

		client :=_interface.NewClient()

		players, err := client.GetPlayerTagsFromLocation(locationWorker.Loc.ID)

		if err!=nil{
			log.Println(err)
		}

		//log.Println(playerTags.GetTags())
		playerTags := players.GetTags()

		for _,nextPlayerTag := range playerTags{

			currentPlayer,err := client.GetRequestForPlayer(parser.ToUrlTag(nextPlayerTag))

			if err!=nil{
				log.Println(err)
			}else{
				queries.UpdatePlayer(db,currentPlayer,0)
			}
		}

	}
	return locationWorker.Loc.Name
}

