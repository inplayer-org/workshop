package client

import (
	"net/url"
)

func CardsQuery()url.Values{
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
	q.Add("card_members","true")
	q.Add("card_member_fields","all")
	q.Add("lists","all")
	q.Add("list_fields","all")
	q.Add("members","all")
	q.Add("members_fields","all")
	q.Add("memberships","all")

	return q

}

func LabelsQuery()url.Values{

	var req url.URL

	q:=req.Query()

	q.Add("fields","all")

	return q

}

func ListQuery()url.Values{

	var req url.URL

	q:=req.Query()

	q.Add("fields","all")

	return q

}

func MembersQuery()url.Values{

	var req url.URL

	q:=req.Query()

	q.Add("fields","all")
	q.Add("cards","all")

	return q

}