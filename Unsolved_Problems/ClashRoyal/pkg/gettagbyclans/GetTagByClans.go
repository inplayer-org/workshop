package gettagbyclans

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Tag struct {
	Tag []struct {
		tag string `json:"tag"`
	}
}

func GetTagByClans() {
	var tag Tag
	client := &http.Client{}

	req,err :=http.NewRequest("GET",`https://api.clashroyale.com/v1/clans/\#9VPC0JCP/members`,nil)

	if err!=nil{
		panic(err)
	}

	req.Header.Add("Content-Type","application/json")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImNhOGU4NzI3LTgxMjMtNGE3NC04MGEwLTU0MDNmOWY5NDllMSIsImlhdCI6MTUzNDg0NzQ2NCwic3ViIjoiZGV2ZWxvcGVyLzdjMDEzZWE0LTE1YTMtZDJhNS04MmVlLTJiMzkxYmM0YWM0MSIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyI2Mi4xNjIuMTY4LjE5NCJdLCJ0eXBlIjoiY2xpZW50In1dfQ.CDPfyLPRgd2YRtYGSdq3dNvVCgXZgUWVpFMRZ_guXuGQp3q2IqJbDBkHS_kC2QiA6W6rFieetKSQ8XBKtAY0rQ")

	resp,err:=client.Do(req)

	if err!=nil{
		panic(err)
	}
	json.NewDecoder(resp.Body).Decode(&tag)

	fmt.Print(tag)
}



func main() {
	GetTagByClans()
}