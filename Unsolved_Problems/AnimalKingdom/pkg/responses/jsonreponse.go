package responses

import (
	"net/http"
	"log"
	"encoding/json"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, dataStruct interface{}) {
	log.Println(dataStruct)
	response, err := json.Marshal(dataStruct)
	if err!=nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
	}else{
		w.WriteHeader(code)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}