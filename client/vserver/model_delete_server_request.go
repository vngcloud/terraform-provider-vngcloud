/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Delete Server Request
type DeleteServerRequest struct {
	// Delete all volume are attached
	DeleteAllVolume bool `json:"deleteAllVolume,omitempty"`
}
