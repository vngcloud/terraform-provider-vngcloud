/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type PersistentVolumeDto struct {
	ClusterId string `json:"clusterId,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
	IopsId    int64  `json:"iopsId,omitempty"`
	Name      string `json:"name,omitempty"`
	Size      int32  `json:"size,omitempty"`
	Status    string `json:"status,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	VmId      string `json:"vmId,omitempty"`
	VolumeId  string `json:"volumeId,omitempty"`
}