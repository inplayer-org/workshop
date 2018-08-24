package queries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"fmt"
)

func UpdateClans(db *sql.DB,clan structures.Clan)error{
	//log.Println("Tried to insert clan -> ",clan)
	if !Exists(db,"clans","clanTag", clan.Tag) {
		_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES (?,?)", clan.Tag, clan.Name)

		if err != nil {
			fmt.Println("1 ->",err)
		}
	}else{
		_,err := db.Exec("UPDATE clans SET clanName=(?) WHERE clanTag=(?)", clan.Name,clan.Tag)
		if err != nil {
			fmt.Println("2 ->",err)
		}

		//fmt.Println("3 ->",err)
	}

	return nil

}