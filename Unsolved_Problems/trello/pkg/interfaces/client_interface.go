package interfaces

type ClientInterface interface {
	GetMember(string) (DataStructure, error)
	GetLabel(string) (DataStructure, error)
	GetBoardInfo(string) (DataStructure, error)
}
