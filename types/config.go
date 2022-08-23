package types

// CadAPI : SBDB Close-Approach Data API
var CadAPI = "https://ssd-api.jpl.nasa.gov/cad.api?"

//400 Bad Response Data
var Check400BadResponseData = Check400BadResponse{
	MoreInfo: "https://ssd-api.jpl.nasa.gov/doc/cad.html",
	Message:  "one or more query parameter was not recognized",
	Code:     "400"}

//405 Incorrect method Data
var Check405BadResponseData = Check400BadResponse{
	MoreInfo: "https://ssd-api.jpl.nasa.gov/doc/cad.html",
	Message:  "Invalid HTTP method: must be GET or POST",
	Code:     "405"}
