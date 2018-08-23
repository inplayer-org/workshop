package locqueries

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateLocations(db *sql.DB,locs locations.Locations)error{

	for _,elem:=range locs.Location {

		//if not exist
		_,err:=db.Exec("INSERT INTO locations(id,countryName,isCountry,countryCode) VALUES ((?).(?),(?),(?));",elem.ID,elem.Name,elem.IsCountry,elem.CountryCode)

		if err!=nil {
			return err
		}

		//else
		_,err=db.Exec("UPDATE locations SET countryName=(?),isCounrty=(?),countryCode=(?) WHERE id=(?)",elem.Name,elem.IsCountry,elem.CountryCode,elem.ID)

		if err!=nil {
			return err
		}

	}

	return nil

}

