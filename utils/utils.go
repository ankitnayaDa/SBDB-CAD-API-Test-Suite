package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)


const (
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
	defaultRetryMax     = 4
	getMethod           = "GET"
)

//GetResponse (url, token, "POST", payload) (string, error)
//Description: Generic method for connecting the given url with required parameters and returns the response body(Json string)
//Prints the error message and exists if there is an error.
//Returns: response body received from the url.
func RestClient(url, methodType string, payload string) (string, error) {
	var resp *http.Response

	req, err := http.NewRequest(methodType, url, strings.NewReader(payload))

	if err != nil {
		return "", err
	}

	proxyURL, err := http.ProxyFromEnvironment(req)

	if err != nil {
		return "", err
	}

	req.Header.Add("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport, Timeout: time.Minute*3 + time.Second*40}

	if req.Method == getMethod {
		resp, err = httpDo(client, req)
	} else {
		resp, err = client.Do(req)
	}

	if err != nil {
		return "", fmt.Errorf("error occurred while hitting request : %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	text := string(body)
	return text, err
}


func httpDo(client *http.Client, request *http.Request) (*http.Response, error) {
	i := 0

	for {
		var code int

		//Attempt to request
		response, err := client.Do(request)

		//Check if we want to continue retry
		checkOk, err := checkForRetry(response, err)

		if err != nil {
			fmt.Printf("[ERR] %s %s request failed: %v", request.Method, request.URL, err)
		}

		if !checkOk {
			return response, err
		}

		if response != nil {
			drainBody(response.Body)
		}

		remain := defaultRetryMax - i
		if remain == 0 {
			break
		}

		wait := backOff(defaultRetryWaitMin, defaultRetryWaitMax, i)
		desc := fmt.Sprintf("%s %s", request.Method, request.URL)
		if code > 0 {
			desc = fmt.Sprintf("%s %s", request.Method, request.URL)
		}

		fmt.Printf("[DEBUG] %s: retrying in %s (%d left)\n", desc, wait, remain)
		time.Sleep(wait)
		i++
	}

	return nil, fmt.Errorf("%s %s giving up after %d attempts", request.Method, request.URL, defaultRetryMax+1)
}

func checkForRetry(response *http.Response, err error) (bool, error) {

	if err != nil {
		return true, err
	}

	//Check the response code. retry if status code is range from 500
	if response.StatusCode == 0 || response.StatusCode >= 500 {
		return true, nil
	}

	return false, nil
}


func drainBody(body io.ReadCloser) {

	defer body.Close()

	data, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println("[Error] body read failed before draining", err)
	}

	fmt.Printf("draining the body received %s\n", string(data))

}

func backOff(min, max time.Duration, attemptNumber int) time.Duration {
	multi := math.Pow(2, float64(attemptNumber)) * float64(min)
	sleep := time.Duration(multi)

	if float64(sleep) != multi || sleep > max {
		sleep = max
	}

	return sleep
}

//PrettifyJSON (jsonString string)  string
//Description: Returns indented json string which can be used for logging
//Returns: Indented json string
func PrettifyJSON(jsonString string) string {
	if jsonString == "" {
		return ""
	}
	js, err := simplejson.NewJson([]byte(jsonString))
	if err != nil {
		return ""
	}
	prettyJSON, err := js.EncodePretty()
	if err != nil {
		return ""
	}
	return string(prettyJSON)
}

//ToJSONString (payload interface{})  string
//Description : Converts structure information to json string
//Returns: json string
func ToJSONString(payload interface{}) string {
	js, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return string(js)

}

func JsonUnmarshal(JSONStr string, unmarshalObj interface{}) error {
	err := json.Unmarshal([]byte(JSONStr), &unmarshalObj)

	if err != nil {
		return fmt.Errorf("json unmarshal Error: %s, response received : %s", err.Error(), JSONStr)
	}
	return nil
}
