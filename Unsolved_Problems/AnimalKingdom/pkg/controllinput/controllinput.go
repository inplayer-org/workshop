package controllinput

import (
	"strconv"
	"strings"
	"unicode"
)




//Changes the value of the string argument and return bool
func CheckString(str *string) bool{
	*str = strings.ToLower(*str)
	return StringLengthBetween(*str,1,30) &&  LettersOnly(*str)
}

func LettersOnly(str string) bool{
	for _, l := range str {
		if !unicode.IsLetter(l) {
			return false
		}
	}
	return true
}

func IntOnly(str string) bool{
	if _, err := strconv.Atoi(str); err != nil{
		return false
	}
	return true
}

func StringLengthBetween(str string, first int, second int) bool{
	return len(str)>first && len(str)<second
}

