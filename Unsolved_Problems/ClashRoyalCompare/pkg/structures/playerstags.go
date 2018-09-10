package structures

//PlayerTags - Contains only player Tags for a slice of Player structure
type PlayerTags struct {
	Player []struct {
		Tag          string `json:"tag"`
		/*	Name         string `json:"name"`
			ExpLevel     int    `json:"expLevel"`
			Trophies     int    `json:"trophies"`
			Rank         int    `json:"rank"`
			PreviousRank int    `json:"previousRank"`
			Clan         struct {
				Tag     string `json:"tag"`
				Name    string `json:"name"`
				BadgeID int    `json:"badgeId"`
			} `json:"clan"`
			Arena struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"arena"`*/
	} `json:"items"`
}

//GetTags - Returns string slice of all tags present in a PlayersTags structure
func (p *PlayerTags) GetTags()[]string{

	var ret []string

	for _,elem:=range p.Player {
		ret=append(ret,elem.Tag)
	}

	return ret

}


