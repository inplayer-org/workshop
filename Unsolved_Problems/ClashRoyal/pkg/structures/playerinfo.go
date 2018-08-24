package structures

import "database/sql"

type PlayerInfo struct {
	Tag    string `json:"tag"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Trophies int `json:"trophies"`
	Clan string  `json:"clan"`
	LocationID int `json:"location_id"`
}



func (p *PlayerInfo) GetNamePlayer(db *sql.DB) error {

	err := db.QueryRow("SELECT playerName,playerTag from players where playerName='%s'like ",p.Name).Scan(&p.Name,&p.Tag)
	if err!=nil {
		return err
	}

	return nil
}



func GetAllPlayers(db *sql.DB)([]PlayerInfo,error){

	rows, err:= db.Query("SELECT playerTag,playerName,wins,losses,trophies,clanTag,locationID from players")

	if err!=nil{

		return nil,err
	}

	defer rows.Close()

	return playerRows(rows)
}

func playerRows(rows *sql.Rows)([]PlayerInfo,error){
	var players  []PlayerInfo

	for rows.Next() {
		var p PlayerInfo
		err:=rows.Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies,&p.Clan,&p.LocationID)

		if err!=nil {

			return nil,err
		}

		players=append(players,p)
	}

	return players,nil
}
