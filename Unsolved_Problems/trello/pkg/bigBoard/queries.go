package bigBoard

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/relations"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
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

	exist,err:=validators.ExistsElementInColumn(DB,"Boards",bb.ID,"ID")

	if err != nil {
		return errors.Database(err)
	}

	if exist {

		err=bb.updateByID(DB)

		if err != nil {
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

	return bb.Insert(DB)
}

func (bb *BigBoard) updateByID(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Boards SET nameBoard=?,descrip=?,shortUrl=? WHERE ID=?",bb.Name,bb.Desc,bb.ShortURL,bb.ID)

	return err

}