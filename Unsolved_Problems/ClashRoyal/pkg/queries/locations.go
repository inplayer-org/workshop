package queries

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
)

//insertInto inserts location into database
func InsertIntoLocationsTable(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("INSERT INTO locations(id,countryName,isCountry,countryCode) VALUES ((?),(?),(?),(?));", id, name, isCountry, code)

	if err != nil {
		return errors.Database(err)
	}

	return nil

}

//update for an id updates a location
func UpdateLocationsTable(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("UPDATE locations SET countryName=(?),isCountry=(?),countryCode=(?) WHERE id=(?)", name, isCountry, code, id)

	if err != nil {
		return errors.Database(err)
	}

	return nil

}
// Returning slice of Locations Info from DB Table locations ALLInfo about Location
func GetAllLocations(db *sql.DB)([]structures.Locationsinfo,error){

	rows, _ := db.Query("SELECT id,countryName,isCountry,countryCode from locations")


	defer rows.Close()

	return locationrows(rows)
}

func locationrows (rows *sql.Rows)([]structures.Locationsinfo,error){
	var locations  []structures.Locationsinfo

	for rows.Next() {
		var l structures.Locationsinfo
		err:=rows.Scan(&l.ID,&l.Name,&l.IsCountry,&l.CountryCode)

		if err!=nil {
			return nil,errors.Database(err)
		}

		locations=append(locations,l)
	}

	return locations,nil
}