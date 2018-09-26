package labels

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func InsertLabel(DB *sql.DB,label Label)error{

	_,err := DB.Exec("INSERT INTO Trello.Labels (ID,IDboard,nameLabel,color) VALUES (?,?,?,?);",label.ID,label.IDboard,label.nameLabel,label.color)


	return err

}



func Update(DB *sql.DB,label Label)error{

	exists,err := validators.ExistsElementInColumn(DB,"Labels",label.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return updatelabel(DB,label)
	}


	return InsertLabel(DB,label)

}


func updatelabel(DB *sql.DB,label Label)error{

	_,err := DB.Exec("UPDATE Trello.Labels SET IDboard=?,nameLabel=?,color=? WHERE ID=?",label.IDboard,label.nameLabel,label.color,label.ID)

	return err
}



