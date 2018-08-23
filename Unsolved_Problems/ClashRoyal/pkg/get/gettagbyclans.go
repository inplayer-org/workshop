package get

import (
	"net/http"
	"encoding/json"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)






func GetTagByClans(clanTag string) []string{
	var tag structures.PlayerTags
	client := &http.Client{}

	req,err :=http.NewRequest("GET","https://api.clashroyale.com/v1/clans/"+ clanTag +"/members",nil)

	if err!=nil{
		panic(err)
	}

	req.Header.Add("Content-Type","application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjM3OTUwMGRiLTU3YWYtNDcwNy04NWE2LTFhMjMwMjRjMDEyOCIsImlhdCI6MTUzNDkyOTM3Nywic3ViIjoiZGV2ZWxvcGVyLzdjMDEzZWE0LTE1YTMtZDJhNS04MmVlLTJiMzkxYmM0YWM0MSIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCIsIjYyLjE2Mi4xNjguMTk0IiwiNjIuMTYyLjE2OC4xOTQiXSwidHlwZSI6ImNsaWVudCJ9XX0.ha9ITX8-_1sHRi6y2pFCWUJiyv2dvlX8BWqG5x1l9mLE1FbFfNIe-ZcgMZlglcjhE4uaHSSFAaaC-FMIyXYywg")
	resp,err:=client.Do(req)

	if err!=nil{
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&tag)
	return parser.ToUrlTags(tag.GetTags())
}

