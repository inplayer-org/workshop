package members

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type Member struct{
	ID string `json:"id"`
	FullName string `json:"fullName"`
	Initials string `json:"initials"`
	Username string `json:"username"`
	Email string `json:"email"`
	IDboards []string `json:"idBoards"`
}

func (m *Member) NewDataStructure() interfaces.DataStructure{
	return m
}

func (m *Member)GetIDboards()[]string{
	return m.IDboards
}
