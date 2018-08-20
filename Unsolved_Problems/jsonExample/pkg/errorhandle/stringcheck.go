package errorhandle

import (
	"unicode"
	"strings"
	"regexp"
	"errors"
)
type IsString struct {
	arf int
	msg string
}

var Err = &IsString{10,"It shouldn't contain ints"}

func (e *IsString) Error() string {
	return e.msg
}


	func CheckString(Msg *string)  error{
		*Msg = strings.ToLower(*Msg)
		//e:=&IsString{"It shouldn't contain ints"}
		if  LettersOnly(*Msg) != true {
			//fmt.Println(Err)
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

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func CheckEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}