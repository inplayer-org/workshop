package structures

import "database/sql"

type PlayerStats struct {
	Tag    string `json:"tag"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Trophies int `json:"trophies"`
	Clan Clan  `json:"clan"`
	LocationID interface{} `json:"location_id"`
}


func (p *PlayerStats) GetNamePlayer(db *sql.DB) error {

	err := db.QueryRow("SELECT playerName,playerTag from players where playerName='%s'like",p.Name).Scan(&p.Name,&p.Tag)
	if err!=nil {
		return err
	}

	return nil
}



func GetPlayersLike(db *sql.DB,name string)([]PlayerStats,error){
	var players [] PlayerStats
	rows,err:=db.Query("SELECT playerTag,playerName,wins,losses,trophies,clanTag,locationid FROM players Where playerName Like (?)",name)
	if err !=nil {
		return nil,err
	}

	for rows.Next(){

		var p PlayerStats
		err = rows.Scan(&p.Name,&p.Tag,&p.Wins,&p.Losses,&p.Trophies,&p.Clan.Name,&p.LocationID)

		if err !=nil {
			return nil,err
		}

		players = append(players,p)
	}

	return players,nil

}