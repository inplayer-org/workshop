package members

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func Insert(DB *sql.DB,member Member)error{

	_,err := DB.Exec("INSERT INTO Trello.Members (ID,fullname,initials,username,email) VALUES (?,?,?,?,?);",member.ID,member.FullName,member.Initials,member.Username,member.Email)


	return err

}

func Update(DB *sql.DB,member Member)error{

	exists,err := validators.ExistsElementInColumn(DB,"Members",member.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return updateByID(DB,member)
	}

	exists,err = validators.ExistsElementInColumn(DB,"Members",member.Username,"username")

	if err!= nil{
		return err
	}

	if exists{
		return updateByUsername(DB,member)
	}

	return Insert(DB,member)

}

func updateByID(DB *sql.DB,member Member)error{

	_,err := DB.Exec("UPDATE Trello.Members SET fullname=?,initials=?,username=?,email=? WHERE ID=?",member.FullName,member.Initials,member.Username,member.Email,member.ID)

	return err
}

func updateByUsername(DB *sql.DB,member Member)error{

	_,err := DB.Exec("UPDATE Trello.Members SET ID=?,fullname=?,initials=?,email=? WHERE username=?",member.ID,member.FullName,member.Initials,member.Email,member.Username)

	return err
}