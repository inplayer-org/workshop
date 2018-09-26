package client

import (
	"encoding/json"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/errors"
)


//MyClient structure have client that Do rquests
type MyClient struct {
	client *http.Client
}

//NewClient constructs MyClient
func NewClient() interfaces.ClientInterface {
	return &MyClient{&http.Client{},

	}
}

//SetHeaders sets the headers to make the request
func SetHeaders(req *http.Request){

	q := req.URL.Query()
	q.Add("key","9ecdc5f04a4ccb643b83d4fd2b920416")
	q.Add("token","125a712b04063b34d2c22392704bb38a5fc88bb48f665c0f6bdf2d516f473c9d")


	req.Header.Add("Content-Type","application/json")
	req.URL.RawQuery = q.Encode()
}

//NewGetRequest makes the request with the headers
func NewGetRequest(url string)(*http.Request,error){
	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil {
		return nil, err
	}
	SetHeaders(req)
	return req,nil
}

func (c *MyClient)GetLabel(labelID string)(interfaces.DataStructure,error){
	var label labels.Label

	urlStr:="https://api.trello.com/1/labels/"+labelID
	req,err:=NewGetRequest(urlStr)

	if err!=nil {
		return label.NewDataStructure(),err

	}
	resp,err:=c.client.Do(req)

	if err!=nil{
		return label.NewDataStructure(),err

}
	if err:=errors.CheckStatusCode(resp);err!=nil{
		return label.NewDataStructure(),err
	}

	json.NewDecoder(resp.Body).Decode(&label)

	return label.NewDataStructure(),nil


}

func (c *MyClient)GetMember(memberID string)(interfaces.DataStructure,error){

	var member members.Member

	urlStr:="https://api.trello.com/1/members/"+memberID
	req,err:=NewGetRequest(urlStr)

	if err!=nil{
		return member.NewDataStructure(),err
	}
	resp,err:=c.client.Do(req)

	if err!=nil{
		return member.NewDataStructure(),err
	}


	//fail to parse header,timeout,no header provided
	if err:=errors.CheckStatusCode(resp);err!=nil{
	return member.NewDataStructure(),err
	}

	json.NewDecoder(resp.Body).Decode(&member)

	return member.NewDataStructure(),nil


}

