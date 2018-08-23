package structures

type PlayerStats struct {
	Tag    string `json:"tag"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Trophies string `json:"trophies"`
	Clan Clan  `json:"clan"`
}

type ByWins []PlayerStats

func (p ByWins) Len() int {
	return len(p)
}
func (p ByWins) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ByWins) Less(i, j int) bool {
	if p[i].Wins!=p[j].Wins {
		return p[i].Wins>p[j].Wins
	}
	if p[i].Losses!=p[j].Losses {
		return p[i].Losses<p[j].Losses
	}
	return p[i].Name<p[j].Name
}
