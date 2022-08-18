package tests

import (
	"SBDB-CAD-API-Test-Suite/libs"
	"SBDB-CAD-API-Test-Suite/types"
	"SBDB-CAD-API-Test-Suite/utils"
	"testing"
)

/*############################################################################################################
        NDAC_TC_TAGS: NDAC_CLD_SM, ALL
        AUTHOR      : sunil.kumar_k@nokia.com
        DESCRIPTION : Test Add duplicate IMSIs for an email id
##############################################################################################################*/

func Test_Get_Cad_For_Asteroid(t *testing.T) {

	utils.LogInfo("Test desription: Test Add duplicate IMSIs for an email id")
	utils.LogInfo("***** Execution started *****")

	libs.Get_All_Cad_For_Asteroid(types.CadAPI)

}
