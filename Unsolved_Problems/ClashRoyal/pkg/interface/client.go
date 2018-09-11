package _interface

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"encoding/json"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"net/url"
	"strconv"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
)


//ClientInterface imeto ne e dobro
type ClientInterface interface {
	GetLocations() (structures.Locations,error)
	GetPlayerTagsFromLocation(int) (structures.PlayerTags,error)
	GetRequestForPlayer(string) (structures.PlayerStats,error)
	GetTagByClans(string) (structures.PlayerTags,error)
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

//Get request for Clans with clanTag as string returning all Tagsmembers of 1 clan
func (c *MyClient) GetTagByClans(clanTag string) (structures.PlayerTags,error) {
	tag:=parser.ToRequestTag(clanTag)

	var  playerTags structures.PlayerTags
	urlStr :="https://api.clashroyale.com/v1/clans/"+tag+"/members"
	req,err:=NewGetRequest(urlStr)

	//fail to parse url
	if err!= nil {
		return playerTags,errors.Default("URL",err)
	}

	resp,err:=c.client.Do(req)


	//fail to parse header,timeout,no header provided
	if err:=errors.CheckStatusCode(resp);err!=nil{
		return playerTags,err
	}
	json.NewDecoder(resp.Body).Decode(&playerTags)
	fmt.Println(playerTags)
	return playerTags,nil

}

//GetRequestForPlayer makes request and gets players tag name wins losses trophies clanTag and locationID
func (c *MyClient) GetRequestForPlayer (tag string) (structures.PlayerStats,error) {

	tag=parser.ToRequestTag(tag)

	var currentPlayer structures.PlayerStats

	fmt.Println(tag)

	urlStr := "https://api.clashroyale.com/v1/players/"

	url.Parse(urlStr+tag)

	req,err:=NewGetRequest(urlStr+tag)

	if err!=nil{
		return currentPlayer,errors.Default("URL",err)
	}

	for {
		resp, err := c.client.Do(req)

		if err:=errors.CheckStatusCode(resp);err!=nil{
			return currentPlayer,err
		}

//		if resp.StatusCode>=200 && resp.StatusCode<=300{
 	if err:=errors.CheckStatusCode(resp);err!=nil{
			json.NewDecoder(resp.Body).Decode(&currentPlayer)
			currentPlayer.Tag = tag
		//	queries.UpdatePlayer(c.db,currentPlayer,0)  not using anymore updating players in handlres >>>

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


func (c *MyClient) GetLocations()(structures.Locations,error){

	var locations structures.Locations

	req,err:=NewGetRequest("https://api.clashroyale.com/v1/locations")

	//cant parse url
	if err != nil {
		return locations,errors.Default("URL",err)
	}

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err:= errors.CheckStatusCode(resp);err!=nil {
		return locations,err
	}

	json.NewDecoder(resp.Body).Decode(&locations)

//	Locs(c.db,locations) updating locations db in handlers not here anymore

	return locations,nil
}

func (c *MyClient)GetPlayerTagsFromLocation(id int)(structures.PlayerTags,error)  {

	var playerTags structures.PlayerTags

	urlStr:="https://api.clashroyale.com/v1/locations/" + strconv.Itoa(id) + "/rankings/players"


	req,err:=NewGetRequest(urlStr)

	//fail to parse url
	if err!=nil {
		return playerTags,errors.Default("URLerr",err)
	}

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err:= errors.CheckStatusCode(resp);err!=nil {
		return playerTags,err
	}

	json.NewDecoder(resp.Body).Decode(&playerTags)

	return playerTags,nil

}