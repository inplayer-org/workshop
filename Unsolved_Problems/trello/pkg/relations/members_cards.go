package relations


import (
"database/sql"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
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


func UpdateMembersCardsRelTable(DB *sql.DB,idCard string,idMember ...string)error{


	for _,elem := range idMember{
		exist,err := validators.ExistsRelation(DB,"Cards_Members_REL",idCard,elem,"IDcard","IDmember")

		if err!=nil{
			return nil
		}

		if !exist {
			err = InsertIntoMembersCardsRel(DB,idCard,elem)
			if err!=nil{
				return err
			}
		}
	}

	err := deleteMembersCardsRel(DB,idCard,idMember...)

	return err

}