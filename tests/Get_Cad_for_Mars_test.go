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

	assert.Nil(t,libs.Get_Cad_Sort_Date(CAD_API),"UNEXPECTED:User not able to view All Close Apporch object data for mars based on specific dates with sort based on date ")
}

func Test_All_Cad_For_Mars_Based_Distance_Sort_Distance(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All Close Apporch object data for mars based on specific maximun and minimum distances with sort based on dist")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Mars&dist-max=10&dist-min=0.05&sort=dist"

	assert.Nil(t,libs.Get_Cad_Sort_Distance(CAD_API),"UNEXPECTED:User not able to view All Close Apporch object data for mars based on specific maximun and minimum distances with sort based on dist")
}

func Test_Zero_Count_Result_On_Mars(t *testing.T) {

	utils.LogInfo("Test description: Verify Zero count results in mars for class IEO")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Mars&class=IEO"

	assert.Nil(t,libs.Get_Zero_Count_Results_Mars(CAD_API),"UNEXPECTED:Data received")
}

func Test_All_NeaComet_For_Mars_Based_Distance(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All NEAs and comets for mars based on distance")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Mars&nea-comet=true&sort=dist"

	assert.Nil(t,libs.Get_Cad_Sort_Distance(CAD_API),"UNEXPECTED:User not able to view All NEAs and comets for mars based on distance")
}

func Test_All_Neo_For_Earth_Based_Distance(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All NEOS for earth sorted based on distance")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Earth&neo=true&sort=dist"

	assert.Nil(t,libs.Get_Cad_Sort_Distance(CAD_API),"UNEXPECTED:User not able to view All NEOS for earth sorted based on distance")
}

func Test_All_NO_For_Earth_Based_Distance(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All numbered-objects for earth sorted based on distance with filter based on dates")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=earth&kind=n&date-min=2020-01-01&date-max=2022-01-01&sort=date"

	assert.Nil(t,libs.Get_Cad_Sort_Date(CAD_API),"UNEXPECTED:User not able to view All numbered-objects for earth sorted based on distance with filter based on dates")
}

func Test_Particular_Object_For_Earth_Based_Distance(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view details for neo 2020 QW3 for earth sorted based on distance")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"des=2020 QW3&date-min=1900-01-01&date-max=2100-01-01&dist-max=0.2&sort=dist"

	assert.Nil(t,libs.Get_Cad_Sort_Distance(CAD_API),"UNEXPECTED:User not able to view details for neo 2020 QW3 for earth sorted based on distance")
}

func Test_Particular_Object_For_Earth_Limit(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view  details for neo 2020 QW3 for earth with limit 2")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"des=2020 QW3&date-min=1900-01-01&date-max=2100-01-01&dist-max=0.2"

	assert.Nil(t,libs.Get_All_Cad_With_Limit(CAD_API,"2"),"UNEXPECTED:User not able to view  details for neo 2020 QW3 for earth with limit 2")
}

func Test_Particular_Object_For_Earth_With_Diamater_Fullname(t *testing.T) {

	utils.LogInfo("Test description: Verify user able to view All NEOS for earth with diameter and fullname")
	utils.LogInfo("***** Execution started *****")

	CAD_API := types.CadAPI+"body=Earth&neo=true&diameter=true&fullname=true"

	assert.Nil(t,libs.Get_Cad_With_Diamater_Fullname(CAD_API),"UNEXPECTED:User not able to view All NEOS for earth with diameter and fullname")
}