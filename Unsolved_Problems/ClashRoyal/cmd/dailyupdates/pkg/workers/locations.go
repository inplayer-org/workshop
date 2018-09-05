package workers

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

type LocationWorker struct{
	Loc structures.Locationsinfo
}

func NewLocationWorker(loc structures.Locationsinfo)Worker{

	return &LocationWorker{loc}

}

func(locationWorker *LocationWorker) FinishUpdate(db *sql.DB)string{

	if locationWorker.Loc.IsCountry {
		playerTags, err := update.GetPlayerTagsPerLocation(locationWorker.Loc.ID)

		if err!=nil{
			log.Println(err)
		}

		allErrors := update.Players(db, parser.ToUrlTags(playerTags.GetTags()), locationWorker.Loc.ID)

		for _, err := range allErrors {
			log.Println(err)
		}
	}
	return locationWorker.Loc.Name
}

