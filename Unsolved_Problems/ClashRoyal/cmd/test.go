package main

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/HandlersFunc"
)

func main() {
	a := HandlersFunc.App{}
	a.Initialize("root", "darko123", "Clash_Royale")

	a.Run(":3030")
}
