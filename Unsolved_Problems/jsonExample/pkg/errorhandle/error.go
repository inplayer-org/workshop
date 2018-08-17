package errorhandle

import (
	"unicode"
	"strings"
)
type IsString struct {
	msg string
}

func (error *IsString) Error() string {
	return error.msg
}


	func CheckString(Msg *string)  error{
		*Msg = strings.ToLower(*Msg)
		e:=&IsString{"It shouldn't contain ints"}
		if  LettersOnly(*Msg) != true {
			return e
		}
		return nil
	}


	func LettersOnly(str string) bool{
		for _, l := range str {
		if !unicode.IsLetter(l) {
		return false
	}
	}
		return true

	}

