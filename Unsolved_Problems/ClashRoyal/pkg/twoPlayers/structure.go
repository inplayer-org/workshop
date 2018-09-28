package twoPlayers

import "repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/playerStats"

type TwoPlayers struct {
	Player1 playerStats.PlayerStats
	Player2 playerStats.PlayerStats
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
