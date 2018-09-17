package clans

import (

	"database/sql"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"fmt"

)

//Cheking for clan if exist in db if not inserting if exist updating
func UpdateClans(db *sql.DB,clan Clan)error{

	if clan.Name=="" || clan.Tag==""{
		return nil
	}

	if !parser.Exists(db,"clans","clanTag", clan.Tag) {

		_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES (?,?)", clan.Tag, clan.Name)

		if err != nil {
			return err
		}

	}else{
		_,err := db.Exec("UPDATE clans SET clanName=(?) WHERE clanTag=(?)", clan.Name,clan.Tag)

		if err != nil {
			return err
		}

	}

	return nil
}


// Returning slice of clan structure with clanname u get clanname and clanTag
func GetClansLike(db *sql.DB,name string)([]Clan,error){

	var clan [] Clan
	rows,err:=db.Query("SELECT clanName,clanTag FROM clans Where clanName Like (?)","%"+name+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){


		var c Clan
		err = rows.Scan(&c.Name,&c.Tag)

		if err !=nil {
			return nil,err
		}
		c.Tag = parser.ToRawTag(c.Tag)
		clan = append(clan,c)
	}

	if len(clan)==0{
		return nil,errors.Database(sql.ErrNoRows)
	}


	return clan,nil
}
// Returning string(clan name from DB table clans)
func GetClanName(db *sql.DB,clanTag string)(string,error){

	fmt.Println(clanTag)

	var clanName string

	err := db.QueryRow("SELECT clanName FROM clans WHERE clanTag=?",clanTag).Scan(&clanName)

	if err!=nil{
		return clanName,err
	}

	return clanName,nil

}

//GetAllClans - Returns slice of all Clans present in the database
func GetAllClans(db *sql.DB) ([]Clan, error) {

	var clanss []Clan
	var clan Clan

	rows, err := db.Query("SELECT clanTag,clanName FROM clans;")

	if err != nil {
		return clanss, err
	}

	for rows.Next() {
		err := rows.Scan(&clan.Tag, &clan.Name)

		if err != nil {
			return clanss, err
		}

		clanss = append(clanss, clan)
	}

	return clanss, nil
}