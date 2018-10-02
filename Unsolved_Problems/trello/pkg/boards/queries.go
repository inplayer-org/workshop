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




// Returning string(board name from DB table boards)
func GetBoardName(db *sql.DB,boardID string)(string,error){


	var boardName string

	err := db.QueryRow("SELECT nameBoard FROM Boards WHERE ID=?",boardID).Scan(&boardName)

	if err!=nil{
		return boardName,err
	}

	return boardName,nil

}

// Returns slice of all boards present in the database
func GetAllBoards(db *sql.DB) ([]Board, error) {

	var boards []Board
	var board Board

	rows, err := db.Query("SELECT ID,nameBoard,descrip,shortUrl FROM Boards;")

	if err != nil {
		return boards, err
	}

	for rows.Next() {
		err := rows.Scan(&board.ID, &board.Name,&board.Desc,&board.ShortUrl)

		if err != nil {
			return boards, err
		}

		boards = append(boards, board)
	}

	return boards, nil
}


func GetBoardbyID(db *sql.DB,boardID string)(Board,error){



	var board Board

	err := db.QueryRow("SELECT nameBoard,descrip,shortUrl FROM Boards WHERE ID=?",boardID).Scan(&board.Name,&board.Desc,&board.ShortUrl)

	if err!=nil{
		return board,err
	}

	return board,nil

}
