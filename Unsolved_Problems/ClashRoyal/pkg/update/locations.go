package update

import (
	"database/sql"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"strconv"
	"time"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
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
		//log.Println("Finished for -> ",elem.Name)
		done<-err
	}

//DailyUpdateLocation makes request and updates the location table
func DailyUpdate(db *sql.DB)(structures.Locations,error){
	client := _interface.NewClient()
	locations,err:=client.GetLocations()

	if err!=nil{
		return locations,err
	}

	err=Locations(db,locations)

	if err!=nil{
		return locations,err
	}

	return locations,nil

}

