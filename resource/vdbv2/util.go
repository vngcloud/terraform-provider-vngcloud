package vdbv2

import (
	"fmt"
	"io"
	"net/http"
)

func CheckErrorResponse(httpResponse *http.Response) bool {
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		return true
	}
	return false
}

func GetResponseBody(httpResponse *http.Response) string {
	localVarBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return fmt.Sprint("Error parse response body: ", err)
	}
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
