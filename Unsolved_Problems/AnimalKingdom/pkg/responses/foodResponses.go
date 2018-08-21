package responses

func InvalidStructureData() string{
	return "Invalid structure data," +
		" Food name and type has to be between 2 and 30 characters and can't contain numbers." +
		" Food id shouldn't be lower than 1"
}

func MixedRequest() string{
	return "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters" +
		". Please provide correct request"
}


//MultipleChoices - tableRow string, entry string, structureValue string
func MultipleChoices(tableRow string, entry string, structureValue string) string{
	return "Entry "+tableRow+" ("+entry+") and JSON "+tableRow+" ("+structureValue+") has to be equal"

}

func AlreadyExist(tableRow string ,entry string)string{
	return "Food with "+tableRow+" ("+entry+") already exists in database, If you want to update entry use the PUT method"
}


func NotAllowedToUpdateBy(tableRow string)string{
	return "You aren't allowed to update the database by Food " + tableRow
}

func NotFound(tableRow string,entry string)string{
	return "Food "+ tableRow +" ("+entry+") not present in database"
}

func ErrorDuringExec(queryName string)string{
	return "Something went wrong during "+queryName+" in the database"
}