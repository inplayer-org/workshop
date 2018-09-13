package players
import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/clans"
)

type TwoPlayers struct {
	Player1 PlayerStats
	Player2 PlayerStats
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func (p TwoPlayers) DiffWins() int {
	return Abs(p.Player1.Wins - p.Player2.Wins)
}
func (p TwoPlayers) DiffLosses() int {
	return Abs(p.Player1.Losses - p.Player2.Losses)
}
func (p TwoPlayers) DiffTrophies() int {
	return Abs(p.Player1.Trophies - p.Player2.Trophies)
}





type RankedPlayer struct{
	Rank int
	Player PlayerStats
}



type PlayerStats struct {
	Tag        string      `json:"tag"`
	Name       string      `json:"name"`
	Wins       int         `json:"wins"`
	Losses     int         `json:"losses"`
	Trophies   int         `json:"trophies"`
	Clan       clans.Clan        `json:"clan"`
	LocationID interface{} `json:"location_id"`
}


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
