package relations


import (
"database/sql"
	"strconv"
)

func InsertIntoMembersCardsRel(DB *sql.DB,idCard string,idMember string)error{

	_,err := DB.Exec("INSERT INTO Trello.Cards_Members_REL (IDcard, IDmember) VALUES (?,?);",idCard,idMember)
	return err

}

func deleteMembersCardsRel(DB *sql.DB,idCard string,idMember ...string)error{

	query := "DELETE FROM Cards_Members_REL WHERE IDcard=" + strconv.Quote(idCard)

	for _,elem := range idMember{
		query+=  " && IDlabel!=" + strconv.Quote(elem)
	}
	query+=";"

	_,err := DB.Exec(query)

	return err
}
