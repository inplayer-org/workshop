package client

import (
	"net/url"
)

func QueryForCardsFromBoard()url.Values{
	var req url.URL

	q:=req.Query()

	q.Add("cards","all")
	q.Add("card_pluginData","true")

	return q

}

func BigBoardQuery()url.Values{
	var req url.URL

	q:=req.Query()

	q.Add("cards","all")
	q.Add("cards_fields","all")
	q.Add("labels","all")
	q.Add("label_fields","all")
	q.Add("card_members","all")
	q.Add("card_member_fields","all")
	q.Add("lists","all")
	q.Add("list_fields","all")
	q.Add("members","all")
	q.Add("members_fields","all")

	return q

}
