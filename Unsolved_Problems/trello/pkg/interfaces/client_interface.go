package interfaces

type ClientInterface interface {
	GetMember(string)(DataStructure,error)
	GetLabel(string)(DataStructure,error)
	GetBoard(string) (DataStructure, error)
	GetList(string)(DataStructure,error)
	GetCard(string)(DataStructure,error)
}
