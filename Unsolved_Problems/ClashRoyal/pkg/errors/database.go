package errors

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

//Checks and prints errors correlated to the sql database package and returns it
func Database(err error)error{
	if err==sql.ErrNoRows{
		newErr := errors.Errorf("ERROR ErrNoRows: %s",err)
		log.Println(newErr)
		return err
	}else if err==sql.ErrConnDone{
		newErr := errors.Errorf("ERROR ErrConnDone: %s",err)
		log.Println(newErr)
		return err
	}else if err!=nil{
		newErr := errors.Errorf("ERROR Undefined: %s",err)
		log.Println(newErr)
		return err
	}
	return nil
}