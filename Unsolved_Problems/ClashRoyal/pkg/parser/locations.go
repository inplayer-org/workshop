package parser

import "repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"

func FirstCharFrom(name string)(string,error){
	if len(name)==0{
		return "",errors.NewResponseError("Empty string","Can't get the first character from empry string",400)
	}
	 return name[:1],nil
}
