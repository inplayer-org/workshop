package client

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
	"encoding/json"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/labels"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/boards"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/lists"
	"net/url"
)

func (c *MyClient)GetLabel(labelID string)(interfaces.DataStructure,error){
	var label labels.Label

	urlStr:="https://api.trello.com/1/labels/"+labelID
	req,err:=NewGetRequest(urlStr,nil)

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
	req,err:=NewGetRequest(urlStr,nil)

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

func (c *MyClient) GetBoard(boardID string) (interfaces.DataStructure, error) {

	var board boards.Board

	urlStr := "https://api.trello.com/1/boards/" + boardID
	req, err := NewGetRequest(urlStr,nil)

	if err != nil {
		return board.NewDataStructure(), err
	}
	resp, err := c.client.Do(req)

	if err != nil {
		return board.NewDataStructure(), err
	}

	if err := errors.CheckStatusCode(resp); err != nil {
		return board.NewDataStructure(), err
	}

	json.NewDecoder(resp.Body).Decode(&board)

	return board.NewDataStructure(), nil
}

func (c *MyClient)GetCard(cardID string,values url.Values)(interfaces.DataStructure,error){

	var card cards.Card

	urlStr:="https://api.trello.com/1/cards/"+cardID
	req,err:=NewGetRequest(urlStr,values)

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

func (c *MyClient)GetList(listID string)(interfaces.DataStructure,error){

	var list lists.List

	urlStr:="https://api.trello.com/1/lists/"+listID
	req,err:=NewGetRequest(urlStr,nil)

	if err!=nil {
		return list.NewDataStructure(),err

	}
	resp,err:=c.client.Do(req)

	if err!=nil{
		return list.NewDataStructure(),err

	}
	if err:=errors.CheckStatusCode(resp);err!=nil{
		return list.NewDataStructure(),err
	}

	json.NewDecoder(resp.Body).Decode(&list)

	return list.NewDataStructure(),nil

}
