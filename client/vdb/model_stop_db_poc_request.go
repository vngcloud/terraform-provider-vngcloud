/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vdb

type StopDbPocRequest struct {
	DbInstanceId string      `json:"dbInstanceId,omitempty"`
	Extra        interface{} `json:"extra,omitempty"`
	ProjectId    string      `json:"projectId,omitempty"`
}