package structures

//Structure with a slice of Location
type Locations struct {
	//Structure with information for the Location
	Location []Location `json:"items"`
}

type Location struct{
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode"`
}