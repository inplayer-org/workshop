package cards

import (
	"database/sql"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
)

//insertInto inserts location into database
func InsertIntoCardsTable(db *sql.DB, name string, id int, maxlevel int, iconurl string) error {

	_, err := db.Exec("INSERT INTO cards(name,id,maxlevel,iconurl) VALUES ((?),(?),(?),(?));", name, id, maxlevel, iconurl)

	if err != nil {
		return err
	}

	return nil

}

//update for an id updates a location
func UpdateCardsTable(db *sql.DB, name string, id int, maxlevel int, iconurl string) error {

	_, err := db.Exec("UPDATE cards SET name=(?),id=(?),maxlevel=(?) WHERE iconurl=(?)", name, id, maxlevel, iconurl)

	if err != nil {
		return err
	}

	return nil

}

// // Returning slice of Locations Info from DB Table locations ALLInfo about Location
// func GetAllCards(db *sql.DB)([]Cards,error){

// 	rows, _ := db.Query("SELECT name,id,maxlevel,iconurl from cards")

// 	defer rows.Close()

// 	return cardsrows(rows)
// }

// func cardsrows (rows *sql.Rows)([]Cards,error){
// 	var card  []Cards

// 	for rows.Next() {
// 		var c Cards
// 		err:=rows.Scan(&c.Items)
// 		if err!=nil {
// 			return nil,err
// 		}

// 		card=append(card,c)
// 	}

// 	return card,nil
// }

func GetAllCards(db *sql.DB) ([]CardsInfo, error) {

	var cards []CardsInfo
	var card CardsInfo

	rows, err := db.Query("SELECT name,id,maxlevel,iconurl FROM cards")

	if err != nil {
		return cards, err
	}

	for rows.Next() {
		err := rows.Scan(&card.Name, &card.ID, &card.MaxLevel, &card.IconUrls.Medium)

		if err != nil {
			return cards, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

// Returning slice of card structure with cardname u get cardname,id,maxlevel,iconurl
func GetCardLike(db *sql.DB, name string) ([]Cards, error) {

	var card []Cards
	rows, err := db.Query("SELECT name,id,maxlevel,iconurl FROM cards Where name Like (?)", "%"+name+"%")

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var c Cards
		err = rows.Scan(&c.Items)

		if err != nil {
			return nil, err
		}
		card = append(card, c)
	}

	if len(card) == 0 {
		return nil, errors.Database(sql.ErrNoRows)
	}

	return card, nil
}
