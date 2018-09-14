package rankedPlayer


import (
"database/sql"
"fmt"
	"strconv"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/playerStats"
)



func GetSortedRankedPlayers(DB *sql.DB, orderBy string, numberOfPlayers int) ([]RankedPlayer, error) {

	var Players []RankedPlayer
	expression := "SELECT playerTag,playerName,wins,losses,trophies,clans.clanName,rankedPlayer.clanTag from rankedPlayer JOIN clans where rankedPlayer.clanTag=clans.clanTag order by " + orderBy + " desc limit " + strconv.Itoa(numberOfPlayers)
	rows, err := DB.Query(expression)

	if err != nil {
		return nil, err
	}

	rank := 1

	log.Println(rows.Columns())

	for rows.Next() {

		var currentPlayer playerStats.PlayerStats
		err = rows.Scan(&currentPlayer.Tag, &currentPlayer.Name, &currentPlayer.Wins, &currentPlayer.Losses, &currentPlayer.Trophies, &currentPlayer.Clan.Name, &currentPlayer.Clan.Tag)

		if err != nil {
			return nil, err
		}

		currentPlayer.Tag = parser.ToRawTag(currentPlayer.Tag)
		currentPlayer.Clan.Tag = parser.ToRawTag(currentPlayer.Clan.Tag)

		Players = append(Players, RankedPlayer{Player: currentPlayer, Rank: rank})
		rank++
	}

	return Players, nil
}

// Returning and sorting rankedPlayer from 1 location by wins from DB table clans
func GetPlayersByLocation(db *sql.DB, name int) ([]RankedPlayer, error) {

	var player []RankedPlayer
	rows, err := db.Query("SELECT playerName,wins,losses,trophies,clanTag from rankedPlayer where locationID=? order by wins desc limit 200", name)

	if err != nil {
		return nil, err
	}

	rank := 1
	for rows.Next() {

		var t RankedPlayer
		t.Rank = rank

		rows.Scan(&t.Player.Name, &t.Player.Wins, &t.Player.Losses, &t.Player.Trophies, &t.Player.Clan.Tag)
		err := db.QueryRow("SELECT clanName from clans where clanTag=?", t.Player.Clan.Tag).Scan(&t.Player.Clan.Name)

		if err != nil {
			return nil, err
		}

		player = append(player, t)
		rank++
	}
	return player, nil
}

//feature
/*func ClanNotFoundByName(db *sql.DB,name string)(structures.PlayerStats,error){

	var p structures.PlayerStats

	err:=db.QueryRow("SELECT playerTag,playerName,wins,losses,trophies from rankedPlayer where playername=?",name).Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies)

	if err!=nil {
		return p,err
	}

	p.Clan.Name=""

	return p,nil
}*/


//Slice of RankedPlayer returning all rankedPlayer from 1 clan sorted by wins
func GetPlayersByClanTag(db *sql.DB, clanTag string) ([]RankedPlayer, error) {

	var playerss []RankedPlayer

	rows, err := db.Query("SELECT rankedPlayer.playerTag,rankedPlayer.playerName,rankedPlayer.wins,rankedPlayer.losses,rankedPlayer.trophies,rankedPlayer.clanTag,clans.clanName from rankedPlayer join clans where clans.clanTag=rankedPlayer.clanTag and rankedPlayer.clanTag=? order by wins desc limit 50", clanTag)

	if err != nil {
		return nil, err
	}

	rank := 1

	for rows.Next() {

		var player RankedPlayer
		player.Rank = rank

		err = rows.Scan(&player.Player.Tag, &player.Player.Name, &player.Player.Wins, &player.Player.Losses, &player.Player.Trophies, &player.Player.Clan.Tag, &player.Player.Clan.Name)

		if err != nil {
			return nil, err
		}

		player.Player.Tag = parser.ToRawTag(player.Player.Tag)
		player.Player.Clan.Tag = parser.ToRawTag(player.Player.Clan.Tag)

		playerss = append(playerss, player)
		rank++
	}

	return playerss, nil
}
