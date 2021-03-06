package workers

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"strconv"
)

type LocationWorker struct{
	Loc structures.Locationsinfo
}

func NewLocationWorker(loc structures.Locationsinfo)Worker{

	return &LocationWorker{loc}

}

//Worker for sending request for all player tags from a specific location (present in the database) to the clash royale api and writing the response to database
func(locationWorker *LocationWorker) FinishUpdate(db *sql.DB,client _interface.ClientInterface)string{

	//Checks if the location is a country and skips the whole operation if it's not
	// (there aren't player rankings available for locations that aren't a country in the clash royale api)
	if locationWorker.Loc.IsCountry {

		//Get PlayerTags structure of all players from the location
		players, err := client.GetPlayerTagsFromLocation(locationWorker.Loc.ID)

		if err!=nil{
			errors.Database(err)
			return "Failed to get data for location " + locationWorker.Loc.Name + " with ID " + strconv.Itoa(locationWorker.Loc.ID)
		}


		//Converting PlayerTags structure into string[]
		playerTags := players.GetTags()

		//Requesting and updating information for every player
		for _,nextPlayerTag := range playerTags{

			currentPlayer,err := client.GetRequestForPlayer(nextPlayerTag)

			if err!=nil{
				errors.Database(err)
			}else{
				queries.UpdatePlayer(db,currentPlayer,locationWorker.Loc.ID)
			}
		}

	}
	return locationWorker.Loc.Name
}

