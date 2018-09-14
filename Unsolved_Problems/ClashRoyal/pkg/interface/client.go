package _interface

import (
	"encoding/json"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"net/url"
	"strconv"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/clans"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/playerStats"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/playerTags"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/cards"
)


//ClientInterface imeto ne e dobro
type ClientInterface interface {
	GetLocations() (locations.Locations,error)
	GetPlayerTagsFromLocation(int) (playerTags.PlayerTags,error)
	GetRequestForPlayer(string) (playerStats.PlayerStats,error)
	GetTagByClans(string) (playerTags.PlayerTags,error)
	GetClan(string)(clans.Clan,error)
	GetCards() (cards.Cards,error)
}

//MyClient structure have client that Do rquests
type MyClient struct {
	client *http.Client
}

//NewClient constructs MyClient
func NewClient() ClientInterface {
	return &MyClient{&http.Client{},

	}
}

//SetHeaders sets the headers to make the request
func SetHeaders(req *http.Request){
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")
}
//Returning all cards from API and error check
func (c *MyClient)GetCards() (cards.Cards,error) {
	var card cards.Cards

	req,err:=NewGetRequest("https://api.clashroyale.com/v1/cards")

	//cant parse url
	if err != nil {
		return card,errors.Default("URL",err)
	}

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err:= errors.CheckStatusCode(resp);err!=nil {
		return card,err
	}

	json.NewDecoder(resp.Body).Decode(&card)

	return card,nil
}

func (c *MyClient)GetClan(tag string)(clans.Clan,error){
	clanTag:=parser.ToRequestTag(tag)

	var clan clans.Clan
	urlStr:="https://api.clashroyale.com/v1/clans/"+clanTag
	req,err:=NewGetRequest(urlStr)

	if err!=nil{
		return clan,err
	}
	resp,err:=c.client.Do(req)


	//fail to parse header,timeout,no header provided
	if err:=errors.CheckStatusCode(resp);err!=nil{
		return clan,err
	}
	json.NewDecoder(resp.Body).Decode(&clan)

	return clan,nil


}

//Get request for Clans with clanTag as string returning all Tagsmembers of 1 clan
func (c *MyClient) GetTagByClans(clanTag string) (playerTags.PlayerTags,error) {
	tag:=parser.ToRequestTag(clanTag)

	var  tags playerTags.PlayerTags
	urlStr :="https://api.clashroyale.com/v1/clans/"+tag+"/members"
	req,err:=NewGetRequest(urlStr)

	//fail to parse url
	if err!= nil {
		return tags,errors.Default("URL",err)
	}

	resp,err:=c.client.Do(req)


	//fail to parse header,timeout,no header provided
	if err:=errors.CheckStatusCode(resp);err!=nil{
		return tags,err
	}
	json.NewDecoder(resp.Body).Decode(&tags)
	fmt.Println(tags)
	return tags,nil

}

//GetRequestForPlayer makes request and gets rankedPlayer tag name wins losses trophies clanTag and locationID
func (c *MyClient) GetRequestForPlayer (tag string) (playerStats.PlayerStats,error) {

	rtag:=parser.ToRequestTag(tag)

	var currentPlayer playerStats.PlayerStats

	urlStr := "https://api.clashroyale.com/v1/players/"

	url.Parse(urlStr+rtag)

	req,err:=NewGetRequest(urlStr+rtag)

	if err!=nil{
		return currentPlayer,errors.Default("URL",err)
	}

	for {
		resp, err := c.client.Do(req)

		if err:=errors.CheckStatusCode(resp);err!=nil{
			return currentPlayer,err
		}

//		if resp.StatusCode>=200 && resp.StatusCode<=300{
 		if err:=errors.CheckStatusCode(resp);err==nil{
			json.NewDecoder(resp.Body).Decode(&currentPlayer)
			currentPlayer.Tag = tag
		//	queries.UpdatePlayer(c.db,currentPlayer,0)  not using anymore updating rankedPlayer in handlres >>>

			break
		}
		//log.Println("REQUEST PROBLEM !! -> ",resp.Status,",  Retrying ...")
		//if resp.StatusCode==http.StatusNotFound{
			return currentPlayer,err
		//}
	}

	return currentPlayer,nil
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


func (c *MyClient) GetLocations()(locations.Locations,error){

	var locs locations.Locations

	req,err:=NewGetRequest("https://api.clashroyale.com/v1/locations")

	//cant parse url
	if err != nil {
		return locs,errors.Default("URL",err)
	}

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err:= errors.CheckStatusCode(resp);err!=nil {
		return locs,err
	}

	json.NewDecoder(resp.Body).Decode(&locs)

//	Locs(c.db,locations) updating locations db in handlers not here anymore

	return locs,nil
}

func (c *MyClient)GetPlayerTagsFromLocation(id int)(playerTags.PlayerTags,error)  {

	var tags playerTags.PlayerTags

	urlStr:="https://api.clashroyale.com/v1/locations/" + strconv.Itoa(id) + "/rankings/players"


	req,err:=NewGetRequest(urlStr)

	//fail to parse url
	if err!=nil {
		return tags,errors.Default("URLerr",err)
	}

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err:= errors.CheckStatusCode(resp);err!=nil {
		return tags,err
	}

	json.NewDecoder(resp.Body).Decode(&tags)

	return tags,nil

}