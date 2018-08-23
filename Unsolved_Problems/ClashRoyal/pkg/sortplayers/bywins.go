package sortplayers

import (
	"fmt"
	"net/http"
	"encoding/json"
	url2 "net/url"
	"sort"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)



//Sort players in order by priority list 1.Wins ,2.Losses ,3.Name
func byWins(playerTags []string){

	//Mock Object
	//var playerTag []string
	//playerTag = append(playerTag,"%258GPVRJJPL" )
	//playerTag = append(playerTag,"%2582LVJ9LC")
	//playerTag = append(playerTag,"%25PYJV8290" )
	//playerTag = append(playerTag,"%25Q2RJ2RP2" )

	var currentPlayer structures.PlayerStats
	var Players []structures.PlayerStats


	baseUrl := "https://api.clashroyale.com/v1/players/"

	client := http.Client{}

	reqHeader :=	http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImM4NmY2ZjFjLWZlMGEtNDFiZi05NWRlLWY0MWYyYTY0ODQyNCIsImlhdCI6MTUzNDg1MTkxNSwic3ViIjoiZGV2ZWxvcGVyLzNiYmFmOGRhLTg0YmMtOWQyMi1iM2QwLTRlNDA3NmRhMWEzOCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ._3ZnRmHHPLamGQO6QtNdXXVe7V6hAjrpA0z3sTBl7wJ937U8KEUozers1ZiLQUImDGTz_m-XLAFBlE-12DqVAw")

	for _,nextTag := range playerTags{

		req,err := url2.Parse(baseUrl+nextTag)

		if err!=nil{
			fmt.Println(err)
		}

		resp,err := client.Do(&http.Request{Method:"GET",Header:reqHeader,URL:req})

		if err!=nil{
			fmt.Println(err)
		}

		json.NewDecoder(resp.Body).Decode(&currentPlayer)

		Players = append(Players,currentPlayer)

	}

	//Mock Test Entries
	//Player = append(Player,structures.PlayerStats{Name:"B Player",Wins:4279,Losses:3000,Tag:"TestTag"})
	//Player = append(Player,structures.PlayerStats{Name:"A Player",Wins:4279,Losses:3000,Tag:"TestTag1"})


	//fmt.Println("Pre sort Array = ",Player)
	//fmt.Println("is it sorted ? - ",sort.IsSorted(ByWins(Player)),"\n")
	//
	sort.Sort(ByWins(Players))
	//

	fmt.Printf("%4s  %-30s %-6s %-6s \n","Rank","Name","Wins","Loses")


	for i,elem := range Players{
		fmt.Printf("%4d. %-30s %-6d %-6d \n",i+1,elem.Name,elem.Wins,elem.Losses,)
	}

	fmt.Println()
	//fmt.Println("is it sorted ? - ",sort.IsSorted(ByWins(Player)),"\n")

	}