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

func(locationWorker *LocationWorker) FinishUpdate(db *sql.DB,client _interface.ClientInterface)string{

	//Checks if the location is a country and skips the whole operation if it's not
	// (there aren't player rankings available for locations that aren't a country in the clash royale api)
	if locationWorker.Loc.IsCountry {

		//Get PlayerTags structure of all players from the location
		players, err := client.GetPlayerTagsFromLocation(locationWorker.Loc.ID)

		if err!=nil{
			log.Println(err)
		}

		//Converting PlayerTags structure into string[]
		playerTags := players.GetTags()

		//Requesting and updating information for every player
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

