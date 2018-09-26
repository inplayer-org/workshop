package labels

import "database/sql"

func InsertLabel(DB *sql.DB,label Label)error{

	_,err := DB.Exec("INSERT INTO Trello.Labels (ID,IDboard,nameLabel,color) VALUES (?,?,?,?);",label.ID,label.IDboard,label.nameLabel,label.color)


	return err

}



