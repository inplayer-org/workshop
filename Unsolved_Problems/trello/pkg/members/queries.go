package members

import "database/sql"

func InsertMember(DB *sql.DB,member Member)error{

	_,err := DB.Exec("INSERT INTO Trello.Members (ID,fullname,initials,username,email) VALUES (?,?,?,?,?);",member.ID,member.FullName,member.Initials,member.Username,member.Email)


	return err

}

