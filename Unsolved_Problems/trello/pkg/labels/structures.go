package labels

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type Label struct {
	ID string `json:"id"`
	IDboard string `json:"idBoard"`
	NameLabel string `json:"name"`
	Color string `json:"color"`

}

func (l *Label) NewDataStructure() interfaces.DataStructure{
	return l
}

