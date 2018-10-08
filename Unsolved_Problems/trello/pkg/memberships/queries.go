package memberships

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (membership *Membership)Insert(DB *sql.DB)error{

	_,err := DB.Exec("INSERT INTO Trello.Memberships (ID,IDmember,IDboard,memberType,unconfirmed,deactivated) VALUES (?,?,?,?,?,?);",membership.ID,membership.IDmember,membership.IDboard,membership.MemberType,membership.Unconfirmed,membership.Deactivated)


	return err

}

func  (membership *Membership)Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Memberships",membership.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return membership.updateByID(DB)
	}

	return membership.Insert(DB)

}

func (membership *Membership)updateByID(DB *sql.DB)error{

	_,err := DB.Exec("UPDATE Trello.Memberships SET IDmember=?,IDboard=?,memberType=?,unconfirmed=?,deactivated=? WHERE ID=?",membership.IDmember,membership.IDboard,membership.MemberType,membership.Unconfirmed,membership.Deactivated,membership.ID)

	return err
}


func GetBoardsByMember(db *sql.DB,memberid string)(string,error){



	var boardID string

	err := db.QueryRow("SELECT IDboard FROM Memberships WHERE IDmember=?",memberid).Scan(&boardID)

	if err!=nil{
		return boardID,err
	}

	return boardID,nil

}
