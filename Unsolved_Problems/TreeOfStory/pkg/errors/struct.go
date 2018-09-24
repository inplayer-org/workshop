package errors

import "fmt"

type MyError struct{

	Message string
}

func NewMyError(m string)error{
	return &MyError{m}
}

func(e *MyError)Error()string{
	return fmt.Sprintf("%s", e.Message)
}