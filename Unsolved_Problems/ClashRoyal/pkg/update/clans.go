package update

import (
	"database/sql"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

//not used
/*func UpdateClans(db *sql.DB, clan []structures.Clan) error {

	for _, elem := range clan {

		//if not exist
		_, err := db.Exec("INSERT INTO clans(clanTag,clanName) VALUES (?,?)", elem.Tag, elem.Name)

		if err != nil {
			return err
		}

	}
	_, err := db.Exec("UPDATE clans SET clanTag=(?), clanName=(?) WHERE clanTag=(?)")
	if err != nil {
		return err
	}

	return err
}*/

//GetAllClans - Returns slice of all structure Clan present in the database
func GetAllClans(db *sql.DB) ([]structures.Clan, error) {

	var clans []structures.Clan
	var clan structures.Clan

	rows, err := db.Query("SELECT clanTag,clanName FROM clans;")

	if err != nil {
		return clans, err
	}

	for rows.Next() {
		err := rows.Scan(&clan.Tag, &clan.Name)

		if err != nil {
			return clans, err
		}

		clans = append(clans, clan)
	}

	return clans, nil
}


//GetRequestForPlayersFromClan -
func GetRequestForPlayersFromClan(db *sql.DB,clanTag string)error{
	client := _interface.NewClient()
	tags,err:= client.GetTagByClans(clanTag)

	if err!=nil{
		return err
	}

	done := make(chan error)

	countChan := 0

	for _,playerTag := range tags.Player {
		go chanRequest(db,client,playerTag.Tag,done)
		countChan++
	}

	for ;countChan>0;countChan--{
		log.Println("done = ",<-done)
	}

	return nil
}

func chanRequest(db *sql.DB,clientInterface _interface.ClientInterface,playerTag string,done chan <- error){
	players,err:=clientInterface.GetRequestForPlayer(parser.ToUrlTag(playerTag))

	if err!=nil{
		done<-err
	}else{
		var i int
		if players.LocationID!=nil {
			i=players.LocationID.(int)
		}else{
			i=0
		}

		err:=queries.UpdatePlayer(db,players,i)
		done<-err
	}
}

//not used
/*func GetClans(clanTag string)(structures.Clan,error){

	var clans structures.Clan

	client := &http.Client{}

	req,err :=http.NewRequest("GET","https://api.clashroyale.com/v1/clans/"+ clanTag +"",nil)

	if err!=nil{
		return clans,err
	}

	req.Header.Add("Content-Type","application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjBkMTUxODQ4LWM0ZTgtNGU1Zi05NzRiLWQzNjQ1ZjAxMzk2MiIsImlhdCI6MTUzNDg1NDQ2MCwic3ViIjoiZGV2ZWxvcGVyL2U1ODJhZWJlLWNlNGUtNGVhMC1hZTgwLTk5MTdhMmNkMGZhYyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.8-GoA48DGZScCOi6EU4AAuJUcXbY2kqqHwsEXg22w4hDHJegjuSaS6jjDSoZcZFSS9x6Fbkd825eSagpAjbX4Q")

	resp,err:=client.Do(req)

	if err!=nil{
		return clans,err
	}

	json.NewDecoder(resp.Body).Decode(&clans)

	return clans,nil

}*/
/*
func GetTagByClans(clanTag string) []string {

	var tag structures.PlayerTags
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.clashroyale.com/v1/clans/"+clanTag+"/members", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjM3OTUwMGRiLTU3YWYtNDcwNy04NWE2LTFhMjMwMjRjMDEyOCIsImlhdCI6MTUzNDkyOTM3Nywic3ViIjoiZGV2ZWxvcGVyLzdjMDEzZWE0LTE1YTMtZDJhNS04MmVlLTJiMzkxYmM0YWM0MSIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCIsIjYyLjE2Mi4xNjguMTk0IiwiNjIuMTYyLjE2OC4xOTQiXSwidHlwZSI6ImNsaWVudCJ9XX0.ha9ITX8-_1sHRi6y2pFCWUJiyv2dvlX8BWqG5x1l9mLE1FbFfNIe-ZcgMZlglcjhE4uaHSSFAaaC-FMIyXYywg")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&tag)
	return parser.ToUrlTags(tag.GetTags())
}

*/