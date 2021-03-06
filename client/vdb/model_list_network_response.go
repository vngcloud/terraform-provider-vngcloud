/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vdb

type ListNetworkResponse struct {
	Code        int32           `json:"code,omitempty"`
	ErrorMsg    string          `json:"errorMsg,omitempty"`
	ListNetwork []NetworkV2Info `json:"listNetwork,omitempty"`
	Success     bool            `json:"success,omitempty"`
	Total       int32           `json:"total,omitempty"`
}
