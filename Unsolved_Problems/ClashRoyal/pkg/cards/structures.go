package cards


type Cards struct {
	Items []struct {
		Name     string `json:"name"`
		ID       int    `json:"id"`
		MaxLevel int    `json:"maxLevel"`
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"items"`
}



