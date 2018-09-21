package treeNodes

type Node struct{
	ID int `json:"id"`
	UserID int `json:"userId"`
	NodeName string `json:"name"`
	ParentNode string `json:"parentNode"`
	NodeWeight int `json:"nodeWeight"`
	FileName string `json:"fileName"`
}


