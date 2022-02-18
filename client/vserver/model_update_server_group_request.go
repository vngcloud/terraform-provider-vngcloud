/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdateServerGroupRequest struct {
	// Description of server group
	Description string `json:"description,omitempty"`
	// Id of project
	ProjectId string `json:"projectId,omitempty"`
	// Id of the server group
	ServerGroupId string `json:"serverGroupId"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
}