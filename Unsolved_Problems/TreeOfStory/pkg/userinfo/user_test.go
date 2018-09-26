package userinfo_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"repo.inplayer.com/workshop/Unsolved_Problems/TreeOfStory/pkg/userinfo"
)

func newTestUserInfo() userinfo.Userinfo {
	i := userinfo.Userinfo{
		UserName: "jana12",
		FullName: "Jana",
	}
	return i
}
func CreateUser(t *testing.T) {
	log.Println("tukas")
	db, _ := sql.Open("mysql", "root:1111/TestTreeOfStory")
	i := newTestUserInfo()
	// Insert an User
	createdItem, err := userinfo.InsertIntoUserInfoTableTest(db, i)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the item's id was set
	if createdItem.ID == 0 {
		t.Fatal("id not set")
	}

}

// func TestUpdateItemByID(t *testing.T) {
// 	s := newTestStore()
// 	defer s.Close()

// 	i := newTestItem()

// 	createdItem, err := s.CreateItem(i)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Update the items title
// 	newTitle := randomString(16)
// 	itemToUpdate := createdItem
// 	itemToUpdate.Title = newTitle
// 	itemToUpdate.Completed = true
// 	err = s.UpdateItemByID(fmt.Sprintf("%v", createdItem.ID), itemToUpdate)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	updatedItem, err := s.FindItemByID(fmt.Sprintf("%v", createdItem.ID))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Verify the title
// 	if updatedItem.Title != itemToUpdate.Title {
// 		t.Fatalf("title not udpdated; expected %v, got %v\n", itemToUpdate.Title, updatedItem.Title)
// 	}

// 	// Verify completed status
// 	if updatedItem.Completed != itemToUpdate.Completed {
// 		t.Fatalf("completed status not updated; expected %v, got %v\n", itemToUpdate.Completed, updatedItem.Completed)
// 	}
// }
