package cards

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"time"
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
	IDmembers []string `json:"idMembers"`
	Labels []labels.Label `json:"labels"`
	ShortLink string `json:"shortLink"`//POSSIBLE CHANGE
	ShortURL string `json:"shortUrl"`//POSSIBLE CHANGE


}

func (c *Card) NewDataStructure() interfaces.DataStructure{
	return c
}

func (c *Card) GetIDboards()[]string{
	var ret []string
	ret=append(ret,c.IDBoard)
	return ret
}

