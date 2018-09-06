package structures

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
