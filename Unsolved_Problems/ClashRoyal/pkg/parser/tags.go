package parser

//ToUrlTag - Removes the first character (should be #) and concatenates %25 in from for making a request from the Clash Royal api by tag for a single string entry
func ToUrlTag(tag string) string{

		return "%25" + tag[1:]

}

//ToUrlTags - Removes the first character (should be #) and concatenates %25 in from for making a request from the Clash Royal api by tag for a slice of strings
func ToUrlTags(tags []string)[]string {

	var ret []string

	for _,elem:=range tags{
		ret=append(ret,ToUrlTag(elem))
	}
	return ret
}