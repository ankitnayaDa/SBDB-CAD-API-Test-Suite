package libs

import (
	"SBDB-CAD-API-Test-Suite/types"
	"SBDB-CAD-API-Test-Suite/utils"
	"errors"
	"reflect"
	"strings"
	"time"
)

//get all close-approach data for Mars
func Get_All_Cad_For_Mars( API string) error {
	resposebody, err := utils.GetRestClientResponse(API, "GET", "")
	if err != nil {
		return err
	}
	var unmarshalledActualResponse types.CadApiResponse
	if err := utils.JsonUnmarshal(resposebody, &unmarshalledActualResponse); err != nil {
		return errors.New("Error in unmarshalling")
	}
	if unmarshalledActualResponse.Count == "0"{
		return errors.New("No Objects found in Mars Orbit")
	}
	return nil
}

//get all close-approach data for Mars
func Get_Cad_For_Mars_Based_Date_Sort_Date( API string) error {
	resposebody, err := utils.GetRestClientResponse(API, "GET", "")
	if err != nil {
		return err
	}
	var unmarshalledActualResponse types.CadApiResponse
	if err := utils.JsonUnmarshal(resposebody, &unmarshalledActualResponse); err != nil {
		return errors.New("Error in unmarshalling")
	}
	if unmarshalledActualResponse.Count == "0"{
		return errors.New("No Objects found in Mars Orbit")
	} else {
		return CheckSortBasedOnTimeStamp(unmarshalledActualResponse)
	}
}

//get all close-approach data for Mars
func Get_Cad_For_Mars_Based_Distance_Sort_Distance( API string) error {
	resposebody, err := utils.GetRestClientResponse(API, "GET", "")
	if err != nil {
		return err
	}
	var unmarshalledActualResponse types.CadApiResponse
	if err := utils.JsonUnmarshal(resposebody, &unmarshalledActualResponse); err != nil {
		return errors.New("Error in unmarshalling")
	}
	if unmarshalledActualResponse.Count == "0"{
		return errors.New("No Objects found in Mars Orbit")
	}
	return nil
}


//Check For 400 Bad Response
func Verify_400_Response( API string ,expectedResponse interface{}) error {
	responsebody, err := utils.GetRestClientResponse(API, "GET", "")
	if err != nil {
		return err
	}
	return compareResponse(responsebody,expectedResponse)
}

//Check For 405 Method Not Allowed response
func Verify_405_Response( API string ,expectedResponse interface{}) error {
	responsebody, err := utils.GetRestClientResponse(API, "PUT", "")
	if err != nil {
		return err
	}
	return compareResponse(responsebody,expectedResponse)
}

//Verify : verifies if actualResponse and expectedResponse are same
func compareResponse(actualResponse string, expectedResponse interface{}) error {
	switch expectedResponse.(type) {
	case types.Check400BadResponse:
		var unmarshalledActualResponse types.Check400BadResponse
		if err := utils.JsonUnmarshal(actualResponse, &unmarshalledActualResponse); err != nil {
			return errors.New("Error in unmarshalling")
		}
		return CHECKIsEqual(unmarshalledActualResponse, expectedResponse)
	default:
		return errors.New("Unknown response type")
	}
}


func CHECKIsEqual(actual interface{}, expected interface{}) error {
	if !reflect.DeepEqual(actual, expected) {
		return errors.New("Expected reponse was :" + utils.ToJSONString(expected))
	}
	return nil
}

func CheckSortBasedOnTimeStamp(response types.CadApiResponse) error {
	var times time.Time
	for i, t := range response.Data[0]{
		if i == 3{
			splittedString := strings.Split(t, " ")
			times, _ = time.Parse("2006-Jan-02", splittedString[0])
			utils.LogInfo("Received 1st time stamp",times)

		}
	}
	for _, t := range response.Data {
			splittedString := strings.Split(t[3], " ")
			parseime, _ := time.Parse("2006-Jan-02", splittedString[0])
			utils.LogInfo("Compairing Date %s with date %s ", parseime,times)
			if times.After(parseime) {
				utils.LogInfo("Data is not sorted based on timestamp")
				return errors.New("data is not sorted")
			}
			times, _ = time.Parse("2006-Jan-02", splittedString[0])
		}
			utils.LogInfo("Data is sorted based on timestamp")
		return nil
}
