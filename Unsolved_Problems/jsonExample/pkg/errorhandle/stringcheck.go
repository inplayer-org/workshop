package errorhandle

import (
	"strings"
	"unicode"
	"regexp"
	"errors"
)

type BadString struct {
	msg string
}

var Err = &BadString{"It shouldn't contain ints"}

func (e *BadString) Error() string {
	return e.msg
}

func (e *BadString)setMsg(msg string){
	e.msg=msg+e.Error()
}

func CheckString(Msg *string)  error{
		*Msg = strings.ToLower(*Msg)
		if  LettersOnly(*Msg) != true {
			Err.setMsg(*Msg)
			return Err
		}
		return nil
	}

func LettersOnly(str string) bool{
		for _, l := range str {
		if unicode.IsNumber(l) {
		return false
	}
	}
		return true


	}

var (
	ErrBadFormat        = errors.New("invalid format of Email Address")

	emailRegexp = regexp.MustCompile( "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$" )
)

func CheckEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

var ErrBadSalary = errors.New("invalid input for salary")

func CheckSalary(salary string) error{

	if !strings.HasSuffix(salary, "den.") {
		return ErrBadSalary
	}

	s := []rune(salary)
	for i:=0; i< len(s)-4;  i++{
		if !unicode.IsNumber(s[i]) {
			return ErrBadSalary
		}
	}
	return nil

	}