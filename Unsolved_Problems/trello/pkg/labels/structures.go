package labels

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type Label struct {
	ID string `json:"id"`
	IDboard string `json:"idBoard"`
	nameLabel string `json:"name"`
	color string `json:"color"`

}

func (label *Label) NewDataStructure() interfaces.DataStructure{
	return label
}