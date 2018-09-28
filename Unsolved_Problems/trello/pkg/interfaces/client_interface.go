package interfaces

import "net/url"

type ClientInterface interface {
	GetMember(string)(DataStructure,error)
	GetLabel(string)(DataStructure,error)
	GetBoard(string) (DataStructure, error)
	GetList(string)(DataStructure,error)
	GetCard(string,url.Values)(DataStructure,error)
}
