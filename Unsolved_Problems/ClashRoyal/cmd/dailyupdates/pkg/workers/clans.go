package workers

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func Clans(db *sql.DB,clanInfoChan <- chan structures.Clan,done chan <- string){


	for clan :=  range clanInfoChan {

		playerTags := update.GetTagByClans(parser.ToUrlTag(clan.Tag))


		allErrors := update.Players(db, playerTags, 0)

		for _, err := range allErrors {
			log.Println(err)
		}

		done <- clan.Name
	}
}