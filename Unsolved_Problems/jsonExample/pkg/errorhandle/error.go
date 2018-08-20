package errorhandle

import (
	"unicode"
	"strings"
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

