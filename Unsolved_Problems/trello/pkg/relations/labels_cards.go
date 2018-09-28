package relations

import (
	"database/sql"
	"strconv"
)

func InsertCardsLabelsRelation(DB *sql.DB,idCard string,idLabel string)error{

	_,err := DB.Exec("INSERT INTO Trello.Cards_Labels_REL (IDcard, IDlabel) VALUES (?,?);",idCard,idLabel)
	return err

}

//Deletes all cards labels relations that aren't present in the idLabel slice
func DeleteRemovedCardsLabelsRelation(DB *sql.DB,idCard string,idLabel ...string)error{

	query := "DELETE FROM Cards_Labels_REL WHERE IDcard=" + strconv.Quote(idCard)

	for _,elem := range idLabel{
		query+=  " && IDlabel!=" + strconv.Quote(elem)
	}
	query+=";"

	_,err := DB.Exec(query)

	return err
}

func UpdateCardsLabelsRelationsForCard(DB *sql.DB,idCard string){





}