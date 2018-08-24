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
)

func Players(DB *sql.DB,playerTags []string,locationID int){

	var currentPlayer structures.PlayerStats


	baseUrl := "https://api.clashroyale.com/v1/players/"

	client := http.Client{}

	reqHeader :=	http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImM4NmY2ZjFjLWZlMGEtNDFiZi05NWRlLWY0MWYyYTY0ODQyNCIsImlhdCI6MTUzNDg1MTkxNSwic3ViIjoiZGV2ZWxvcGVyLzNiYmFmOGRhLTg0YmMtOWQyMi1iM2QwLTRlNDA3NmRhMWEzOCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ._3ZnRmHHPLamGQO6QtNdXXVe7V6hAjrpA0z3sTBl7wJ937U8KEUozers1ZiLQUImDGTz_m-XLAFBlE-12DqVAw")

	for _,nextTag := range playerTags{

		req,err := url.Parse(baseUrl+nextTag)

		if err!=nil{
			fmt.Println(err)
		}

		resp,err := client.Do(&http.Request{Method:"GET",Header:reqHeader,URL:req})

		if err!=nil{
			fmt.Println(err)
		}

		json.NewDecoder(resp.Body).Decode(&currentPlayer)
		currentPlayer.Tag = "#"+currentPlayer.Tag[1:]
		queries.UpdatePlayer(DB,currentPlayer,locationID)


	}
	log.Println("Finished Updating for location",locationID )
}
