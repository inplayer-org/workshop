package update

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
)


//GetRequestForPlayersFromClan - Refres Button - Refreshing all informations for players in 1 clan and updating in DB that inf
func GetRequestForPlayersFromClan(db *sql.DB,clanTag string)error{
	client := _interface.NewClient()
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
	players,err:=clientInterface.GetRequestForPlayer(parser.ToUrlTag(playerTag))

	if err!=nil{
		done<-err
	}else{
		var i int
		if players.LocationID!=nil {
			i=players.LocationID.(int)
		}else{
			i=0
		}

		err:=queries.UpdatePlayer(db,players,i)
		done<-err
	}
}

