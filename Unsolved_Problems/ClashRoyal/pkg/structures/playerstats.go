package structures

import "database/sql"

type PlayerStats struct {
	Tag    string `json:"tag"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Trophies int `json:"trophies"`
	Clan Clan  `json:"clan"`
	Tags *PlayerStats `"json:"tags"`
	Names *PlayerStats `"json:"names"`
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

/*
func (p *PlayerStats) GetTagPlayer(db *sql.DB) error {

	err := db.QueryRow("SELECT playerTag,playerName from players where playerTag='%s'",p.Tag).Scan(&p.Tag,&p.Name)
	if err!=nil {
		return err
	}
	s,err:=SelectPlayerByTag(db, p.Tag)

	if err!=nil {
		return err
	}

	p.Tags=&s

	return nil
}

func SelectPlayerByTag(db *sql.DB,tag string) (PlayerStats,error){
	var p PlayerStats

	err := db.QueryRow("SELECT playerTag,playerName from players where playertag='%s'",tag).Scan(&p.Tag,&p.Name)
	if err!=nil {
		return p,err
	}

	return p,nil
}


*/


func (p *PlayerStats) GetNamePlayer(db *sql.DB) error {

	err := db.QueryRow("SELECT playerName,playerTag from players where playerName='%s'like ",p.Name).Scan(&p.Name,&p.Tag)
	if err!=nil {
		return err
	}

	return nil
}

