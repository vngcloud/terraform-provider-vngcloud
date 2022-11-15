/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdateL7PolicyRequest struct {
	// Action. This indicates how the listener will route traffic. The value can be REDIRECT_TO_POOL or REDIRECT_TO_URL.
	Action string `json:"action"`
	// Compare operation. The value can be CONTAINS or EQUAL_TO
	CompareType string `json:"compareType"`
	// Keep the query string or not.
	KeepQueryString bool `json:"keepQueryString,omitempty"`
	// Policy's id to be updated.
	L7policyId string `json:"l7policyId"`
	// Position of the policy
	Position int64 `json:"position,omitempty"`
	// Redirect HTTP code for redirecting to other URL.
	RedirectHttpCode int32 `json:"redirectHttpCode,omitempty"`
	// Pool for forwarding.
	RedirectPoolId string `json:"redirectPoolId,omitempty"`
	// URL for forwarding.
	RedirectUrl string `json:"redirectUrl,omitempty"`
	// Which attribute to compare. The value can be PATH or HOST_NAME
	Type_ string `json:"type"`
	// The value to compare with attribute.
	Value string `json:"value,omitempty"`
}
