package libs

import (
	"SBDB-CAD-API-Test-Suite/types"
	"SBDB-CAD-API-Test-Suite/utils"
	"errors"
	"reflect"
	"strconv"
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

func Get_Zero_Count_Results_Mars( API string) error {
	resposebody, err := utils.GetRestClientResponse(API, "GET", "")
	if err != nil {
		return err
	}
	var unmarshalledActualResponse types.CadApiResponse
	if err := utils.JsonUnmarshal(resposebody, &unmarshalledActualResponse); err != nil {
		return errors.New("Error in unmarshalling")
	}
	if unmarshalledActualResponse.Count == "0"{
		return nil
	} else {
		return errors.New("Objects found in Mars Orbit")
	}
}

//get all close-approach data for Mars
func Get_Cad_Sort_Date( API string) error {
	API = strings.Replace(API," ","%20",-1)
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
func Get_Cad_Sort_Distance( API string) error {
	API = strings.Replace(API," ","%20",-1)
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
	return CheckSortBasedOnDistance(unmarshalledActualResponse)
}

func Get_All_Cad_With_Limit( API string,limit string) error {
	API = strings.Replace(API," ","%20",-1)
	API = API + "&limit="+limit
	resposebody, err := utils.GetRestClientResponse(API, "GET", "")
	if err != nil {
		return err
	}
	var unmarshalledActualResponse types.CadApiResponse
	if err := utils.JsonUnmarshal(resposebody, &unmarshalledActualResponse); err != nil {
		return errors.New("Error in unmarshalling")
	}
	Limit := strconv.Itoa(len(unmarshalledActualResponse.Data))
	if unmarshalledActualResponse.Count == "0" {
		return errors.New("No Objects found in Mars Orbit")
	} else if Limit != limit {
		utils.LogInfo("Data limit did not match",Limit, limit)
		return errors.New("Data limit did not match")
	}
	return nil
}

//get all close-approach data for Mars
func Get_Cad_With_Diamater_Fullname( API string) error {
	API = strings.Replace(API," ","%20",-1)
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
	return CheckDiamaterWithFullname(unmarshalledActualResponse)
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


func CheckSortBasedOnDistance(response types.CadApiResponse) error {
	var old_distance float64
	for i, t := range response.Data[0]{
		if i == 4{
			old_distance, _ := strconv.ParseFloat(t, 32)
			utils.LogInfo("Received 1st distance",old_distance)

		}
	}
	for _, t := range response.Data {
		distance, _ := strconv.ParseFloat(t[4], 32)
		received_distance := distance
		utils.LogInfo("Compairing new distance %f with old distance %f ", old_distance,received_distance)
		if received_distance < old_distance  {
			utils.LogInfo("Data is not sorted based on distance")
			return errors.New("data is not sorted")
		}
		old_distance = distance

	}
	utils.LogInfo("Data is sorted based on distance")
	return nil
}

func CheckDiamaterWithFullname(response types.CadApiResponse) error {
	for _, t := range response.Data {
		if t[13] == ""{
			return errors.New("object fullname not avilable")
		}
		utils.LogInfo(t[11],t[12])
		if t[11] == "" && t[12] == "" {
			return errors.New("object diameter not avilable")
		}
	}
	return nil
}