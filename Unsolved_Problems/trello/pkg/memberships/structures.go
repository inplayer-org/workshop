package memberships

import "repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"

type Membership struct{
	ID string `json:"id"`
	IDmember string `json:"idMember"`
	IDboard string `json:"idBoard"`
	MemberType string `json:"memberType"`
	Unconfirmed bool `json:"unconfirmed"`
	Deactivated string `json:"deactivated"`
}

func (m *Membership) NewDataStructure() interfaces.DataStructure{
	return m
}