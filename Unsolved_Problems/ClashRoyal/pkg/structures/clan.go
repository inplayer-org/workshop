package structures

import "database/sql"

type Clan struct {
	//Structure with information for the Clans
	Tag          string    `json:"tag"`
	Name        string `json:"name"`

}

func (c *Clan) GetTagClan(db *sql.DB) error {

	err := db.QueryRow("SELECT clanTag,clanName from clans where clanTag='%s'",c.Tag).Scan(&c.Tag,&c.Name)
	if err!=nil {
		return err
	}

	return nil
}


func (c *Clan) GetNameClan(db *sql.DB) error {

	err := db.QueryRow("SELECT clanTag from clans where clanName=(?)",c.Name).Scan(&c.Tag)
	if err!=nil {
		return err
	}

	return nil
}

