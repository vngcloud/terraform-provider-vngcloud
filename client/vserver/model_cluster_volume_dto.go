/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ClusterVolumeDto struct {
	ClusterId          int64  `json:"clusterId,omitempty"`
	DateCreated        string `json:"dateCreated,omitempty"`
	Iops               int32  `json:"iops,omitempty"`
	Name               string `json:"name,omitempty"`
	ProjectId          int64  `json:"projectId,omitempty"`
	Size               int64  `json:"size,omitempty"`
	Status             string `json:"status,omitempty"`
	VmUuid             string `json:"vmUuid,omitempty"`
	VolumeTypeId       string `json:"volumeTypeId,omitempty"`
	VolumeTypeZoneName string `json:"volumeTypeZoneName,omitempty"`
	VolumeUuid         string `json:"volumeUuid,omitempty"`
}
