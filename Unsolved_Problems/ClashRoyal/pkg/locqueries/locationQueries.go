package locqueries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	_ "github.com/go-sql-driver/mysql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"strconv"
)

func UpdateLocations(db *sql.DB,locs locations.Locations)error{

	for _,elem:=range locs.Location {

		if queries.Exists(db, "locations", "id", strconv.Itoa(elem.ID)) {

			_, err := db.Exec("INSERT INTO locations(id,countryName,isCountry,countryCode) VALUES ((?).(?),(?),(?));", elem.ID, elem.Name, elem.IsCountry, elem.CountryCode)

			if err != nil {
				return err
			}

		} else {

			_, err := db.Exec("UPDATE locations SET countryName=(?),isCounrty=(?),countryCode=(?) WHERE id=(?)", elem.Name, elem.IsCountry, elem.CountryCode, elem.ID)

			if err != nil {
				return err
			}

		}
	}

	return nil

}

func GetLocationID(db *sql.DB,name string)(int,error){

	var id int

	err:=db.QueryRow("SELECT id FROM locations WHERE countryName=(?);",name).Scan(&id)

	if err!= nil {
		return 0,nil
	}

	return id ,nil

}
