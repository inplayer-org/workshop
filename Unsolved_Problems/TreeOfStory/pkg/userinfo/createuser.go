package userinfo

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/pkg/files"
)

// if query  insert into USerinfotable calling CreateDirectory to create directory with id of UserInfo
func CreateUser(db *sql.DB,elem Userinfo) error {

	var err error
	var userID int

		if !queries.Exists(db, "UserInfo", "id", strconv.Itoa(elem.ID)) {
			userID,err = InsertIntoUserInfoTable(db, elem.ID, elem.UserName, elem.FullName)
		} else {
			userID,err = UpdateUserInfoTable(db, elem.ID, elem.UserName, elem.FullName)
		}

		if err != nil {
			path := "/home/darko/go/src/repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/user/" + strconv.Itoa(userID)
			files.CreateDirectory(path)
		}


		return nil
	}


