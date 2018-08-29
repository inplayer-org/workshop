package update

import (
	"database/sql"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"strconv"
	"time"
)

//UpdateLocations if the location exists in database then it updates it if it doeesnt then it inserts it
func UpdateLocations(db *sql.DB,locs structures.Locations)error {
	var err error
	done := make(chan error)
	defer close(done)
	responsesCount := 0
	for _, elem := range locs.Location {
		responsesCount++
		go UpdateCurrentLocation(db,elem,done)
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

	func UpdateCurrentLocation(db *sql.DB,elem structures.Location,done chan <- error){
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
	//log.Println("Finished updating for location ->",elem)
