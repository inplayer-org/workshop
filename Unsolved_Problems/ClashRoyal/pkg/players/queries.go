package players


import (
"database/sql"
"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/clans"
	"strconv"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

func Exists(DB *sql.DB,table string,column string,value string) bool{

	var result int

	query:=fmt.Sprintf(`SELECT COUNT(%s) FROM %s WHERE %s="%s"`,column,table,column,value)

	DB.QueryRow(query).Scan(&result)

	if result==0{
		return false
	}

	return true

}



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
		if Exists(DB,"players","playerTag",player.Tag){
			return update(DB,player,locationID,clanTag)
		}else {
			return insert(DB,player,locationID,clanTag)
		}
	} else {
		if Exists(DB, "players", "playerTag", player.Tag) {
			return update(DB, player, nil, clanTag)
		} else {
			return insert(DB, player, nil, clanTag)
		}
	}

}

func GetSortedRankedPlayers(DB *sql.DB, orderBy string, numberOfPlayers int) ([]RankedPlayer, error) {

	var Players []RankedPlayer
	expression := "SELECT playerTag,playerName,wins,losses,trophies,clans.clanName,players.clanTag from players JOIN clans where players.clanTag=clans.clanTag order by " + orderBy + " desc limit " + strconv.Itoa(numberOfPlayers)
	rows, err := DB.Query(expression)

	if err != nil {
		return nil, err
	}

	rank := 1

	log.Println(rows.Columns())

	for rows.Next() {

		var currentPlayer PlayerStats
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

// Returning and sorting players from 1 location by wins from DB table clans
func GetPlayersByLocation(db *sql.DB, name int) ([]RankedPlayer, error) {

	var player []RankedPlayer
	rows, err := db.Query("SELECT playerName,wins,losses,trophies,clanTag from players where locationID=? order by wins desc limit 200", name)

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

	err:=db.QueryRow("SELECT playerTag,playerName,wins,losses,trophies from players where playername=?",name).Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies)

	if err!=nil {
		return p,err
	}

	p.Clan.Name=""

	return p,nil
}*/
// Enterning clanTag as string and returning error if clan is not found from table players (If player dont have clan cant be listed)
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

//Returning  Slice of all players with same name
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

// Geting Player Tag with enterning player Name
func GetPlayerName(db *sql.DB, tag string) (string, error) {

	var name string

	err := db.QueryRow("SELECT playerName FROM players WHERE playerTag=?", tag).Scan(&name)

	if err!= nil{
		return name,err
	}

	return name,nil
}

//Slice of RankedPlayer returning all players from 1 clan sorted by wins
func GetPlayersByClanTag(db *sql.DB, clanTag string) ([]RankedPlayer, error) {

	var playerss []RankedPlayer

	rows, err := db.Query("SELECT players.playerTag,players.playerName,players.wins,players.losses,players.trophies,players.clanTag,clans.clanName from players join clans where clans.clanTag=players.clanTag and players.clanTag=? order by wins desc limit 50", clanTag)

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
