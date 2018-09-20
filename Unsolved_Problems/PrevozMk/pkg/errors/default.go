package errors

import "github.com/pkg/errors"

//Gets the type of error and the error message,prints it and transforms it into a single error
func Default(typeOfError string,err error)error{
	newErr := errors.Errorf("ERROR %s: %s",typeOfError,err)
	return newErr
}