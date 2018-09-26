package userinfo

import (
	"database/sql"
	"errors"
)

//insertInto inserts userinfo into database
func InsertIntoUserInfoTableTest(db *sql.DB, u Userinfo) (Userinfo, error) {

	result, err := db.Exec("INSERT INTO UserInfo(UserName,FullName) VALUES ((?),(?));", u.UserName, u.FullName)

	if err != nil {
		return u, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return u, err
	}

	// Set the todo Item's id before returning
	u.ID = int(id)

	return u, nil

}

//update for an id updates a userinfo
func UpdateUserInfoTableTest(db *sql.DB, u Userinfo) error {

	result, err := db.Exec("UPDATE UserInfo SET UserName=(?),FullName=(?) WHERE ID=(?)", u.UserName, u.FullName, u.ID)

	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("no rows updated")
	}
	return nil
}
