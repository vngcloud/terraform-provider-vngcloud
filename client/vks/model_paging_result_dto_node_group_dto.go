/*
 * vks-api API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0-SNAPSHOT
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vks

type PagingResultDtoNodeGroupDto struct {
	Items    []NodeGroupDto `json:"items,omitempty"`
	Total    int64          `json:"total,omitempty"`
	Page     int32          `json:"page,omitempty"`
	PageSize int32          `json:"pageSize,omitempty"`
}
