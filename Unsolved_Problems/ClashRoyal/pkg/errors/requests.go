package errors

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type ResponseError struct {
	Reason string `json:"reason"`
	Message string `json:"message"`
}

func CheckStatusCode(response *http.Response)error{

	if response.StatusCode==200{
		return nil
	}

	var respErr ResponseError
	json.NewDecoder(response.Body).Decode(&respErr)
	err := errors.Errorf("reason : %s\nmessage : %s",respErr.Reason,respErr.Message)
	log.Println(err)
	return err

}