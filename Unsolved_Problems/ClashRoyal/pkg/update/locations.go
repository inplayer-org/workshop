package update

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"strconv"
	"time"
	"log"
)

//UpdateLocations if the location exists in database then it updates it if it doeesnt then it inserts it
func Locations(db *sql.DB,locs structures.Locations)error {

	var err error
	done := make(chan error)
	defer close(done)
	responsesCount := 0

	for _, elem := range locs.Location {
		responsesCount++
		go CurrentLocation(db,elem,done)
		time.Sleep(time.Millisecond*15)
	}

	for ; responsesCount > 0; responsesCount-- {
		err = <-done

		if err!=nil{
			return err
		}

	}

	return nil
}

func CurrentLocation(db *sql.DB,elem structures.Locationsinfo,done chan <- error){
	var err error
	for {
		if !queries.Exists(db, "locations", "id", strconv.Itoa(elem.ID)) {
			err = queries.InsertIntoLocationsTable(db, elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)
		} else {
			err = queries.UpdateLocationsTable(db, elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)
		}
		if err==nil{
			break
		}
		fmt.Println("Error for ",elem.Name," -> ",err)
	}
		log.Println("Finished for -> ",elem.Name)
		done<-err
	}

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

//DailyUpdateLocation makes request and updates the location table
func DailyUpdate(db *sql.DB)(structures.Locations,error){

	locations,err:=Get()

	if err!=nil{
		return locations,err
	}

	err=Locations(db,locations)

	if err!=nil{
		return locations,err
	}

	return locations,nil

}