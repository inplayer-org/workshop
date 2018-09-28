package lists

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (list *List)Insert(DB *sql.DB)error{

	_,err := DB.Exec("INSERT INTO Trello.Lists (ID,nameList,IDboard) VALUES (?,?,?);",list.ID,list.Name,list.IDBoard)


	return err

}

func  (list *List)Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Lists",list.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return list.updateByID(DB)
	}

	return list.Insert(DB)

}

func (list *List)updateByID(DB *sql.DB)error{

	_,err := DB.Exec("UPDATE Trello.Lists SET nameList=?,IDboard=? WHERE ID=?",list.Name,list.IDBoard,list.ID)

	return err
}


// Returning slice of list structure with boardid u get listID and listName
func GetListFromBoard(db *sql.DB,boardID string)([]List,error){

	var lists []List

	rows,err:=db.Query("SELECT ID,nameList FROM Lists Where idBoard Like (?)","%"+boardID+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){


		var l List
		err = rows.Scan(&l.ID,&l.Name)

		if err !=nil {
			return nil,err
		}

		lists = append(lists,l)
	}


	return lists,nil
}