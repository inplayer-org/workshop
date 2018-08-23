package update

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)


func UpdateClans(db *sql.DB,clan []structures.Clan)error{

for _,elem:=range clan {

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
