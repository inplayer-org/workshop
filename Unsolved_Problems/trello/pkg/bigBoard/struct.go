package bigBoard

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/lists"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/memberships"
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
	Memberships []memberships.Membership `json:"memberships"`
}

func (bb *BigBoard) NewDataStructure() interfaces.DataStructure {
	return bb
}
