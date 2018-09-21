package userinfo

import (
	"database/sql"

)

//insertInto inserts userinfo into database
func InsertIntoUserInfoTable(db *sql.DB,id int,username string,fullname string)error{

	_, err := db.Exec("INSERT INTO UserInfo(ID,UserName,FullName) VALUES ((?),(?),(?));", id, username, fullname)

	if err != nil {
		return err
	}

	return nil

}

//update for an id updates a userinfo
func UpdateUserInfoTable(db *sql.DB,id int,username string,fullname string)error {

	_, err := db.Exec("UPDATE UserInfo SET ID=(?),UserName=(?),FullName=(?) WHERE ID=(?)", username, fullname, id)

	if err != nil {
		return err
	}

	return nil
}


// Returning slice of Locations Info from DB Table locations ALLInfo about Location
func GetAllUsers(db *sql.DB)([]Userinfo,error){

	rows, _ := db.Query("SELECT ID,UserName,FullName from UserInfo")


	defer rows.Close()

	return locationrows(rows)
}

func locationrows (rows *sql.Rows)([]Userinfo,error){
	var users  []Userinfo

	for rows.Next() {
		var u Userinfo
		err:=rows.Scan(&u.ID,&u.UserName,&u.FullName)

		if err!=nil {
			return nil,err
		}

		users=append(users,u)
	}

	return users,nil
}