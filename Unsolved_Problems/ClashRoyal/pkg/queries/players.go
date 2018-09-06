package queries

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"strconv"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
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

func GetSortedRankedPlayers(DB *sql.DB,orderBy string,numberOfPlayers int)([]structures.RankedPlayer,error){

		var Players []structures.RankedPlayer
	expression := "SELECT playerTag,playerName,wins,losses,trophies,clans.clanName,players.clanTag from players JOIN clans where players.clanTag=clans.clanTag order by "+ orderBy + " desc limit " + strconv.Itoa(numberOfPlayers)
	rows,err := DB.Query(expression)

	if err!=nil{
		return nil,err
	}

	rank :=1

	log.Println(rows.Columns())

	for rows.Next(){

		var currentPlayer structures.PlayerStats
		err = rows.Scan(&currentPlayer.Tag,&currentPlayer.Name,&currentPlayer.Wins,&currentPlayer.Losses,&currentPlayer.Trophies,&currentPlayer.Clan.Name,&currentPlayer.Clan.Tag)

		if err!=nil{
			return nil,err
		}

		currentPlayer.Tag = parser.ToRawTag(currentPlayer.Tag)
		currentPlayer.Clan.Tag = parser.ToRawTag(currentPlayer.Clan.Tag)

		Players = append(Players,structures.RankedPlayer{Player:currentPlayer,Rank:rank})
		rank++
	}

	return Players,nil
}

func GetPlayersByLocation(db *sql.DB,name int)([]structures.RankedPlayer,error){

	var players []structures.RankedPlayer
	rows,err:=db.Query("SELECT PlayerName,wins,losses,trophies,clanTag from players where locationID=? order by wins desc limit 200",name)

	if err!=nil {
		return nil,err
	}

	rank:=1
	for rows.Next(){

		var t structures.RankedPlayer
		t.Rank=rank

		rows.Scan(&t.Player.Name,&t.Player.Wins,&t.Player.Losses,&t.Player.Trophies,&t.Player.Clan.Tag)
		err:=db.QueryRow("SELECT clanName from clans where clanTag=?",t.Player.Clan.Tag).Scan(&t.Player.Clan.Name)

		if err!=nil {
			return nil,err
		}

		players=append(players,t)
		rank++
	}
	return players,nil
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

func ClanNotFoundByTag(db *sql.DB,tag string)(structures.PlayerStats,error){

	var p structures.PlayerStats

	t:=parser.ToHashTag(tag)

	err:=db.QueryRow("SELECT playerTag,playerName,wins,losses,trophies from players where playerTag=?",t).Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies)

	if err!=nil {
		return p,err
	}

	p.Tag=tag
	p.Clan.Name=""

	return p,nil
}

func GetFromTag(db *sql.DB,tag string)(structures.PlayerStats,error){

	var p structures.PlayerStats

	t:=parser.ToHashTag(tag)

	err:=db.QueryRow("SELECT players.playerTag,players.playerName,players.wins,players.losses,players.trophies,players.clanTag, clans.clanName From players inner join clans where players.clanTag=clans.clanTag and clans.clanTag=players.clanTag and players.playerTag=?",t).Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies,&p.Clan.Tag,&p.Clan.Name)
	p.Tag=tag
	if err!=nil{
		fmt.Println(err)
		return p,err
	}
	p.Clan.Tag = parser.ToRawTag(p.Clan.Tag)
	return p,nil
}

func GetPlayersLike(db *sql.DB,name string)([]structures.PlayerStats,error){

	var players [] structures.PlayerStats
	rows,err:=db.Query("SELECT players.playerTag,players.playerName,players.wins,players.losses,players.trophies,players.clanTag, clans.clanName FROM players inner join clans Where players.clanTag=clans.clanTag and playerName Like (?)","%"+name+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){

		var p structures.PlayerStats
		err = rows.Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies,&p.Clan.Tag,&p.Clan.Name)

		if err !=nil {
			return nil,err
		}

		p.Tag=parser.ToRawTag(p.Tag[1:])
		p.Clan.Tag = parser.ToRawTag(p.Clan.Tag[1:])

		players = append(players,p)
	}

	return players,nil
}

func GetPlayerName(db *sql.DB,tag string)(string,error){

	var name string

	err := db.QueryRow("SELECT playerName FROM players WHERE playerTag=?",tag).Scan(&name)

	return name,err
}

func GetPlayersByClanTag(db *sql.DB,clanTag string)([]structures.RankedPlayer,error){

	clanTag=parser.ToHashTag(clanTag)

	var players []structures.RankedPlayer

	rows,err:=db.Query("SELECT players.playerTag,players.playerName,players.wins,players.losses,players.trophies,players.clanTag,clans.clanName from players join clans where clans.clanTag=players.clanTag and players.clanTag=? order by wins desc limit 50",clanTag)

	if err != nil {
		return nil,err
	}

	rank:=1

	for rows.Next(){

		var player structures.RankedPlayer
		player.Rank=rank

		err=rows.Scan(&player.Player.Tag,&player.Player.Name,&player.Player.Wins,&player.Player.Losses,&player.Player.Trophies,&player.Player.Clan.Tag,&player.Player.Clan.Name)

		if err != nil {
			return nil,err
		}

		player.Player.Tag=parser.ToRawTag(player.Player.Tag)
		player.Player.Clan.Tag=parser.ToRawTag(player.Player.Clan.Tag)

		players=append(players,player)
		rank++
	}

	return players,nil
	}