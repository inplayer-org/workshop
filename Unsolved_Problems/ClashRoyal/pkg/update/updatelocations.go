package update

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"database/sql"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)

//UpdateLocations if the location exists in database then it updates it if it doeesnt then it inserts it
func UpdateLocations(db *sql.DB,locs structures.Locations)error{

	for _,elem:=range locs.Location {

		if !queries.Exists(db, "locations", "id", strconv.Itoa(elem.ID)) {

			err := queries.InsertIntoLocationsTable(db,elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)

			if err != nil {
				return err
			}

		} else {

			err := queries.UpdateLocationsTable(db,elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)

			if err != nil {
				return err
			}

		}
		//log.Println("Finished updating for location ->",elem)
	}

	return nil

}