/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type InterfaceClusterNodeGroupEntity struct {
	BackendStatus      string    `json:"backendStatus,omitempty"`
	ClusterId          string    `json:"clusterId,omitempty"`
	ClusterNodeGroupId int32     `json:"clusterNodeGroupId,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	DeletedAt          time.Time `json:"deletedAt,omitempty"`
	FlavorId           string    `json:"flavorId,omitempty"`
	Id                 int32     `json:"id,omitempty"`
	Name               string    `json:"name,omitempty"`
	NodeCount          int32     `json:"nodeCount,omitempty"`
	NodeGroupDefault   bool      `json:"nodeGroupDefault,omitempty"`
	ProjectId          string    `json:"projectId,omitempty"`
	Status             string    `json:"status,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	Uuid               string    `json:"uuid,omitempty"`
}