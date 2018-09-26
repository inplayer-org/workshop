package interfaces

type ClientInterface interface {
	GetMember(string)(DataStructure,error)
	GetLabel(string)(DataStructure,error)
	GetList(string)(DataStructure,error)
}
