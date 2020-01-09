package fxn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Verify(ref string) bool {
	url := "https://ravesandboxapi.flutterwave.com/flwv3-pug/getpaidx/api/v2/verify" //Please make sure to change to that of production server when you are ready to go live.

	data := make(map[string]string)
	data["txref"] = ref
	// merchant secret key
	data["SECKEY"] = "FLWSECK_TEST-01783ff073eee2afa44b3387283324ae-X"

	//convert data map to json byte[]
	jsonData, _ := json.Marshal(data)

	// make request to endpoint
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	// Make `POST` request and handle response
	response, err := client.Do(request)

	if err == nil && response != nil {
		responseJSON, _ := ioutil.ReadAll(response.Body)
		var responseData map[string]interface{}
		err = json.Unmarshal(responseJSON, &responseData)

		if err == nil {
			fmt.Println(responseData)

			if strings.Compare(responseData["status"].(string), "successful") == 0 {
				amount, _ := responseData["data"].(map[string]interface{})["amount"].(float64)
				if amount == 1000 {
					fmt.Println(true)
					return true
				}
			}
		}
	}
	return false
}
