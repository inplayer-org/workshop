package queries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

func UpdateClans(db *sql.DB,clan structures.Clan)error{

	if clan.Name=="" || clan.Tag==""{
		return nil
	}

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

	}

	return nil
}

//feature
/*func GetAllClans(db *sql.DB)([]structures.Clan,error){

	rows, _ := db.Query("SELECT * from clans")


	defer rows.Close()

	return clanRows(rows)
}

func clanRows(rows *sql.Rows)([]structures.Clan,error){
	var clans  []structures.Clan

	for rows.Next() {
		var c structures.Clan
		err:=rows.Scan(&c.Tag,&c.Name)

		if err!=nil {
			return nil,err
		}

		clans=append(clans,c)
	}

	return clans,nil
}*/

func GetClansLike(db *sql.DB,name string)([]structures.Clan,error){

	var clans [] structures.Clan
	rows,err:=db.Query("SELECT clanName,clanTag FROM clans Where clanName Like (?)","%"+name+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){

		var c structures.Clan
		err = rows.Scan(&c.Name,&c.Tag)

		if err !=nil {
			return nil,err
		}
		c.Tag = parser.ToRawTag(c.Tag)
		clans = append(clans,c)
	}

	return clans,nil
}

func GetClanName(db *sql.DB,clanTag string)(string,error){

	fmt.Println(clanTag)

	var clanName string

	err := db.QueryRow("SELECT clanName FROM clans WHERE clanTag=?",clanTag).Scan(&clanName)

	return clanName,err

}

//GetAllClans - Returns slice of all Clans present in the database
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