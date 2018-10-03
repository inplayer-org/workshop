package lists

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type List struct{
	ID string `json:"id"`
	Name string `json:"name"`
	IDBoard string `json:"idBoard"`
}

func (l *List) NewDataStructure() interfaces.DataStructure{
	return l
}
