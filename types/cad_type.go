package types

//400 Bad Response structure
type Check400BadResponse struct {
	MoreInfo string `json:"moreInfo"`
	Message  string `json:"message"`
	Code     string `json:"code"`
}

//405 Method Not Allowed Response structure
type CadApiResponse struct {
	Signature struct {
		Source  string `json:"source"`
		Version string `json:"version"`
	} `json:"signature"`
	Count  string     `json:"count"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
}