package queries

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"
)


func update(DB *sql.DB,player structures.PlayerStats,locationID interface{},clanTag interface{})error{

		_,err := DB.Exec("update players SET playerName=(?),wins=(?),losses=(?),trophies=(?),clanTag=(?),locationID=(?) where playerTag=(?);",
			player.Name, player.Wins, player.Losses, player.Trophies, clanTag, locationID,player.Tag)
		UpdateError(err)
		return err

}

func insert(DB *sql.DB,player structures.PlayerStats,locationID interface{},clanTag interface{}) error{

		_, err := DB.Exec(`INSERT INTO players(playerTag,playerName,wins,losses,trophies,clanTag,locationID) values((?),(?),(?),(?),(?),(?),(?));`,
			player.Tag, player.Name, player.Wins, player.Losses, player.Trophies, clanTag, locationID)
		InsertError(err)
		return err
}

func UpdatePlayer(DB *sql.DB,player structures.PlayerStats,locationID int)error{

	var clanTag interface{}
	var locID interface{}

	clanTag = nil
	locID = nil

	//log.Println("Updating for player ",player)
	if !(player.Clan.Tag=="") && !(player.Clan.Name==""){
		UpdateClans(DB,structures.Clan{Tag:player.Clan.Tag,Name:player.Clan.Name})
		clanTag = player.Clan.Tag
	}

	if locationID!=0{
		locID = locationID
	}


	if Exists(DB,PlayersTable,PlayerTag,player.Tag){
		return update(DB,player,locID,clanTag)
	}else {
		return insert(DB,player,locID,clanTag)
	}
}
