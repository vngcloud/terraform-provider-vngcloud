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
	return fmt.Sprint("Status Code: ", httpResponse.StatusCode, " , ", string(localVarBody))
}
