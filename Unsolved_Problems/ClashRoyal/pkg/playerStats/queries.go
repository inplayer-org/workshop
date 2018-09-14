package playerStats

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/clans"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

func update(DB *sql.DB,player PlayerStats,locationID interface{},clanTag interface{})error{
	if locationID!=nil {
		_, err := DB.Exec("update players SET playerName=(?),wins=(?),losses=(?),trophies=(?),clanTag=(?),locationID=(?) where playerTag=(?);",
			player.Name, player.Wins, player.Losses, player.Trophies, clanTag, locationID, player.Tag)
		return err
	}else{
		_, err := DB.Exec("update players SET playerName=(?),wins=(?),losses=(?),trophies=(?),clanTag=(?) where playerTag=(?);",
			player.Name, player.Wins, player.Losses, player.Trophies, clanTag, player.Tag)
		return err
	}
}

func insert(DB *sql.DB, player PlayerStats, locationID interface{}, clanTag interface{}) error {

	_, err := DB.Exec(`INSERT INTO players(playerTag,playerName,wins,losses,trophies,clanTag,locationID) values((?),(?),(?),(?),(?),(?),(?));`,
		player.Tag, player.Name, player.Wins, player.Losses, player.Trophies, clanTag, locationID)

	return err
}

func UpdatePlayer(DB *sql.DB, player PlayerStats, locationID interface{}) error {

	var clanTag interface{}

	clanTag = nil


	if !(player.Clan.Tag=="") && !(player.Clan.Name==""){

		err:=clans.UpdateClans(DB,clans.Clan{Tag:player.Clan.Tag,Name:player.Clan.Name})
		if err!=nil {
			return err
		}

		clanTag = player.Clan.Tag
	}

	if locationID!=nil{
		if parser.Exists(DB,"players","playerTag",player.Tag){
			return update(DB,player,locationID,clanTag)
		}else {
			return insert(DB,player,locationID,clanTag)
		}
	} else {
		if parser.Exists(DB, "players", "playerTag", player.Tag) {
			return update(DB, player, nil, clanTag)
		} else {
			return insert(DB, player, nil, clanTag)
		}
	}

}

// Enterning clanTag as string and returning error if clan is not found from table rankedPlayer (If player dont have clan cant be listed)
func ClanNotFoundByTag(db *sql.DB, tag string) (PlayerStats, error) {

	var p PlayerStats

	err := db.QueryRow("SELECT playerTag,playerName,wins,losses,trophies from players where playerTag=?", tag).Scan(&p.Tag, &p.Name, &p.Wins, &p.Losses, &p.Trophies)

	if err != nil {
		return p, err
	}

	p.Tag = parser.ToRawTag(tag)
	p.Clan.Name = ""

	return p, nil
}

// PlayerTag String and returning all informations for 1 player from PLayerstats with clan information Joining Clans into Players table.
func GetFromTag(db *sql.DB, tag string) (PlayerStats, error) {

	var p PlayerStats

	t:=parser.ToRawTag(tag)

	err:=db.QueryRow("SELECT players.playerTag,players.playerName,players.wins,players.losses,players.trophies,players.clanTag,clans.clanName From players inner join clans where players.clanTag=clans.clanTag and players.playerTag=?",tag).Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies,&p.Clan.Tag,&p.Clan.Name)
	p.Tag=t
	if err!=nil{
		return p,err
	}
	p.Clan.Tag = parser.ToRawTag(p.Clan.Tag)
	return p, nil
}

//Returning  Slice of all rankedPlayer with same name
func GetPlayersLike(db *sql.DB, name string) ([]PlayerStats, error) {

	var player []PlayerStats

	rows, err := db.Query("SELECT playerTag,playerName,wins,losses,trophies,clanTag FROM players WHERE playerName Like (?)", "%"+name+"%")
	if err!=nil{
		return nil,err
	}
	for rows.Next() {
		var temp interface{}
		var p PlayerStats
		err = rows.Scan(&p.Tag, &p.Name, &p.Wins, &p.Losses, &p.Trophies,&temp)
		if err != nil {
			return nil, err
		}

		if temp!=nil{
			var s []uint8
			s = temp.([]uint8)
			p.Clan.Tag = string(s)
			p.Clan.Name,err = clans.GetClanName(db,p.Clan.Tag)
			p.Clan.Tag = parser.ToRawTag(p.Clan.Tag)
		}


		p.Tag = parser.ToRawTag(p.Tag)
		player = append(player, p)
	}

	return player, nil
}
