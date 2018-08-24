package structures

import (
	"database/sql"
)

type Clan struct {
	//Structure with information for the Clans
	Tag          string    `json:"tag"`
	Name        string `json:"name"`

}



func (c *Clan) GetNameClan(db *sql.DB) error {

	err := db.QueryRow("SELECT clanTag from clans where clanName=(?)",c.Name).Scan(&c.Tag)
	if err!=nil {
		return err
	}

	return nil
}

func GetAllClans(db *sql.DB)([]Clan,error){

	rows, _ := db.Query("SELECT * from clans")


	defer rows.Close()

	return clanRows(rows)
}

func clanRows(rows *sql.Rows)([]Clan,error){
	var clans  []Clan

	for rows.Next() {
		var c Clan
		err:=rows.Scan(&c.Tag,&c.Name)

		if err!=nil {
			return nil,err
		}

		clans=append(clans,c)
	}

	return clans,nil
}
