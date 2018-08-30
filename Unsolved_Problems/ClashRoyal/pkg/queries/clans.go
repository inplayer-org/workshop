package queries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"fmt"
)

func UpdateClans(db *sql.DB,clan structures.Clan)error{
	//log.Println("Tried to insert clan -> ",clan)
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

		//fmt.Println("3 ->",err)
	}

	return nil

}

func GetAllClans(db *sql.DB)([]structures.Clan,error){

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
}


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

		clans = append(clans,c)
	}

	return clans,nil

}

func GetClanName(db *sql.DB,clanTag string)(string,error){
	var clanName string
	err := db.QueryRow("SELECT clanName FROM clans WHERE clanTag=?",clanTag).Scan(&clanName)
	return clanName,err

}