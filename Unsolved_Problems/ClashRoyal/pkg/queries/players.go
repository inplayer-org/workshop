package queries

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"
)


func update(DB *sql.DB,player structures.PlayerStats,locationID int)error{
	if locationID==0{
		_,err := DB.Exec("update players SET playerName=(?),wins=(?),losses=(?),trophies=(?),clanTag=(?) where playerTag=(?);",
			player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag, player.Tag)
		UpdateError(err)
		return err
	}else{
		_,err := DB.Exec("update players SET playerName=(?),wins=(?),losses=(?),trophies=(?),clanTag=(?),locationID=(?) where playerTag=(?);",
			player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag, locationID,player.Tag)
		UpdateError(err)
		return err
	}
}

func insert(DB *sql.DB,player structures.PlayerStats,locationID int) error{
	if locationID == 0{
		_,err := DB.Exec(`INSERT INTO players(playerTag,playerName,wins,losses,trophies,clanTag) values((?),(?),(?),(?),(?),(?));`,
			player.Tag, player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag)
		InsertError(err)
		return err
	}else {
		_, err := DB.Exec(`INSERT INTO players(playerTag,playerName,wins,losses,trophies,clanTag,locationID) values((?),(?),(?),(?),(?),(?),(?));`,
			player.Tag, player.Name, player.Wins, player.Losses, player.Trophies, player.Clan.Tag, locationID)
		InsertError(err)
		return err
	}
}

func UpdatePlayer(DB *sql.DB,player structures.PlayerStats,locationID int)error{

	//log.Println("Updating for player ",player)
	if !(player.Clan.Tag=="") && !(player.Clan.Name==""){
		UpdateClans(DB,structures.Clan{Tag:player.Clan.Tag,Name:player.Clan.Name})
	}
	if Exists(DB,PlayersTable,PlayerTag,player.Tag){
		return update(DB,player,locationID)
	}else {
		return insert(DB,player,locationID)
	}
}
