package main

import (
	locations2 "repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"fmt"
)

func main (){

	locations,err:=locations2.GetLocations()

	if err!=nil {
		panic(err)
	}

	locationsMap:=locations2.LocationMap(locations)

	fmt.Println(locationsMap)

	}