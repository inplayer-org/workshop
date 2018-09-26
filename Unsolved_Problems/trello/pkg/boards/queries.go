package boards

import (
	"database/sql"

	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
)

func (board *Board) Insert(DB *sql.DB) error {

	_, err := DB.Exec("INSERT INTO Trello.Boards (ID,nameBoard,descrip,shortUrl) VALUES (?,?,?,?);", board.ID, board.Name, board.Desc, board.ShortUrl)

	return err

}

func (board *Board) Update(DB *sql.DB) error {

	exists, err := validators.ExistsElementInColumn(DB, "Boards", board.ID, "ID")

	if err != nil {
		return err
	}

	if exists {
		return board.updateByID(DB)
	}

	return board.Insert(DB)

}

func (board *Board) updateByID(DB *sql.DB) error {

	_, err := DB.Exec("UPDATE Trello.Boards SET nameBoard=?,descrip=?,shortUrl=? WHERE ID=?", board.Name, board.Desc, board.ShortUrl, board.ID)

	return err
}
