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
		return errors.Errorf("Error in inserting node ->about<- into our database, nodeName for this userID(%d) already exists",userID)
	}

	_,err = DB.Exec(`Insert into TreeNodes(userID,nodeName,parentNode,nodeWeight,fileName) Values (?,"about","/",0,"about");`,userID)

	if err!=nil{
		return err
	}

	return nil
}

//Unchecked, needs to be tested
func Insert(DB *sql.DB,userID int,nodeName string,parentNode string,nodeWeight int,fileName string)error{

	exist,err := nodeExistsForUser(DB,userID,nodeName)

	if err!=nil{
		return err
	}

	if exist{
		return errors.Errorf("Error in inserting node %s into our database, nodeName for this userID(%d) already exists",nodeName,userID)
	}

	exist,err = nodeExistsForUser(DB,userID,parentNode)

	if err!=nil {
		return err
	}

	if !exist{
		return errors.Errorf("Error in inserting node %s into our database, nodeParent for this userID(%d) doesn't exist",parentNode)
	}

	weightCheck,err := GetWeightForNode(DB,userID,parentNode)

	if err!=nil{
		return err
	}

	if weightCheck>nodeWeight{
		return errors.Errorf("Error in inserting node %s into our database, nodeWeight(%d) for this node can't be lower than the weight of it's parentNode(%d)",parentNode,nodeWeight,weightCheck)
	}

	_,err = DB.Exec(`Insert into TreeNodes(userID,nodeName,parentNode,nodeWeight,fileName) Values (?,?,?,?,?);`,userID,nodeName,parentNode,nodeWeight,fileName)

	if err!=nil{
		return err
	}

	return nil


}

func GetWeightForNode(DB *sql.DB,userID int,nodeName string)(int,error){

	var weight int

	err := DB.QueryRow("SELECT nodeWeight FROM TreeNodes WHERE userID=? && nodeName=?",userID,nodeName).Scan(&weight)

	if err!=nil{
		return weight,err
	}

	return weight,nil

}

