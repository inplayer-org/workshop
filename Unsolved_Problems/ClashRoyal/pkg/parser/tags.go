package parser

func ToUrlTag(tag string) string{

		return "%25" + tag[1:]

}

func ToUrlTags(tags []string)[]string {

	var ret []string

	for _,elem:=range tags{
		ret=append(ret,ToUrlTag(elem))
	}
	return ret
}