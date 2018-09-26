package boards

type Board struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	ShortUrl string `json:"shortUrl"`
}
