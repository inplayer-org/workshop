package parser

func TagToTagUrlTag(tag string) string{

		return "%25" + tag[1:]

}

func TagsToTagsUrlTags(tags []string)[]string {

	var ret []string

	for _,elem:=range tags{
		ret=append(ret,TagToTagUrlTag(elem))
	}
	return ret
}