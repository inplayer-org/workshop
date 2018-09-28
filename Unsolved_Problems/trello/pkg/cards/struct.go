package cards

import (
	"time"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"database/sql"
)

type Card struct {
	ID string `json:"id"`
	Badges struct {
		CheckItems        int  `json:"checkItems"`
		CheckItemsChecked int  `json:"checkItemsChecked"`
		Description       bool `json:"description"`
	} `json:"badges"`
	DateLastActivity time.Time `json:"dateLastActivity"`
	Descrip string `json:"desc"`
	ShortLink string `json:"shortLink"`//POSSIBLE CHANGE
	ShortURL string `json:"shortUrl"`//POSSIBLE CHANGE
	IDmembers []string `json:"idMembers"`

}

func (c *Card) NewDataStructure() interfaces.DataStructure{
	return c
}

func (c *Card) updateByID(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Trello.Cards SET checkItems=?,checkItemsChecked=?,description=?,dateLastActivity=?,descrip=?,shortLink=?,shortUrl=? WHERE ID=?",c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortLink,c.ShortURL,c.ID)

	return err

}

func (c *Card) updateByShortLink(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Trello.Cards SET ID=?,checkItems=?,checkItemsChecked=?,description=?,dateLastActivity=?,descrip=?,shortUrl=? WHERE ShortLink=?",c.ID,c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortURL,c.ShortLink)

	return err

}

func (c *Card) updateByShortURL(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Trello.Cards SET ID=?,checkItems=?,checkItemsChecked=?,description=?,dateLastActivity=?,descrip=?,shortLink=? WHERE shortUrl=?",c.ID,c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortLink,c.ShortURL)

	return err

}
