package user

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)



func (user *User) Insert(DB *sql.DB) error {

	_, err := DB.Exec("INSERT INTO Trello.Users (username,pass,token) VALUES (?,?,?);",user.Username, user.Password, user.Token)

	return err

}


//sa za sa ne trebe
func (user *User)  Update(DB *sql.DB) error {

	exists, err := validators.ExistsElementInColumn(DB, "Users", user.Username, "username")

	if err != nil {
		return err
	}

	if exists {
		return user.updatebyUsername(DB)
	}

	return user.Insert(DB)

}

func (user *User)  updatebyUsername(DB *sql.DB) error {

	_, err := DB.Exec("UPDATE Trello.Users SET pass=?,token=? WHERE username=?", user.Password, user.Token,user.Username)

	return err
}


func (user *User) Delete (DB *sql.DB) error {

	_, err := DB.Exec("DELETE FROM Trello.Users (ID,username,pass,token) VALUES (?,?,?,?);", user.ID, user.Username,user.Password,user.Token)

	return err

}



func GetFromUser(db *sql.DB,userID int)(User,error){

	var user User

	err := db.QueryRow("SELECT username,pass,token FROM Users WHERE ID=?",userID).Scan(&user.Username,&user.Password,&user.Token)

	if err!=nil{
		return user,err
	}

	return user,nil

}


func GetUserID(db *sql.DB,username string)(int,error){



	var id int

	err := db.QueryRow("SELECT ID FROM Users WHERE username=?",username).Scan(&id)

	if err!=nil{
		return id,err
	}

	return id,nil

}
