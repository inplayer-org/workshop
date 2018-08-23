package get

import (
	"net/http"
	"encoding/json"
)

type Clans struct {
	//Structure with information for the Clans
	Clans []struct {
		Tag          string    `json:"tag"`
		Name        string `json:"name"`

	} `json:"items"`
}

func GetClans()(Clans,error){

var clans Clans

client := &http.Client{}

req,err :=http.NewRequest("GET","https://api.clashroyale.com/v1/clans/"+ clanTag +"",nil)

if err!=nil{
return clans,err
}

req.Header.Add("Content-Type","application/json")
req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")

resp,err:=client.Do(req)

if err!=nil{
return clans,err
}

json.NewDecoder(resp.Body).Decode(&clans)

return clans,nil

}
