package update

import (
	"database/sql"
	"strconv"
	"sync"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)

var wg sync.WaitGroup

//UpdateLocations if the location exists in database then it updates it if it doesn't then it inserts it
func Locations(db *sql.DB, locs structures.Locations) error {

	done := make(chan error,300)

	for _, elem := range locs.Location {
		wg.Add(1)
		go CurrentLocation(db, elem, done)
	}
	wg.Wait()

	close(done)

	for err := range done {

		if err != nil {
			return err
		}
	}

	return nil
}

func CurrentLocation(db *sql.DB, elem structures.Locationsinfo, done chan<- error) {
	var err error
	defer wg.Done()

	for {
		if !queries.Exists(db, "locations", "id", strconv.Itoa(elem.ID)) {
			err = queries.InsertIntoLocationsTable(db, elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)
		} else {
			err = queries.UpdateLocationsTable(db, elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)
		}

		if err==nil{
			break
		}

		//Not printed since it's retried
		//if err!=nil{
		//	log.Println(err)
		//}
	}
	done <- err

}

//DailyUpdateLocation makes request and updates the location table
func DailyUpdate(db *sql.DB) (structures.Locations, error) {
	client := _interface.NewClient()
	locations, err := client.GetLocations()

	if err != nil {
		return locations, err
	}

	err = Locations(db, locations)

	if err != nil {
		return locations, err
	}

	return locations, nil

}
