package session

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (session *Session) Insert(DB *sql.DB) error {

	_, err := DB.Exec("INSERT INTO Trello.Sessions (uid,IDuser) VALUES (?,?);", session.UID, session.IDuser)

	return err

}

func (session *Session) Update(DB *sql.DB) error {

	exists, err := validators.ExistsElementInColumn(DB, "Sessions", session.UID, "uid")

	if err != nil {
		return err
	}

	if exists {
		return session.updateByID(DB)
	}

	return session.Insert(DB)

}

func ( session *Session)  updateByID(DB *sql.DB) error {

	_, err := DB.Exec("UPDATE Trello.Sessions SET IDuser=? WHERE uid=?",  session.IDuser,session.UID)

	return err
}

func (session *Session) Delete (DB *sql.DB) error {

	_, err := DB.Exec("DELETE  FROM Trello.Sessions WHERE uid=?;")

	return err

}


func GetFromSessions(db *sql.DB,sessionID string)(int,error){

	var IDuser int

	err := db.QueryRow("SELECT IDuser FROM Sessions WHERE uid=?",sessionID).Scan(&IDuser)

	if err!=nil{
		return IDuser,err
	}

	return IDuser,nil

}
