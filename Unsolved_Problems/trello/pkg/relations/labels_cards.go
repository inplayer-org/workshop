package relations

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertCardsLabelsRelation(DB *sql.DB,idCard string,idLabel string)error{

	_,err := DB.Exec("INSERT INTO Trello.Cards_Labels_REL (IDcard, IDlabel) VALUES (?,?);",idCard,idLabel)
	return err

}

func deleteCardsLabelsRelation(DB *sql.DB,idCard string,idLabel ...string)error{

	query := fmt.Sprintf(`DELETE FROM Cards_Labels_REL WHERE IDcard="%s"`,idCard)


	for _,elem := range idLabel{
		query+= fmt.Sprintf(` && IDlabel!="%s"`,elem)
	}
	query+=";"

	_,err := DB.Exec(query)

	log.Println("Executed query : ",query)

	return err
}