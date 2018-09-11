package errors

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type ResponseError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

//Checks the status code of the response and transforms it into an error type that correlates to the messages from the clash royale api
func CheckStatusCode(response *http.Response) error {

	if response.StatusCode == 200 {
		return nil
	}

	var respErr ResponseError
	json.NewDecoder(response.Body).Decode(&respErr)
	err := errors.Errorf("reason : %s\nmessage : %s", respErr.Reason, respErr.Message)
	log.Println(err)
	return err

}
