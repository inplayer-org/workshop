package update

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)


//GetAllClans - Returns slice of all structure Clan present in the database
func GetAllClans(db *sql.DB) ([]structures.Clan, error) {

	var clans []structures.Clan
	var clan structures.Clan

	rows, err := db.Query("SELECT clanTag,clanName FROM clans;")

	if err != nil {
		return clans, err
	}

	for rows.Next() {
		err := rows.Scan(&clan.Tag, &clan.Name)

		if err != nil {
			return clans, err
		}

		clans = append(clans, clan)
	}

	return clans, nil
}


//GetRequestForPlayersFromClan -
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

