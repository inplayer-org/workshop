package parser

import "errors"

//ToUrlTag - Removes the first character (should be #) and concatenates %25 in from for making a request from the Clash Royal api by tag for a single string entry
func ToRequestTag(tag string) string{

		return "%25" + tag[1:]

}

func ToHashTag(tag string)string {

	return "#"+tag

}

func ToRawTag(tag string)string {

	return tag[1:]

}

func FromAnyToHashTag(tag string)(string,error){
	if len(tag) == 0{
		return "",errors.New("tried to convert empty string")
	}
	for {
		if len(tag) ==0{
			return "",errors.New("string is only consisted of #, %25,% and &")
		}
		switch string(tag[0]) {
		case "#":
			tag=tag[1:]
			break
		case "%":
			if tag[:3]=="%25"{
				tag = tag[3:]
			}else {
				tag=tag[1:]
			}
			break
		case "&":
			tag=tag[1:]
			break
		default:
			return "#"+tag,nil
		}
	}
}