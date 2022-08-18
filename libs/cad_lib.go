package libs

import (
	"SBDB-CAD-API-Test-Suite/types"
	"SBDB-CAD-API-Test-Suite/utils"
)

//get all close-approach data for asteroid
func Get_All_Cad_For_Asteroid( API string) error {
	_, err := utils.GetRestClientResponse(types.CadAPI, "GET", "")
	if err != nil {
		return err
	}
	return nil
}
