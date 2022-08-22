package tests

import (
	"SBDB-CAD-API-Test-Suite/libs"
	"SBDB-CAD-API-Test-Suite/types"
	"SBDB-CAD-API-Test-Suite/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Check_400_Response(t *testing.T) {
	utils.LogInfo("Test description: Test for 400 bad response when request contained invalid keywords and/or content (details returned in the JSON payload)")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"dist-max=10LD&date-min=2018-01-01&sort=dist&dist=10LD"

	Check400BadResponseData := types.Check400BadResponse{
		MoreInfo:"https://ssd-api.jpl.nasa.gov/doc/cad.html",
		Message:"one or more query parameter was not recognized",
		Code:"400"}

	assert.Nil(t,libs.Verify_400_Response(CAD_API,Check400BadResponseData),"UNEXPECTED:400 BAD REQUEST Not Received")
}

func Test_Check_405_Response(t *testing.T) {
	utils.LogInfo("Test description: Verify 405 response when the request used an incorrect method (see the HTTP Request section)")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"dist-max=10LD&date-min=2018-01-01&sort=dist"

	Check400BadResponseData := types.Check400BadResponse{
		MoreInfo:"https://ssd-api.jpl.nasa.gov/doc/cad.html",
		Message:"Invalid HTTP method: must be GET or POST",
		Code:"405"}

	assert.Nil(t,libs.Verify_405_Response(CAD_API,Check400BadResponseData),"UNEXPECTED:405 Not Received")
}

