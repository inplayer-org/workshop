package structures

//Structure for handling information from Foods table
type Food struct{
	FoodID int `json:"foodid"`
	Type string `json:"type"`
	Name string `json:"name"`
}
