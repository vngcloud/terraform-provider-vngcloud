/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Rename Server Request
type RenameServerRequest struct {
	// New name of server
	NewName string `json:"newName"`
	// Id of server
	ServerId string `json:"serverId"`
}
