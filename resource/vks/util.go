package vks

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"io/ioutil"
	"net/http"
)

func CheckErrorResponse(httpResponse *http.Response) bool {
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		return true
	}
	return false
}

func GetResponseBody(httpResponse *http.Response) string {
	localVarBody, _ := io.ReadAll(httpResponse.Body)
	responseMessage := string(localVarBody)
	if httpResponse.StatusCode == 403 {
		responseMessage = "You don't have permission to do this action"
	}
	return fmt.Sprint("Status Code: ", httpResponse.StatusCode, ", ", responseMessage)
}

func CheckContainString(list []string, findElement string) bool {
	for _, element := range list {
		if element == findElement {
			return true
		}
	}
	return false
}

func CheckListStringEqual(firstList []string, secondList []string) bool {
	if len(firstList) == len(secondList) {
		for _, element := range firstList {
			if !CheckContainString(secondList, element) {
				return false
			}
		}
		return true
	}
	return false
}

func MergeSchemas(a, b map[string]*schema.Schema) map[string]*schema.Schema {
	merged := make(map[string]*schema.Schema)

	for k, v := range a {
		merged[k] = v
	}

	for k, v := range b {
		merged[k] = v
	}

	return merged
}

func fetchByKey(key string) (interface{}, error) {
	url := "https://terraform.api.vngcloud.vn/config.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	response, ok := result[key].(string)
	if !ok {
		return nil, fmt.Errorf("%s not found in response", key)
	}

	return response, nil
}
