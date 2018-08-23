package queries

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"
)


func update(DB *sql.DB,player structures.PlayerStats,locationID int){
	if locationID==0{
		_,err := DB.Exec(`update players SET playerName="(?)",wins=(?),losses=(?),trophies=(?),clanTag=(?) where playerTag=(?);`,
			player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag, player.Tag)
		UpdateError(err)
	}else{
		_,err := DB.Exec(`update players SET playerName="(?)",wins=(?),losses=(?),trophies=(?),clanTag=(?),locationID=(?) where playerTag=(?);`,
			player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag, locationID, player.Tag)
		UpdateError(err)
	}
}

func insert(DB *sql.DB,player structures.PlayerStats,locationID int){
	_,err := DB.Exec(`INSERT INTO players(playerTag,playerName,wins,losses,trophies,clanTag,locationID) values((?),(?),(?),(?),(?),(?),(?));`,
		player.Tag, player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag, locationID)
	InsertError(err)
}

func UpdatePlayer(DB *sql.DB,player structures.PlayerStats,locationID int){
	if Exists(DB,PlayersTable,PlayerTag,player.Tag){
		update(DB,player,locationID)
	}else {
		insert(DB,player,locationID)
	}
}
