package interfaces

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/boards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
)

type ClientInterface interface {
	GetMember(string) (members.Member, error)
	GetLabel(string) (labels.Label, error)
	GetBoardInfo(string) (boards.Board, error)
}
