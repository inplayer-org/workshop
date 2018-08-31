package update

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"net/http"
	"net/url"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"log"
	"time"
)

func Players(DB *sql.DB,playerTags []string,locationID int,done chan <- interface{}){

	var currentPlayer structures.PlayerStats


	baseUrl := "https://api.clashroyale.com/v1/players/"

	client := http.Client{}

	reqHeader :=	http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImM4NmY2ZjFjLWZlMGEtNDFiZi05NWRlLWY0MWYyYTY0ODQyNCIsImlhdCI6MTUzNDg1MTkxNSwic3ViIjoiZGV2ZWxvcGVyLzNiYmFmOGRhLTg0YmMtOWQyMi1iM2QwLTRlNDA3NmRhMWEzOCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ._3ZnRmHHPLamGQO6QtNdXXVe7V6hAjrpA0z3sTBl7wJ937U8KEUozers1ZiLQUImDGTz_m-XLAFBlE-12DqVAw")

	for _,nextTag := range playerTags{

		req,err := url.Parse(baseUrl+nextTag)
		errorLimit := 0
		if err!=nil{
			fmt.Println(err)
		}
		for {
			resp, err := client.Do(&http.Request{Method: "GET", Header: reqHeader, URL: req})

			if err != nil {
				fmt.Println(err)
			}
			if resp.StatusCode>=200 && resp.StatusCode<=300{

				json.NewDecoder(resp.Body).Decode(&currentPlayer)
				currentPlayer.Tag = "#"+currentPlayer.Tag[1:]
				queries.UpdatePlayer(DB,currentPlayer,locationID)

				break
			}
			//log.Println("REQUEST PROBLEM !! -> ",resp.Status,",  Retrying ...")
			if resp.StatusCode!=http.StatusTooManyRequests{
				errorLimit++
			}
			if errorLimit>=9{
				log.Println(nextTag,currentPlayer," Failed to get a response from the API ",resp.Status)
				break
			}
			time.Sleep(time.Second*1)
			}



	}
		if locationID!=0 {
			done <- locationID
		}else{
			done <- currentPlayer.Clan.Name
		}
}


func GetRequestForPlayer(db *sql.DB,tag string)int{

	var currentPlayer structures.PlayerStats


	baseUrl := "https://api.clashroyale.com/v1/players/"

	client := http.Client{}

	reqHeader :=	http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImM4NmY2ZjFjLWZlMGEtNDFiZi05NWRlLWY0MWYyYTY0ODQyNCIsImlhdCI6MTUzNDg1MTkxNSwic3ViIjoiZGV2ZWxvcGVyLzNiYmFmOGRhLTg0YmMtOWQyMi1iM2QwLTRlNDA3NmRhMWEzOCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ._3ZnRmHHPLamGQO6QtNdXXVe7V6hAjrpA0z3sTBl7wJ937U8KEUozers1ZiLQUImDGTz_m-XLAFBlE-12DqVAw")

	req,err := url.Parse(baseUrl+tag)

	if err!=nil{
		fmt.Println(err)
	}

	for {
		resp, err := client.Do(&http.Request{Method: "GET", Header: reqHeader, URL: req})

		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode>=200 && resp.StatusCode<=300{

			json.NewDecoder(resp.Body).Decode(&currentPlayer)
			currentPlayer.Tag = "#"+currentPlayer.Tag[1:]
			queries.UpdatePlayer(db,currentPlayer,0)

			break
		}
		//log.Println("REQUEST PROBLEM !! -> ",resp.Status,",  Retrying ...")
		if resp.StatusCode!=http.StatusNotFound{
			return 404
		}
		time.Sleep(time.Second*1)
	}


return 0
}
