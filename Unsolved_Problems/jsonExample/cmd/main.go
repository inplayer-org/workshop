package main

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/application"
)

func main() {
	a := application.App{}
	a.Initialize("root", "darko123", "inplayerdb")

	a.Run(":3030")
}