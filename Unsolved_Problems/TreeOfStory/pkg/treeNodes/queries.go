package treeNodes

import (
	"database/sql"
	"github.com/pkg/errors"
)

//Inserts the first node when a user is created (Only should be used when creating new User)
func InsertAbout(DB *sql.DB,userID int)error{

	exist,err := nodeExistsForUser(DB,userID,"about")

	if err!=nil{
		return err
	}

	if exist{
		return errors.Errorf("Error in inserting node ->about<- into our database, nodeName for this user already exists")
	}

	_,err = DB.Exec(`Insert into TreeNodes(userID,nodeName,parentNode,nodeWeight,fileName) Values (?,"about","/",0,"about");`,userID)

	if err!=nil{
		return err
	}

	return nil
}


