package utils

import (
	"encoding/json"
	"fmt"

	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

// ErrorResponse this class unpack message from returned error from swagger library
type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func GetErrorMessage(e error) error {
	var resp ErrorResponse
	_, ok := e.(vserver.GenericSwaggerError)
	if !ok {
		return e
	}
	err := json.Unmarshal(e.(vserver.GenericSwaggerError).Body(), &resp)
	if err != nil {
		return err
	}
	return fmt.Errorf("%v%v", e, resp.Message)
}
