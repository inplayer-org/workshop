package labels

import "database/sql"

func InsertLabel(DB *sql.DB,label Label)error{

	_,err := DB.Exec("INSERT INTO Trello.Labels (ID,IDboard,nameLabel,color) VALUES (?,?,?,?);",label.ID,label.IDboard,label.nameLabel,label.color)


	return err

}



func UpdateLabel(DB *sql.DB,label Label)error{

	_,err := DB.Exec("UPDATE Trello.Labels SET IDboard=?,nameLabel=?,color=? WHERE ID=?",label.IDboard,label.nameLabel,label.color,label.ID)

	return err
}
