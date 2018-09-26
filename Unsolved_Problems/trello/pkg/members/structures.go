package members

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type Member struct{
	ID string `json:"id"`
	FullName string `json:"fullName"`
	Initials string `json:"initials"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func (member *Member) NewDataStructure() interfaces.DataStructure{
	return member
}