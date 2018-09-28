package playerTags

import "database/sql"

// Geting Player Tag with enterning player Name
func GetPlayerName(db *sql.DB, tag string) (string, error) {

	var name string

	err := db.QueryRow("SELECT playerName FROM players WHERE playerTag=?", tag).Scan(&name)

	if err!= nil{
		return name,err
	}

	return name,nil
}
