package structures

import (
	"database/sql"
)

type Clan struct {
	//Structure with information for the Clans
	Tag          string    `json:"tag"`
	Name        string `json:"name"`

}


func  SelectClanByTag(db *sql.DB, clantag string) interface{} {
	var clanTag, name string
	err := db.QueryRow("SELECT * FROM clans WHERE clanTag=(?)", clantag).Scan(&clanTag, &name)
	if err!=nil {
		return nil
	}
	return Clan{Tag: clanTag, Name: name}
}

func  SelectClanByName(db *sql.DB, name string) interface{} {
	var Cname, clantag string
	err := db.QueryRow("SELECT * FROM clans WHERE clanName=(?)", Cname).Scan(&Cname, &clantag)
	if err!=nil {
		return nil
	}
	return Clan{Name: Cname, Tag: clantag}
}

