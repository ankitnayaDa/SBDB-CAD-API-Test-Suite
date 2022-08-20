package tests

import (
	"SBDB-CAD-API-Test-Suite/libs"
	"SBDB-CAD-API-Test-Suite/types"
	"SBDB-CAD-API-Test-Suite/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_All_Cad_For_Mars(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All Close Apporch object data for mars")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Mars"

	assert.Nil(t,libs.Get_All_Cad_For_Mars(CAD_API),"UNEXPECTED:User not able to view All Close Apporch object data for mars")
}

func Test_All_Cad_For_Mars_Based_Date_Sort_Date(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All Close Apporch object data for mars based on specific dates with sort based on date")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Mars&date-min=2020-01-01&date-max=2022-01-01&sort=date"

	assert.Nil(t,libs.Get_Cad_For_Mars_Based_Date_Sort_Date(CAD_API),"UNEXPECTED:User not able to view All Close Apporch object data for mars based on specific dates with sort based on date ")
}

func Test_All_Cad_For_Mars_Based_Distance_Sort_Distance(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All Close Apporch object data for mars based on specific maximun and minimum distances with sort based on dist")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Mars&dist-max=10&dist-min=0.05&sort=dist"

	assert.Nil(t,libs.Get_Cad_For_Mars_Based_Distance_Sort_Distance(CAD_API),"UNEXPECTED:User not able to view All Close Apporch object data for mars based on specific maximun and minimum distances with sort based on dist")
}