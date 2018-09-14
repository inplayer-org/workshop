package locations

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"


)

//insertInto inserts location into database
func InsertIntoLocationsTable(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("INSERT INTO locations(id,countryName,isCountry,countryCode) VALUES ((?),(?),(?),(?));", id, name, isCountry, code)

	if err != nil {
		return err
	}

	return nil

}

//update for an id updates a location
func UpdateLocationsTable(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("UPDATE locations SET countryName=(?),isCountry=(?),countryCode=(?) WHERE id=(?)", name, isCountry, code, id)

	if err != nil {
		return err
	}

	return nil

}
// Returning slice of Locations Info from DB Table locations ALLInfo about Location
func GetAllLocations(db *sql.DB)([]Locationsinfo,error){

	rows, _ := db.Query("SELECT id,countryName,isCountry,countryCode from locations")


	defer rows.Close()

	return locationrows(rows)
}

func locationrows (rows *sql.Rows)([]Locationsinfo,error){
	var location  []Locationsinfo

	for rows.Next() {
		var l Locationsinfo
		err:=rows.Scan(&l.ID,&l.Name,&l.IsCountry,&l.CountryCode)

		if err!=nil {
			return nil,err
		}

		location=append(location,l)
	}

	return location,nil
}