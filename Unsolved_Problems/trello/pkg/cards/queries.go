package cards

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (c *Card)Insert(DB *sql.DB)error{

	_,err:=DB.Exec("INSERT into Trello.Cards (ID,checkItems,checkItemsChecked,description,dateLastActivity,descrip,shortLink,shortUrl) values (?,?,?,?,?,?,?,?);",c.ID,c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortLink,c.ShortURL)

	return err

	}

func  (c *Card)Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Cards",c.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return c.updateByID(DB)
	}

	exists,err = validators.ExistsElementInColumn(DB,"Cards",c.ShortLink,"shortLink")

	if err!= nil{
		return err
	}

	if exists{
		return c.updateByShortLink(DB)
	}
	exists,err = validators.ExistsElementInColumn(DB,"Cards",c.ShortURL,"shortURL")

	if err!= nil{
		return err
	}

	if exists{
		return c.updateByShortURL(DB)
	}


	return c.Insert(DB)

}
