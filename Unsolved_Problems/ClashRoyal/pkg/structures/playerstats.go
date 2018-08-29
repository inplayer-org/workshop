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

type ByWins []PlayerStats

func (p ByWins) Len() int {
	return len(p)
}
func (p ByWins) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ByWins) Less(i, j int) bool {
	if p[i].Wins!=p[j].Wins {
		return p[i].Wins>p[j].Wins
	}
	if p[i].Losses!=p[j].Losses {
		return p[i].Losses<p[j].Losses
	}
	return p[i].Name<p[j].Name
}



/*func (p *PlayerStats) GetTop100()([]PlayerStats,error){

}*/



func (p *PlayerStats) GetNamePlayer(db *sql.DB) error {

	err := db.QueryRow("SELECT playerName,playerTag from players where playerName='%s'like ",p.Name).Scan(&p.Name,&p.Tag)
	if err!=nil {
		return err
	}

	return nil
}



func  GetAllPlayers(db *sql.DB)([]PlayerStats,error){


	rows, err:= db.Query("SELECT playerTag,playerName,wins,losses,trophies,clanTag,locationID from players")

	if err!=nil{

		return nil,err
	}

	defer rows.Close()

	return playerRows(db,rows)
}

func playerRows(db *sql.DB,rows *sql.Rows)([]PlayerStats,error){
	var players  []PlayerStats
	for rows.Next() {
		var p PlayerStats
		err:=rows.Scan(&p.Tag,&p.Name,&p.Wins,&p.Losses,&p.Trophies,&p.Clan.Tag,&p.LocationID)

		if err!=nil {

			return nil,err
		}

		err=db.QueryRow("select clanName from clans where clanTag=?",p.Clan.Tag).Scan(&p.Clan.Name)

		if err!=nil {

			return nil,err
		}

		players=append(players,p)
	}

	return players,nil
}