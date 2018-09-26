package errors

import (
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

type ResponseError struct {
	Reason string `json:"reason"`
	Message string `json:"message"`
	StatusCode int
}

//String message of the error (implementing of the error interface)
func (err ResponseError) Error()string{
	return "API response error -> reason : "+err.Reason+" message : "+err.Message+" statusCode : "+strconv.Itoa(err.StatusCode)
}

//Constructor for ResponseError
func NewResponseError(reason string, message string,code int)error{
	return &ResponseError{reason,message,code}
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
