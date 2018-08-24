package locations

import (
"net/http"
"encoding/json"
)

//GetLocations gets the locations and returns error when cant make request or client cant do the request
func Get()(structures.Locations,error){

	var locations structures.Locations

	client := &http.Client{}

	req,err :=http.NewRequest("GET","https://api.clashroyale.com/v1/locations",nil)

	if err!=nil{
		return locations,err
	}

	req.Header.Add("Content-Type","application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")

	resp,err:=client.Do(req)

	if err!=nil{
		return locations,err
	}

	json.NewDecoder(resp.Body).Decode(&locations)

	return locations,nil

}


//LocationMap convert all countries from locations into map with key country name and value country id
func ToMap(locations structures.Locations)map[string]int{

	locationMap:=make(map[string]int)

	for _,location:= range locations.Location{
		if location.IsCountry {
			locationMap[location.Name]=location.ID
		}
	}

	return locationMap

}

//FindLocationID returns the id of location or 0 if its not country
func Find(locationMap map[string]int,country string)int{

	return locationMap[country]

}