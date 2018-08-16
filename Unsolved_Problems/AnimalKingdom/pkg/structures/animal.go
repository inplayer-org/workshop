package structures

//Structure for handling information from Animals table
type Animal struct {
	AnimalID int    `json:"animalid"`
	Name     string `json:"name"`
	Species  string `json:"species"`
	Height   int    `json:"height"`
}
