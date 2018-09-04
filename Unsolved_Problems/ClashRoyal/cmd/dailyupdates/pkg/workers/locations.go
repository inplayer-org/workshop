package workers

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func Locations(db *sql.DB,locationInfoChan <- chan structures.Locationsinfo,done chan <- string){


	for location :=  range locationInfoChan {
		if location.IsCountry {
			playerTags, err := update.GetPlayerTagsPerLocation(location.ID)

			if err!=nil{
				log.Println(err)
			}

			allErrors := update.Players(db, parser.ToUrlTags(playerTags.GetTags()), location.ID)

			for _, err := range allErrors {
				log.Println(err)
			}
		}
		done <- location.Name
	}
}