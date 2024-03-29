/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Resize Volume Request
type ResizeVolumeRequest struct {
	// New size of volume
	NewSize int32 `json:"newSize"`
	// Id of new volume type
	NewVolumeTypeId string `json:"newVolumeTypeId"`
}
