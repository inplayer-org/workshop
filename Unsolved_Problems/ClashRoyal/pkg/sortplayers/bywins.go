package sortplayers

import (
	"fmt"
	"net/http"
	"encoding/json"
	url2 "net/url"
	"sort"
)

type PlayerStats struct {
	Tag    string `json:"tag"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
}

type byWins []PlayerStats

func (p byWins) Len() int {
	return len(p)
}
func (p byWins) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p byWins) Less(i, j int) bool {
	if p[i].Wins!=p[j].Wins {
		return p[i].Wins>p[j].Wins
	}
	if p[i].Losses!=p[j].Losses {
		return p[i].Losses<p[j].Losses
	}
		return p[i].Name<p[j].Name
}



//Sort players in order by priority list 1.Wins ,2.Losses ,3.Name
func ByWins(playerTags []string){

	//Mock Object
	//var playerTag []string
	//playerTag = append(playerTag,"%258GPVRJJPL" )
	//playerTag = append(playerTag,"%2582LVJ9LC")
	//playerTag = append(playerTag,"%25PYJV8290" )
	//playerTag = append(playerTag,"%25Q2RJ2RP2" )

	var currentPlayer PlayerStats
	var Players []PlayerStats


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
	//Player = append(Player,PlayerStats{Name:"B Player",Wins:4279,Losses:3000,Tag:"TestTag"})
	//Player = append(Player,PlayerStats{Name:"A Player",Wins:4279,Losses:3000,Tag:"TestTag1"})


	//fmt.Println("Pre sort Array = ",Player)
	//fmt.Println("is it sorted ? - ",sort.IsSorted(byWins(Player)),"\n")
	//
	sort.Sort(byWins(Players))
	//

	fmt.Printf("%4s  %-30s %-6s %-6s \n","Rank","Name","Wins","Loses")


	for i,elem := range Players{
		fmt.Printf("%4d. %-30s %-6d %-6d \n",i+1,elem.Name,elem.Wins,elem.Losses,)
	}

	fmt.Println()
	//fmt.Println("is it sorted ? - ",sort.IsSorted(byWins(Player)),"\n")

	}