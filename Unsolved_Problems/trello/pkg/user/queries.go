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

	exists, err := validators.ExistsElementInColumn(DB, "Users", user.ID, "ID")

	if err != nil {
		return err
	}

	if exists {
		return user.updateByID(DB)
	}

	return user.Insert(DB)

}

func (user *User)  updateByID(DB *sql.DB) error {

	_, err := DB.Exec("UPDATE Trello.Users SET username=?,pass=?,token=? WHERE ID=?", user.Username, user.Password, user.Token,user.ID)

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
