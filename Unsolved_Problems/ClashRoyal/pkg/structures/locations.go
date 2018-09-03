package structures

import "database/sql"

type Locationsinfo struct {


	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode"`
}

//Structure with a slice of Location
type Locations struct {
	//Structure with information for the Location
	Location []Locationsinfo `json:"items"`

}


func (c *Locationsinfo) GetNameLocation(db *sql.DB) error {

	err := db.QueryRow("SELECT countryName,countryCode from locations where countryName=(?)",c.Name).Scan(&c.Name,&c.CountryCode)
	if err!=nil {
		return err
	}

	return nil
}