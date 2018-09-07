package structures

//Structure for handling information from Animals table
type Eat struct {
	Animal Animal `json:"animal"`
	Food   []Food `json:"food"`
}
