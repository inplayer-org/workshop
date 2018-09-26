package client

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"encoding/json"
)

func (c *MyClient)GetMember(memberID string)(members.Member,error){

	var member members.Member

	urlStr:="https://api.trello.com/1/members/"+memberID
	req,err:=NewGetRequest(urlStr)

	if err!=nil{
		return member,err
	}
	resp,err:=c.client.Do(req)

	if err!=nil{
		return member,err
	}

	//fail to parse header,timeout,no header provided
	//if err:=errors.CheckStatusCode(resp);err!=nil{
	//    return clan,err
	//}
	json.NewDecoder(resp.Body).Decode(&member)

	return member,nil


}

func (c *MyClient)GetCard(cardID string)(cards.Card,error){

	var card cards.Card

	urlStr:="https://api.trello.com/1/cards/"+cardID
	req,err:=NewGetRequest(urlStr)

	if err!=nil{
		return card,err
	}
	resp,err:=c.client.Do(req)

	if err!=nil{
		return card,err
	}

	//fail to parse header,timeout,no header provided
	//if err:=errors.CheckStatusCode(resp);err!=nil{
	//    return clan,err
	//}
	json.NewDecoder(resp.Body).Decode(&card)

	return card,nil

}
