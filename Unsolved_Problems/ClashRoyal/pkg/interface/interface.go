package _interface

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"encoding/json"
	"fmt"
	"database/sql"
	"net/http"
	"net/url"
)

//ClientInterface imeto ne e dobro
type ClientInterface interface {
	GetLocations() (structures.Locations,error)
	GetPlayerTagsFromLocation(int) (structures.PlayerTags,error)
	GetPlayerTagByClans(string) (structures.PlayerTags,error)
	GetRequestForPlayer(string) (int,error)

}

type MyClient struct {
	client *http.Client
	db *sql.DB
}

func NewClient(db *sql.DB) ClientInterface {
	return &MyClient{&http.Client{},
		db,
	}
}



func SetHeaders(req *http.Request){
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")
}

func NewGetRequest(url string)(*http.Request,error){
	return http.NewRequest("GET",url,nil)
}

func (c *MyClient) GetRequestForPlayer (db *sql.DB,tag string) (int,error) {

	var currentPlayer structures.PlayerStats


	urlStr := "https://api.clashroyale.com/v1/players/"

	url.Parse(urlStr+tag)

	req,err:=NewGetRequest(urlStr+tag)

	if err!=nil{
		fmt.Println(err)
	}

	for {
		resp, err := c.client.Do(req)

		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode>=200 && resp.StatusCode<=300{

			json.NewDecoder(resp.Body).Decode(&currentPlayer)
			currentPlayer.Tag = "#"+currentPlayer.Tag[1:]
			queries.UpdatePlayer(db,currentPlayer,0)

			break
		}
		//log.Println("REQUEST PROBLEM !! -> ",resp.Status,",  Retrying ...")
		if resp.StatusCode!=http.StatusNotFound{
			return 404,err
		}
	}

	return 0,err
}
func (c *MyClient) GetPlayerTagByClans(clanTag string) (structures.PlayerTags,error) {
	var  playerTags structures.PlayerTags
	urlStr :="https://api.clashroyale.com/v1/clans/"+clanTag+"/members"
	req,err:=NewGetRequest(urlStr)

	//fail to parse url
	if err!= nil {
		return playerTags,err
	}
	SetHeaders(req)
	resp,err:=c.client.Do(req)


	//fail to parse header,timeout,no header provided
	if err!=nil {
		return playerTags, err
	}
	json.NewDecoder(resp.Body).Decode(&playerTags)
	return playerTags,nil
	//should return parse Tags with %25 how?? cant do it with parser pkg(GLS help)
	}


func (c *MyClient) GetLocations()(structures.Locations,error){

	var locations structures.Locations

	req,err:=NewGetRequest("https://api.clashroyale.com/v1/locations")

	//cant parse url
	if err != nil {
		return locations,err
	}

	SetHeaders(req)

	fmt.Println(req.Header["authorization"])

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err!=nil {
		return locations,err
	}

	json.NewDecoder(resp.Body).Decode(&locations)

	Locs(c.db,locations)

	return locations,nil
}

func Locs(db *sql.DB,locations structures.Locations){

	for _,elem:=range locations.Location {

		if queries.Exists(db,"locations","id",strconv.Itoa(elem.ID)){
			queries.UpdateLocationsTable(db,elem.ID,elem.Name,elem.IsCountry,elem.CountryCode)
		} else {
			queries.InsertIntoLocationsTable(db,elem.ID,elem.Name,elem.IsCountry,elem.CountryCode)
		}

	}

}

func (c *MyClient)GetPlayerTagsFromLocation(id int)(structures.PlayerTags,error)  {

	var playerTags structures.PlayerTags

	urlStr:="https://api.clashroyale.com/v1/locations/" + strconv.Itoa(id) + "/rankings/players"


	req,err:=NewGetRequest(urlStr)

	//fail to parse url
	if err!=nil {
		return playerTags,err
	}

	SetHeaders(req)

	resp,err:=c.client.Do(req)

	//fail to parse header,no header,timeout
	if err!=nil {
		return playerTags,err
	}

	json.NewDecoder(resp.Body).Decode(&playerTags)

	return playerTags,nil

}