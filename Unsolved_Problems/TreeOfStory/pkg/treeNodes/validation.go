package treeNodes

import "database/sql"

func nodeExistsForUser(db *sql.DB,userID int,nodeName string)error{

	Row := db.QueryRow("SELECT Exists (SELECT * FROM TreeNodes WHERE userID=? && nodeName=?);",userID,nodeName)

	var exists int

	err := Row.Scan(exists)

	if err!=nil{
		return err
	}

	


}
