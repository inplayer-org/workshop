package bigBoard

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/errors"
)

func(bb *BigBoard)Insert(DB *sql.DB)error{

	_,err:=DB.Exec("INSERT into Boards (ID,nameBoard,decrip,shortUrl) values (?,?,?,?);",bb.ID,bb.Name,bb.Desc,bb.ShortURL)

	if err!=nil{
		return errors.Database(err)
	}

	_,err=DB.Exec("INSERT into Lists (ID,nameList,IDboard) values (?,?,?);",bb.Lists.ID,bb.Lists.Name,bb.Lists.IDBoard)

	if err!=nil{
		return errors.Database(err)
	}


	return nil
}

func(bb *BigBoard)Update(DB *sql.DB)error{

	return nil
}
