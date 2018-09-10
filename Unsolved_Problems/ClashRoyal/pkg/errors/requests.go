package errors

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Reason string `json:"reason"`
	Message string `json:"message"`
	StatusCode int
}

func (err ResponseError) Error()string{
	return "API response error -> reason : "+err.Reason+" message : "+err.Message+" statusCode : "+strconv.Itoa(err.StatusCode)
}

//Checks the status code of the response and transforms it into an error type that correlates to the messages from the clash royale api
func CheckStatusCode(response *http.Response)error{

	if response.StatusCode==200{
		return nil
	}

	var respErr ResponseError
	json.NewDecoder(response.Body).Decode(&respErr)
	respErr.StatusCode = response.StatusCode
	log.Println(respErr.Error())
	return respErr

}