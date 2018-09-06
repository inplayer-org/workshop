package workers

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
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

	client := _interface.NewClient()

	//log.Println(clanWorker.Clan.Tag)
	playerTags := client.GetTagByClans(clanWorker.Clan.Tag)

	//log.Println(playerTags)
	for _,nextPlayerTag := range playerTags{

		currentPlayer,err := client.GetRequestForPlayer(nextPlayerTag)

		if err!=nil{
			log.Println(err)
		}else{
			queries.UpdatePlayer(db,currentPlayer,0)
		}
	}

	return clanWorker.Clan.Name
}


