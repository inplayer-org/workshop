package queries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

//UpdateLocations if the location exists in database then it updates it if it doeesnt then it inserts it
func UpdateLocations(db *sql.DB,locs locations.Locations)error{

	for _,elem:=range locs.Location {

		if Exists(db, "locations", "id", strconv.Itoa(elem.ID)) {

			err := insertIntoDB(db,elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)

			if err != nil {
				return err
			}

		} else {

			err := updateDB(db,elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)

			if err != nil {
				return err
			}

		}
	}

	return nil

}

//insertInto inserts location into database
func insertIntoDB(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("INSERT INTO locations(id,countryName,isCountry,countryCode) VALUES ((?).(?),(?),(?));", id, name, isCountry, code)

	if err != nil {
		return err
	}

	return nil

}


//update for an id updates a location
func updateDB(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("UPDATE locations SET countryName=(?),isCounrty=(?),countryCode=(?) WHERE id=(?)", name, isCountry, code, id)

	if err != nil {
		return err
	}

	return nil

}

//GetLocationID for location returns its id
func GetLocationID(db *sql.DB,name string)(int,error){

	var id int

	err:=db.QueryRow("SELECT id FROM locations WHERE countryName=(?);",name).Scan(&id)

	if err!= nil {
		return 0,nil
	}

	return id ,nil

}
