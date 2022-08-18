package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//GetRestClientResponse (url string,requestType string, payload string) (string, error)
//Description: makes http request and gets response
//Returns: raw response of the request made , error
func GetRestClientResponse(url string, requestType string, payload string, args ...bool) (string, error) {

	LogInfo("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	LogInfo("+                    Request                       +")
	LogInfo("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	LogInfo("API: ", url)
	LogInfo("Request type: ", requestType)
	LogInfo("Payload:\n", PrettifyJSON(payload))

	response, err := RestClient(url, requestType, payload)

	if err != nil {
		return "", err
	}

	LogInfo("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	LogInfo("+                    Response                      +")
	LogInfo("++++++++++++++++++++++++++++++++++++++++++++++++++++\n", PrettifyJSON(response)+"\n")
	return response, err
}

//logMessage(flag string, message ...interface{})
//Description: logs the message in format dateAndTime::messageLevel::CallerScriptName::lineNumber::message
//Returns: nil
func logMessage(flag string, message ...interface{}) {
	_, fileName, lineNumber, ok := runtime.Caller(2) // 2 specifies level 2 in stack trace, returns (function_pointer, file_name, line_number, ok)
	if ok {
		message[0] = time.Now().UTC().Format("2006/01/02 15:04:05.999Z") + "::" + flag + "::" + filepath.Base(fileName) + "::" + strconv.Itoa(lineNumber) + "::" + fmt.Sprint(message[0])
	}
	if strings.ContainsAny(fmt.Sprint(message[0]), "%") {
		log.Printf(fmt.Sprint(message[0]), message[1:]...)
	} else {
		log.Println(message...)
	}
}

//LogInfo (message ...interface{})
//Description: Logs info messages
//Returns: nil
func LogInfo(message ...interface{}) {
	logMessage("INFO", message...)
}

//LogWarning (message ...interface{})
//Description: Logs warning messages
//Returns: nil
func LogWarning(message ...interface{}) {
	logMessage("WARNING", message...)
}

//LogError (message ...interface{})
//Description: Logs error message and exists.
//Returns: nil
func LogError(message ...interface{}) {
	logMessage("ERROR", message...)
}


