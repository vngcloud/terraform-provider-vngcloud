/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdateServerGroupRequestV2 struct {
	// description for  Server Group.
	Description string `json:"description,omitempty"`
	// Name of the Server Group
	Name string `json:"name"`
	// Id of the server group
	ServerGroupId string `json:"serverGroupId"`
}
