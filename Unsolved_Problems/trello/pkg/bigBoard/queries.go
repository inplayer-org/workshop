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

	for _,list:=range bb.Lists {

		err=list.Insert(DB)

		if err != nil {
			return errors.Database(err)
		}
	}

	for _,label:=range bb.Labels {

		err=label.Insert(DB)

		if err != nil {
			return errors.Database(err)
	}
    }

    for _,card:=range bb.Cards{

    	err=card.Insert(DB)

		if err != nil {
			return errors.Database(err)
		}
	}

	for _,member:=range bb.Members{

		err=member.Insert(DB)

		if err != nil {
			return errors.Database(err)
		}
	}



return nil
}

func(bb *BigBoard)Update(DB *sql.DB)error{

	return nil
}
