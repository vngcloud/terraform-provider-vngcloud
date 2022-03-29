/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type SimpleVolume struct {
	BootIndex    int32     `json:"bootIndex,omitempty"`
	Bootable     bool      `json:"bootable,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	Name         string    `json:"name,omitempty"`
	ProjectId    string    `json:"projectId,omitempty"`
	ServerId     string    `json:"serverId,omitempty"`
	Size         int32     `json:"size,omitempty"`
	Status       string    `json:"status,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	Uuid         string    `json:"uuid,omitempty"`
	VolumeTypeId string    `json:"volumeTypeId,omitempty"`
}
