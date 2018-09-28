package labels

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (label *Label) Insert(DB *sql.DB)error{

	_,err := DB.Exec("INSERT INTO Trello.Labels (ID,IDboard,nameLabel,color) VALUES (?,?,?,?);",label.ID,label.IDboard,label.NameLabel,label.Color)


	return err

}



func(label *Label) Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Labels",label.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return label.updatelabel(DB)
	}


	return label.Insert(DB)

}


func (label *Label) updatelabel(DB *sql.DB)error{

	_,err := DB.Exec("UPDATE Trello.Labels SET IDboard=?,nameLabel=?,color=? WHERE ID=?",label.IDboard,label.NameLabel,label.Color,label.ID)

	return err
}



