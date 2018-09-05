package structures

import "database/sql"

//Locationsinfo - structure for handling location info from Clash Royale api
type Locationsinfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode"`
}

//Structure with a slice of LocationsInfo
type Locations struct {
	Location []Locationsinfo `json:"items"`
}

//Not used, Needs to be removed
func (c *Locationsinfo) GetNameLocation(db *sql.DB) error {

	err := db.QueryRow("SELECT countryName,countryCode from locations where countryName=(?)",c.Name).Scan(&c.Name,&c.CountryCode)

	if err!=nil {
		return err
	}

	return nil
}

