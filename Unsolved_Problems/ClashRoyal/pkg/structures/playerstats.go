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


