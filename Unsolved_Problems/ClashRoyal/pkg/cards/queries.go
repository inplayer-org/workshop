package cards

import "database/sql"

//insertInto inserts location into database
func InsertIntoCardsTable(db *sql.DB,name string,id int,maxlevel int,iconurl string)error{

	_, err := db.Exec("INSERT INTO cards(name,id,maxlevel,iconurl) VALUES ((?),(?),(?),(?));",name,id,maxlevel,iconurl)

	if err != nil {
		return err
	}

	return nil

}

//update for an id updates a location
func UpdateCardsTable(db *sql.DB,name string,id int,maxlevel int,iconurl string)error{

	_, err := db.Exec("UPDATE cards SET name=(?),id=(?),maxlevel=(?) WHERE iconurl=(?)", name,id,maxlevel,iconurl)

	if err != nil {
		return err
	}

	return nil

}
// Returning slice of Locations Info from DB Table locations ALLInfo about Location
func GetAllCards(db *sql.DB)([]Cards,error){

	rows, _ := db.Query("SELECT name,id,maxlevel,iconurl from cards")


	defer rows.Close()

	return cardsrows(rows)
}

func cardsrows (rows *sql.Rows)([]Cards,error){
	var card  []Cards

	for rows.Next() {
		var c Cards
		err:=rows.Scan(&c.Items)
		if err!=nil {
			return nil,err
		}

		card=append(card,c)
	}

	return card,nil
}

