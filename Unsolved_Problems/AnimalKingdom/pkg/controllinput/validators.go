package controllinput

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
	"strconv"
	resp "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
	. "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/consts"
)

func ValidateStructureData(food structures.Food,w http.ResponseWriter) bool{
	if food.FoodID==0 || !CheckString(&food.Name) || !CheckString(&food.Type){
		resp.RespondWithError(w,http.StatusBadRequest,responses.InvalidStructureData())
		return false
	}
	return true
}

func ValidateEntry(food structures.Food,entry string,w http.ResponseWriter)string{
	if IntOnly(entry) {
		if entry != strconv.Itoa(food.FoodID) {
			resp.RespondWithError(w, http.StatusMultipleChoices, responses.MultipleChoices("index",entry,strconv.Itoa(food.FoodID)))
		} else {
			return Int
		}

	} else if CheckString(&entry) {
		if entry != food.Name {
			resp.RespondWithError(w, http.StatusMultipleChoices, responses.MultipleChoices("name",entry,food.Name))

		} else {
			return String
		}

	} else {
		resp.RespondWithError(w, http.StatusBadRequest, responses.MixedRequest())
	}
	return ""
}

func ValidateRequest(request string,w http.ResponseWriter)(interface{},string){
	if IntOnly(request) {
		intRequest,_ := strconv.Atoi(request)
			return intRequest,Int
		} else if CheckString(&request) {
			return request,String
		}else {
		resp.RespondWithError(w, http.StatusBadRequest, responses.MixedRequest())
	}
	return "",""
}
