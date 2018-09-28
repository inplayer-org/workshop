package cards

import (
	"time"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
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
	IDBoard string `json:"idBoard"`
	IDList string `json:"idList"`
	ShortLink string `json:"shortLink"`//POSSIBLE CHANGE
	ShortURL string `json:"shortUrl"`//POSSIBLE CHANGE
}

func (c *Card) NewDataStructure() interfaces.DataStructure{
	return c
}

