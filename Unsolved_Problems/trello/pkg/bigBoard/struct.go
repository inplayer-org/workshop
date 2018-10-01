package bigBoard

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/lists"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
)

type BigBoard struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Desc           string      `json:"desc"`
	ShortURL       string      `json:"shortUrl"`
	Cards []cards.Card `json:"cards"`
	Labels []labels.Label `json:"labels"`
	Lists []lists.List `json:"lists"`
	Members []members.Member `json:"members"`
}

func (bb *BigBoard) NewDataStructure() interfaces.DataStructure {
	return bb
}

func (bb *BigBoard) GetIDboards()[]string{
	var ret []string
	ret=append(ret,bb.ID)
	return ret
}