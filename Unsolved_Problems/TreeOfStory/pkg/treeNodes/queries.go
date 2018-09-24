package treeNodes

import (
	"database/sql"
	"github.com/pkg/errors"
)

//Inserts the first node when a user is created (Only should be used when creating new User)
func InsertAbout(DB *sql.DB,userID int)error{

	exist,err := nodeExistsForUser(DB,Node{UserID:userID,ParentNode:"/",NodeName:"about"})

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
func Insert(DB *sql.DB,newNode Node)error{

	err := validateTreeNodeForInsert(DB,newNode)

	if err!=nil{
		return err
	}

	_,err = DB.Exec(`Insert into TreeNodes(userID,nodeName,parentNode,nodeWeight,fileName) Values (?,?,?,?,?);`,newNode.UserID,newNode.NodeName,newNode.ParentNode,newNode.NodeWeight,newNode.FileName)

	if err!=nil{
		return err
	}

	return nil


}

//Returns the weight for a node
func GetWeightForNode(DB *sql.DB,Node Node)(int,error){

	var weight int

	err := DB.QueryRow("SELECT nodeWeight FROM TreeNodes WHERE userID=? && nodeName=?",Node.UserID,Node.NodeName).Scan(&weight)

	if err!=nil{
		return weight,err
	}

	return weight,nil

}

