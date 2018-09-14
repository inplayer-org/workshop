package playerStats

import "repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/clans"

type PlayerStats struct {
	Tag        string      `json:"tag"`
	Name       string      `json:"name"`
	Wins       int         `json:"wins"`
	Losses     int         `json:"losses"`
	Trophies   int         `json:"trophies"`
	Clan       clans.Clan        `json:"clan"`
	LocationID interface{} `json:"location_id"`
}
