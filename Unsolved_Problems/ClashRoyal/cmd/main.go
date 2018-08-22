package main

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
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

	tags:=parser.TagsToTagsUrlTags(playerTags.GetTags())

	fmt.Println(tags)

	}