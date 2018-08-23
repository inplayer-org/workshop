package update

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/get"
)


func UpdateClans(db *sql.DB,clan get.Clans)error{

for _,elem:=range clan.Clans {

	//if not exist
	_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES ((?),(?))", elem.Tag, elem.Name)

	if err != nil {
		return err
	}

}
 _,err := db.Exec("UPDATE clans SET clanTag=(?), clanName=(?) WHERE clanTag=(?)")
	if err != nil {
		return err
	}

 return err
 }
