package update

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
)


func UpdateClans(db *sql.DB,clan []structures.Clan)error{

for _,elem:=range clan {

	//if not exist
	_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES ((?),(?))", elem.Tag, elem.Name)

	if err != nil {
		return err
	}

}
 _,err := db.Exec("UPDATE clans SET clanTag=(?), clanName=(?) WHERE clanTag=(?)")
	if err != nil {
		return err
	}

 return err
 }

//UpdateLocations if the location exists in database then it updates it if it doeesnt then it inserts it
func UpdateLocations(db *sql.DB,locs locations.Locations)error{

	for _,elem:=range locs.Location {

		if queries.Exists(db, "locations", "id", strconv.Itoa(elem.ID)) {

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
	}

	return nil

}