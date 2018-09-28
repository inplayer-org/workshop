package members

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (member *Member)Insert(DB *sql.DB)error{

	_,err := DB.Exec("INSERT INTO Trello.Members (ID,fullname,initials,username,email) VALUES (?,?,?,?,?);",member.ID,member.FullName,member.Initials,member.Username,member.Email)


	return err

}

func  (member *Member)Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Members",member.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return member.updateByID(DB)
	}

	exists,err = validators.ExistsElementInColumn(DB,"Members",member.Username,"username")

	if err!= nil{
		return err
	}

	if exists{
		return member.updateByUsername(DB)
	}

	return member.Insert(DB)

}

func (member *Member)updateByID(DB *sql.DB)error{

	_,err := DB.Exec("UPDATE Trello.Members SET fullname=?,initials=?,username=?,email=? WHERE ID=?",member.FullName,member.Initials,member.Username,member.Email,member.ID)

	return err
}

func (member *Member)updateByUsername(DB *sql.DB)error{

	_,err := DB.Exec("UPDATE Trello.Members SET ID=?,fullname=?,initials=?,email=? WHERE username=?",member.ID,member.FullName,member.Initials,member.Email,member.Username)

	return err
}




// Returning string(member username from DB table boards)
func GetMemberUsername(db *sql.DB,memberID string)(string,error){


	var username string

	err := db.QueryRow("SELECT fullname FROM Members WHERE ID=?",memberID).Scan(&username)

	if err!=nil{
		return username,err
	}

	return username,nil

}


// Returns slice of all members present in the database
func GetAllMembers(db *sql.DB) ([]Member, error) {

	var members []Member
	var member Member

	rows, err := db.Query("SELECT ID,fullname,initials,username,email FROM Members;")

	if err != nil {
		return members, err
	}

	for rows.Next() {
		err := rows.Scan(&member.ID, &member.FullName,&member.Initials,&member.Username,&member.Email)

		if err != nil {
			return members, err
		}

		members = append(members, member)
	}

	return members, nil
}