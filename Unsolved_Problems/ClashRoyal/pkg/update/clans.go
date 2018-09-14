package update

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/playerStats"
)


//GetRequestForPlayersFromClan - Refresh Button - Refreshing all information for rankedPlayer in a clan and updating that information into a database
func GetRequestForPlayersFromClan(db *sql.DB,client _interface.ClientInterface,clanTag string)error{
	tags,err:= client.GetTagByClans(clanTag)

	if err!=nil{

		return err
	}

	done := make(chan error)

	countChan := 0

	for _,playerTag := range tags.Player {
		go chanRequest(db,client,playerTag.Tag,done)
		countChan++
	}

	for ;countChan>0;countChan--{
		log.Println("done = ",<-done)
	}

	return nil
}

func chanRequest(db *sql.DB,clientInterface _interface.ClientInterface,playerTag string,done chan <- error){
	player,err:=clientInterface.GetRequestForPlayer(playerTag)

	if err!=nil{
		done<-err
	}else{
				err:= playerStats.UpdatePlayer(db,player,nil)
		done<-err
	}
}

