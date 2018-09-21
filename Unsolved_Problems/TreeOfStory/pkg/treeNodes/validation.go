package treeNodes

import "database/sql"

func nodeExistsForUser(db *sql.DB,userID int,nodeName string)(bool,error){

	var exists int

	err := db.QueryRow("SELECT Exists (SELECT * FROM TreeNodes WHERE userID=? && nodeName=?);",userID,nodeName).Scan(&exists)

	if err!=nil{
		return false,err
	}

	if exists == 1{
		return true,nil
	}

	return false,nil

}
