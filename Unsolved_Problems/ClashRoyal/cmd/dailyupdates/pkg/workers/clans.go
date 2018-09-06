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

//Worker for sending request for all player tags in a clan (present in the database) to the clash royale api and writing the response to database
func (clanWorker *ClanWorker) FinishUpdate(db *sql.DB,client _interface.ClientInterface)string{

	//Get PlayerTags structure of all players in the clan
	players,err := client.GetTagByClans(clanWorker.Clan.Tag)

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

	return clanWorker.Clan.Name
}


