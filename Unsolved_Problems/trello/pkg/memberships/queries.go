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


func GetBoardsByMember (db *sql.DB,memberID string)([]Membership,error){

	var boards []Membership

	rows,err:=db.Query("SELECT IDboard FROM Memberships Where IDmember Like (?)","%"+memberID+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){


		var m Membership
		err = rows.Scan(&m.IDboard)

		if err !=nil {
			return nil,err
		}

		boards = append(boards,m)
	}


	return boards,nil
}


func GetMemberByBoards (db *sql.DB,boardID string)([]Membership,error){

	var members []Membership

	rows,err:=db.Query("SELECT IDmember FROM Memberships Where IDboard Like (?)","%"+boardID+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){


		var m Membership
		err = rows.Scan(&m.IDmember)

		if err !=nil {
			return nil,err
		}

		members = append(members,m)
	}


	return members,nil
}
