package vserver

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
	localVarBody, _ := io.ReadAll(httpResponse.Body)
	responseMessage := string(localVarBody)
	if httpResponse.StatusCode == 403 {
		responseMessage = "You don't have permission to do this action"
	}
	return fmt.Sprint("Status Code: ", httpResponse.StatusCode, ", ", responseMessage)
}
