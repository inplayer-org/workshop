package cards

import (
	"time"
)

type Card struct {
	ID string `json:"id"`
	CheckItems int `json:"checkItems"`
	CheckItemsChecked int `json:"checkItemsChecked"`
	Description bool `json:"description"`
	DateLastActivity time.Time `json:"dateLastActivity"`
	Descrip string `json:"descrip"`
	ShortLink string `json:"shortLink"`//POSSIBLE CHANGE
	ShortURL string `json:"shortUrl"`//POSSIBLE CHANGE
}
