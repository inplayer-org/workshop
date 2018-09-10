package update

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"fmt"
)


//GetRequestForPlayersFromClan - Refresh Button - Refreshing all information for players in a clan and updating that information into a database
func GetRequestForPlayersFromClan(db *sql.DB,clanTag string)error{
	client := _interface.NewClient()
	tags,err:= client.GetTagByClans(clanTag)

	fmt.Println(tags)

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
	players,err:=clientInterface.GetRequestForPlayer(playerTag)

	if err!=nil{
		done<-err
	}else{
				err:=queries.UpdatePlayer(db,players,nil)
		done<-err
	}
}

