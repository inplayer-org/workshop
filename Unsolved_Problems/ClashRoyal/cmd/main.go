package main

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/gettagbyclans"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/sortplayers"
)

func main (){

	loc,err:=locations.GetLocations()

	if err!=nil {
		panic(err)
	}

	locationsMap:=locations.LocationMap(loc)

	mkdID:=locationsMap["Albania"]

	fmt.Println(mkdID)

	playerTags,err:=locations.GetPlayerTagsPerLocation(mkdID)

	if err!=nil{
		panic(err)
	}

	tagsFromLoc:=parser.ToUrlTags(playerTags.GetTags())

	//fmt.Println(tags)

	tagsFromClan := gettagbyclans.GetTagByClans(parser.ToUrlTag("#2LQGYRV"))

	sortplayers.ByWins(tagsFromClan)

	sortplayers.ByWins(tagsFromLoc)

	}