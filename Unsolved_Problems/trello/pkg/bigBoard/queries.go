package bigBoard

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/relations"
)

func(bb *BigBoard)Insert(DB *sql.DB)error{

	_,err:=DB.Exec("INSERT into Boards (ID,nameBoard,descrip,shortUrl) values (?,?,?,?);",bb.ID,bb.Name,bb.Desc,bb.ShortURL)

	if err!=nil{
		return errors.Database(err)
	}

	for _,list:=range bb.Lists {

		err=list.Update(DB)

		if err != nil {
			return errors.Database(err)
		}
	}

	for _,label:=range bb.Labels {

		err=label.Update(DB)

		if err != nil {
			return errors.Database(err)
	}
    }

	for _,member:=range bb.Members{

		err=member.Update(DB)

		if err != nil {
			return errors.Database(err)
		}

	}

    for _,card:=range bb.Cards{

    	err=card.Update(DB)

		if err != nil {
			return errors.Database(err)
		}

		err=relations.UpdateCardsLabelsRelationsForCard(DB,card.ID,card.Labels...)

		if err != nil {
			return errors.Database(err)
		}

		err=relations.UpdateMembersCardsRelTable(DB,card.ID,card.IDmembers...)

		if err != nil {
			return errors.Database(err)
		}
	}

return nil
}

func(bb *BigBoard)Update(DB *sql.DB)error{

	return nil
}
