package relations


import (
"database/sql"
"fmt"
"log"
)

func InsertIntoMembersCardsRel(DB *sql.DB,idCard string,idMember string)error{

	_,err := DB.Exec("INSERT INTO Trello.Cards_Members_REL (IDcard, IDmember) VALUES (?,?);",idCard,idMember)
	return err

}

func deleteMembersCardsRel(DB *sql.DB,idCard string,idMember ...string)error{

	query := fmt.Sprintf(`DELETE FROM Cards_Members_REL WHERE IDcard="%s"`,idCard)


	for _,elem := range idMember{
		query+= fmt.Sprintf(` && IDlabel!="%s"`,elem)
	}
	query+=";"

	_,err := DB.Exec(query)

	log.Println("Executed query : ",query)

	return err
}