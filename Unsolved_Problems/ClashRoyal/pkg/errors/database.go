package errors

import (
	"database/sql"

)

//Checks and prints errors correlated to the sql database package and returns it
func Database(err error)error{
	if err==sql.ErrNoRows{
		newErr := Default("ErrNoRows",err)
		return newErr
	}else if err==sql.ErrConnDone{
		newErr := Default("ErrConnDone",err)
		return newErr
	}else if err!=nil{
		newErr := Default("Undefined",err)
		return newErr
	}
	return nil
}