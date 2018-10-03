package interfaces

import (
	"database/sql"
)

type DataStructure interface {
	NewDataStructure()DataStructure
	Insert(*sql.DB)error
	Update(*sql.DB)error
}
