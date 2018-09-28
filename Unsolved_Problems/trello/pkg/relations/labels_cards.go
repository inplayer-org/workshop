package relations

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
	"strconv"
)

func InsertCardsLabelsRelation(DB *sql.DB,idCard string,idLabel string)error{

	_,err := DB.Exec("INSERT INTO Trello.Cards_Labels_REL (IDcard, IDlabel) VALUES (?,?);",idCard,idLabel)
	return err

}

//Deletes all cards labels relations that aren't present in the idLabel slice
func DeleteRemovedCardsLabelsRelations(DB *sql.DB,idCard string,idLabel ...labels.Label)error{

	query := "DELETE FROM Cards_Labels_REL WHERE IDcard=" + strconv.Quote(idCard)

	for _,elem := range idLabel{
		query+=  " && IDlabel!=" + strconv.Quote(elem.ID)
	}
	query+=";"

	_,err := DB.Exec(query)

	return err
}

func UpdateCardsLabelsRelationsForCard(DB *sql.DB,idCard string,idLabel ...labels.Label)error{


	for _,elem := range idLabel{
		exists,err := validators.ExistsRelation(DB,"Cards_Labels_REL",idCard,elem.ID,"IDcard","IDlabel")

		if err!=nil{
			return nil
		}

		if !exists {
			err = InsertCardsLabelsRelation(DB,idCard,elem.ID)
			if err!=nil{
				return err
			}
		}
	}

	err := DeleteRemovedCardsLabelsRelations(DB,idCard,idLabel...)

	return err

}