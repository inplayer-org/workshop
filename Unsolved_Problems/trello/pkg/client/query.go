package client

import (
	"net/url"
)

func QueryForCardsRequest()url.Values{
	var req url.URL

	q:=req.Query()


	q.Add("cards","all")
	q.Add("card_pluginData","true")


	return q

}
