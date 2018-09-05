package workers

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)


type ClanWorker struct{
	Clan structures.Clan
}

func NewClanWorker(c structures.Clan) Worker {
	return &ClanWorker{
		Clan:c,
	}
}

func (clanWorker *ClanWorker) FinishUpdate(db *sql.DB)string{

	//log.Println(parser.ToUrlTag(clanWorker.Clan.Tag))
	playerTags := update.GetTagByClans(parser.ToUrlTag(clanWorker.Clan.Tag))


	allErrors := update.Players(db, playerTags, 0)

	for _, err := range allErrors {
		log.Println(err)
	}
	return clanWorker.Clan.Name
}


