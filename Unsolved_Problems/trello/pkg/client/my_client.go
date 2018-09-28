package client

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/interfaces"
	"log"
	"net/url"
)


//MyClient structure have client that Do rquests
type MyClient struct {
	client *http.Client
}

//NewClient constructs MyClient
func NewClient() interfaces.ClientInterface {
	return &MyClient{&http.Client{}}
}

//SetHeaders sets the headers to make the request
func SetHeaders(req *http.Request,values url.Values){//,m map[string]string

	q := req.URL.Query()
	q.Add("key","9ecdc5f04a4ccb643b83d4fd2b920416")
	q.Add("token","125a712b04063b34d2c22392704bb38a5fc88bb48f665c0f6bdf2d516f473c9d")

	req.Header.Add("Content-Type","application/json")

	if values!=nil {
		req.URL.RawQuery = q.Encode() + "&" + values.Encode()
		log.Println(q.Encode() +"&"+values.Encode(),req.Header)
	}else{
		req.URL.RawQuery = q.Encode()
		log.Println(q.Encode())
	}


}

//NewGetRequest makes the request with the headers
func NewGetRequest(url string,values url.Values)(*http.Request,error){
	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil {
		return nil, err
	}
	SetHeaders(req,values)
	return req,nil
}
