package queries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	_ "github.com/go-sql-driver/mysql"
)



//insertInto inserts location into database
func InsertIntoLocationsTable(db *sql.DB,id int,name string,isCountry bool,code string)error{

	_, err := db.Exec("INSERT INTO locations(id,countryName,isCountry,countryCode) VALUES ((?).(?),(?),(?));", id, name, isCountry, code)

	if err != nil {
		return err
	}

	return nil

}


//update for an id updates a location
func UpdateLocationsTable(db *sql.DB,id int,name string,isCountry bool,code string)error{

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

//GetAllLocations returns all locations from locations table in database
func GetAllLocations(db *sql.DB)(locations.Locations,error){

	var locs locations.Locations

	rows,err:=db.Query("SELECT (id,countryName,isCountry,countryCode) FROM locations;")

	if err !=nil {
		return locs,err
	}

	i:=0
	for rows.Next(){
		err:=rows.Scan(&locs.Location[i].ID,&locs.Location[i].Name,&locs.Location[i].IsCountry,&locs.Location[i].CountryCode)

		if err!=nil {
			return locs,err
		}

		i++

		}

	return locs,nil
}
