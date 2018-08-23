package update

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"

)


func UpdateClans(db *sql.DB,clan []structures.Clan)error{

	for _,elem:=range clan {

		//if not exist
		_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES (?,?)", elem.Tag, elem.Name)

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


func GetAllClans(db sql.DB)(string, error) {

	var clans structures.Clan
	rows,err:=db.Query("SELECT (clanTag,clanName) FROM clans;")

	if err !=nil {
		return queries.ClansTable,err
	}

	count:=0
	for rows.Next(){
		err:=rows.Scan(&clans.structures.Clan[count].Tag,&clans.structures.Clan[count].Name)

		if err!=nil {
			return queries.ClansTable, err
		}

		count++

	}

	return queries.ClansTable,nil
}



func DailyUpdateClans(db *sql.DB)(string){

	clans,err:=GetAllClans()

	if err!=nil{
		return clans
	}

	err=UpdateClans(db, []structures.Clan)

	if err!=nil{
		return clans
	}

	return clans

}
