/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type PagingClusterDto struct {
	ListData  []ClusterDto `json:"listData,omitempty"`
	Page      int32        `json:"page,omitempty"`
	PageSize  int32        `json:"pageSize,omitempty"`
	TotalItem int64        `json:"totalItem,omitempty"`
	TotalPage int32        `json:"totalPage,omitempty"`
}
