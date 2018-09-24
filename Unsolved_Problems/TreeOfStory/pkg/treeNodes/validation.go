package treeNodes

import (
	"database/sql"
	"github.com/pkg/errors"
)

//Checks if node with a combination of userID and nodeName exists in the database
func nodeExistsForUser(db *sql.DB,newNode Node)(bool,error){

	var exists int

	err := db.QueryRow("SELECT Exists (SELECT * FROM TreeNodes WHERE userID=? && nodeName=?);",newNode.UserID,newNode.NodeName).Scan(&exists)

	if err!=nil{
		return false,err
	}

	if exists == 1{
		return true,nil
	}

	return false,nil

}

//Validate all of the logical rules too sse if a node is eligible for inserting
func validateTreeNodeForInsert(DB *sql.DB,newNode Node)error{

	exist,err := nodeExistsForUser(DB,newNode)

	if err!=nil{
		return err
	}

	if exist{
		return errors.Errorf("Error in inserting node %s into our database, nodeName for this userID(%d) already exists",newNode.NodeName,newNode.UserID)
	}

	err = parentEligibleCheck(DB,newNode)

	if err!=nil{
		return err
	}

	return nil

}

//Check if the parent node name exists in the database and if the parent's weight is lower (as it should be) than the new node
func parentEligibleCheck(DB *sql.DB,newNode Node)error{

	//Checks if a node with the node's parent name exists in our database
	exist,err := nodeExistsForUser(DB,Node{UserID:newNode.UserID,NodeName:newNode.ParentNode})

	if err!=nil {
		return err
	}

	//Returns error if there isn't a node with node's parent name in our database
	if !exist{
		return errors.Errorf("Error in inserting node %s into our database, nodeParent for this userID(%d) doesn't exist",newNode.ParentNode)
	}

	//Send a temp Node structure with ParentNode as the nodeName and current userID to check if it exists in our database
	parentWeight,err := GetWeightForNode(DB,Node{UserID:newNode.UserID,NodeName:newNode.ParentNode})

	if err!=nil{
		return err
	}

	//Checks if newNode's weight is higher than of it's parent weight
	if parentWeight>newNode.NodeWeight{
		return errors.Errorf("Error in inserting node %s into our database, nodeWeight(%d) for this node can't be lower than the weight of it's parentNode(%d)",newNode.ParentNode,newNode.NodeWeight,parentWeight)
	}

	return nil
}