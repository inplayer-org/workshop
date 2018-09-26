package client

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"encoding/json"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/errors"
)

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


func (c *MyClient)GetCard(cardID string)(interfaces.DataStructure,error){

	var card cards.Card

	urlStr:="https://api.trello.com/1/cards/"+cardID
	req,err:=NewGetRequest(urlStr)

	if err!=nil{
		return card.NewDataStructure(),err
	}
	resp,err:=c.client.Do(req)

	if err!=nil{
		return card.NewDataStructure(),err
	}

	//fail to parse header,timeout,no header provided
	if err:=errors.CheckStatusCode(resp);err!=nil{
	    return card.NewDataStructure(),err
	}
	json.NewDecoder(resp.Body).Decode(&card)

	return card.NewDataStructure(),nil

}
