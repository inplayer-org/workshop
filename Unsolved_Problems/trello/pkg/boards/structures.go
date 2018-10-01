package boards

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type Board struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	ShortUrl string `json:"shortUrl"`
}

func (board *Board) NewDataStructure() interfaces.DataStructure {
	return board
}

func (bb *Board) GetIDboards()[]string{
	var ret []string
	ret=append(ret,bb.ID)
	return ret
}
