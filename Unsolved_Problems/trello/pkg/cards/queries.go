package cards

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
	"time"
)

//sholud be moved in vaidation PKG
type rawDate []byte

func (t rawDate) Time() (time.Time, error) {
	//fmt.Println(t)
	return time.Parse("2006-01-02 15:04:05", string(t))
}


func (c *Card)Insert(DB *sql.DB)error{

	_,err:=DB.Exec("INSERT into Trello.Cards (ID,checkItems,checkItemsChecked,description,dateLastActivity,descrip,shortLink,shortUrl,IDboard,IDlist) values (?,?,?,?,?,?,?,?,?,?);",c.ID,c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortLink,c.ShortURL,c.IDBoard,c.IDList)

	return err

	}

func  (c *Card)Update(DB *sql.DB)error{

	exists,err := validators.ExistsElementInColumn(DB,"Cards",c.ID,"ID")

	if err!= nil{
		return err
	}

	if exists{
		return c.updateByID(DB)
	}

	exists,err = validators.ExistsElementInColumn(DB,"Cards",c.ShortLink,"shortLink")

	if err!= nil{
		return err
	}

	if exists{
		return c.updateByShortLink(DB)
	}
	exists,err = validators.ExistsElementInColumn(DB,"Cards",c.ShortURL,"shortURL")

	if err!= nil{
		return err
	}

	if exists{
		return c.updateByShortURL(DB)
	}


	return c.Insert(DB)

}

func (c *Card) updateByID(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Trello.Cards SET checkItems=?,checkItemsChecked=?,description=?,dateLastActivity=?,descrip=?,shortLink=?,shortUrl=? WHERE ID=?",c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortLink,c.ShortURL,c.ID)

	return err

}

func (c *Card) updateByShortLink(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Trello.Cards SET ID=?,checkItems=?,checkItemsChecked=?,description=?,dateLastActivity=?,descrip=?,shortUrl=? WHERE ShortLink=?",c.ID,c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortURL,c.ShortLink)

	return err

}

func (c *Card) updateByShortURL(DB *sql.DB) error {

	_,err := DB.Exec("UPDATE Trello.Cards SET ID=?,checkItems=?,checkItemsChecked=?,description=?,dateLastActivity=?,descrip=?,shortLink=? WHERE shortUrl=?",c.ID,c.Badges.CheckItems,c.Badges.CheckItemsChecked,c.Badges.Description,c.DateLastActivity,c.Descrip,c.ShortLink,c.ShortURL)

	return err

}


// Returning slice of label structure with boardid u get labelname and labelid
func GetCardsFromBoard(db *sql.DB,boardID string)([]Card,error){

	var cards []Card

	rows,err:=db.Query("SELECT ID,checkItems,checkItemsChecked,description,dateLastActivity,descrip,shortLink,shortUrl,IDlist FROM Cards Where IDboard Like (?)","%"+boardID+"%")

	if err !=nil {
		return nil,err
	}

	for rows.Next(){

		var dt rawDate

		var c Card


		err = rows.Scan(&c.ID,&c.Badges.CheckItems,&c.Badges.CheckItemsChecked,&c.Badges.Description,&dt,&c.Descrip,&c.ShortLink,&c.ShortURL,&c.IDList)

		if err !=nil {
			//fmt.Println(err.Error())
			return nil,err
		}

		//fmt.Println(dt.Time())

		c.DateLastActivity,err=dt.Time()

		if err !=nil {
			return nil,err
		}

		cards = append(cards,c)
	}


	return cards,nil
}

func GetCardsFromList(DB *sql.DB,listID string)([]Card,error){

	var cards []Card
	rows,err:=DB.Query("SELECT ID,checkItems,checkItemsChecked,description,dateLastActivity,descrip,shortLink,shortUrl,IDboard,IDlist FROM Cards Where IDlist=?;",listID)

	if err!=nil{
		return nil,err
	}

	for rows.Next(){

		var dt rawDate

		var c Card

		err = rows.Scan(&c.ID,&c.Badges.CheckItems,&c.Badges.CheckItemsChecked,&c.Badges.Description,&dt,&c.Descrip,&c.ShortLink,&c.ShortURL,&c.IDBoard,&c.IDList)

		if err !=nil {
			//fmt.Println(err.Error())
			return nil,err
		}

		//fmt.Println(dt.Time())

		c.DateLastActivity,err=dt.Time()

		if err !=nil {
			return nil,err
		}

		cards = append(cards,c)
	}


	return cards,nil
}