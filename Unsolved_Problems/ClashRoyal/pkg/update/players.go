package update

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

func NewError(errorType string,locationID int,playerTag string,tags structures.PlayerStats,msg interface{})error{

	return errors.Errorf("ERROR WITH %s for location id %d: playerTag - %s, structure - %s, %s",
		errorType,locationID,playerTag,tags,msg)

}


func Players(DB *sql.DB,playerTags []string,locationID int)[]error{

	var currentPlayer structures.PlayerStats

	var allErrors []error

	baseUrl := "https://api.clashroyale.com/v1/players/"

	client := http.Client{}

	reqHeader :=	http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImM4NmY2ZjFjLWZlMGEtNDFiZi05NWRlLWY0MWYyYTY0ODQyNCIsImlhdCI6MTUzNDg1MTkxNSwic3ViIjoiZGV2ZWxvcGVyLzNiYmFmOGRhLTg0YmMtOWQyMi1iM2QwLTRlNDA3NmRhMWEzOCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ._3ZnRmHHPLamGQO6QtNdXXVe7V6hAjrpA0z3sTBl7wJ937U8KEUozers1ZiLQUImDGTz_m-XLAFBlE-12DqVAw")

	for _,nextTag := range playerTags{



		req,err := url.Parse(baseUrl+nextTag)
		errorLimit := 0
		if err!=nil{
			err = NewError("URL PARSING",locationID, nextTag,currentPlayer,err)
			allErrors = append(allErrors,err)
			continue
		}
		for {
			resp, err := client.Do(&http.Request{Method: "GET", Header: reqHeader, URL: req})

			if err != nil {
				err = NewError("REQUEST",locationID, nextTag,currentPlayer,err)
				allErrors = append(allErrors,err)
				break
			}

			if resp.StatusCode>=200 && resp.StatusCode<=300{

				json.NewDecoder(resp.Body).Decode(&currentPlayer)
				currentPlayer.Tag = "#"+currentPlayer.Tag[1:]
				err := queries.UpdatePlayer(DB,currentPlayer,locationID)

				if err!=nil{
					err = NewError("DATABASE",locationID, nextTag,currentPlayer,err)
					allErrors=append(allErrors,err)
				}
				break



			}
			//log.Println("REQUEST PROBLEM !! -> ",resp.Status,",  Retrying ...")
			if resp.StatusCode!=http.StatusTooManyRequests{
				errorLimit++
			}

			if errorLimit>=9{
				err := NewError("RESPONSE",locationID, nextTag,currentPlayer,"Response Status ->" + resp.Status)
				allErrors = append(allErrors,err)
			}

			}

	}

	return allErrors
}

func GetRequestForPlayer(db *sql.DB,tag string)int{

	var currentPlayer structures.PlayerStats


	baseUrl := "https://api.clashroyale.com/v1/players/"

	client := http.Client{}

	reqHeader :=	http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImM4NmY2ZjFjLWZlMGEtNDFiZi05NWRlLWY0MWYyYTY0ODQyNCIsImlhdCI6MTUzNDg1MTkxNSwic3ViIjoiZGV2ZWxvcGVyLzNiYmFmOGRhLTg0YmMtOWQyMi1iM2QwLTRlNDA3NmRhMWEzOCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ._3ZnRmHHPLamGQO6QtNdXXVe7V6hAjrpA0z3sTBl7wJ937U8KEUozers1ZiLQUImDGTz_m-XLAFBlE-12DqVAw")

	req,err := url.Parse(baseUrl+parser.ToUrlTag(tag))

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
	}

return 0
}

//GetPlayerTagsPerLocation for id returns top 200 playerTags
func GetPlayerTagsPerLocation(ID int)(structures.PlayerTags,error){

	var playerTags structures.PlayerTags

	client := &http.Client{}

	urlStr:="https://api.clashroyale.com/v1/locations/" + strconv.Itoa(ID) + "/rankings/players"

	link,err:=url.Parse(urlStr)

	if err != nil {
		return playerTags,err
	}

	reqHeader:=http.Header{}
	reqHeader.Set("Content-Type","application/json")
	reqHeader.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")

	resp,err:=client.Do(&http.Request{Header:reqHeader,URL:link,Method:"GET"})

	if err!=nil {
		return playerTags,err
	}

	json.NewDecoder(resp.Body).Decode(&playerTags)

	return playerTags,nil
}