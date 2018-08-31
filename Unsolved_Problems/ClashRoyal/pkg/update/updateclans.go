package update

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/get"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)

func UpdateClans(db *sql.DB, clan []structures.Clan) error {

	for _, elem := range clan {

		//if not exist
		_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES (?,?)", elem.Tag, elem.Name)

		if err != nil {
			return err
		}

	}
	_, err := db.Exec("UPDATE clans SET clanTag=(?), clanName=(?) WHERE clanTag=(?)")
	if err != nil {
		return err
	}

	return err
}

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

func GetRequestForPlayersFromClan(db *sql.DB,clanTag string)int{

	clan := get.GetTagByClans(parser.ToUrlTag(clanTag))
	if len(clan)<=0{
		return 404
	}
	done := make(chan int)

	countChan := 0
	for _,playerTag := range clan {
		go chanRequest(db,playerTag,done)
		countChan++
	}
	for ;countChan>0;countChan--{
		log.Println("done = ",<-done)
	}
	return 0
}

func chanRequest(db *sql.DB,playerTag string,done chan <- int){
	done<- GetRequestForPlayer(db, playerTag)
}