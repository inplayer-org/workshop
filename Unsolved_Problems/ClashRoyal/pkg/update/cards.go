package update

import (
	"database/sql"
	"strconv"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
)

//var wg sync.WaitGroup

func Cards(db *sql.DB, c cards.Cards) error {
	done := make(chan error, 300)

	for _, elem := range c.Items {
		var cardInfo cards.CardsInfo
		cardInfo = elem
		wg.Add(1)
		go CurrentCardUpdate(db, cardInfo, done)
	}
	wg.Wait()

	close(done)

	for err := range done {

		if err != nil {
			return err
		}
	}

	return nil
}

func CurrentCardUpdate(db *sql.DB, elem cards.CardsInfo, done chan<- error) {
	var err error
	defer wg.Done()

	for {
		if !parser.Exists(db, "cards", "id", strconv.Itoa(elem.ID)) {
			err = cards.InsertIntoCardsTable(db, elem.Name, elem.ID, elem.MaxLevel, elem.IconUrls.Medium)
		} else {
			err = cards.UpdateCardsTable(db, elem.Name, elem.ID, elem.MaxLevel, elem.IconUrls.Medium)
		}

		if err == nil {
			break
		}

	}
	done <- err

}
func CardsUpdate(db *sql.DB) (cards.Cards, error) {
	client := _interface.NewClient()
	cardss, err := client.GetCards()

	if err != nil {
		return cardss, err
	}

	err = Cards(db, cardss)

	if err != nil {
		return cardss, err
	}

	return cardss, nil

}
