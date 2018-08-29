package structures

import "database/sql"

type Locationsinfo struct {


		ID          int    `json:"id"`
		Name        string `json:"name"`
		IsCountry   bool   `json:"isCountry"`
		CountryCode string `json:"countryCode"`
	}

func (c *Locationsinfo) GetNameLocation(db *sql.DB) error {

	err := db.QueryRow("SELECT countryName,countryCode from locations where countryName=(?)",c.Name).Scan(&c.Name,&c.CountryCode)
	if err!=nil {
		return err
	}

	return nil
}



func GetAllLocations(db *sql.DB)([]Locationsinfo,error){

	rows, _ := db.Query("SELECT id,countryName,isCountry,countryCode from locations")


	defer rows.Close()

	return locationrows(rows)
}

func locationrows(rows *sql.Rows)([]Locationsinfo,error){
	var locations  []Locationsinfo

	for rows.Next() {
		var l Locationsinfo
		err:=rows.Scan(&l.ID,&l.Name,&l.IsCountry,&l.CountryCode)

		if err!=nil {
			return nil,err
		}

		locations=append(locations,l)
	}

	return locations,nil
}
