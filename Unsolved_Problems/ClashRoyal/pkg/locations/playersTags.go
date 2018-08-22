package locations

import (
	"net/http"
	"net/url"
	"encoding/json"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)



func GetPlayerTagsPerLocation(ID int)(structures.PlayerTags,error){

	var playerTags structures.PlayerTags

	client := &http.Client{}

	urlStr:="https://api.clashroyale.com/v1/locations/" + strconv.Itoa(ID) + "/rankings/players"

	link,err:=url.Parse(urlStr)

	if err != nil {
		return playerTags,err
	}

	reqHeader:=http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")

	resp,err:=client.Do(&http.Request{Header:reqHeader,URL:link,Method:"GET"})

	if err!=nil {
		return playerTags,err
	}

	json.NewDecoder(resp.Body).Decode(&playerTags)

	return playerTags,nil
}
