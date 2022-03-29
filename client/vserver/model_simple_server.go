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

type SimpleServer struct {
	CreatedAt          time.Time                   `json:"createdAt,omitempty"`
	FlavorId           string                      `json:"flavorId,omitempty"`
	ImageId            string                      `json:"imageId,omitempty"`
	InternalInterfaces []InterfaceNetworkInterface `json:"internalInterfaces,omitempty"`
	Licence            bool                        `json:"licence,omitempty"`
	Name               string                      `json:"name,omitempty"`
	NetworkId          string                      `json:"networkId,omitempty"`
	ProjectId          string                      `json:"projectId,omitempty"`
	Status             string                      `json:"status,omitempty"`
	UpdatedAt          time.Time                   `json:"updatedAt,omitempty"`
	Uuid               string                      `json:"uuid,omitempty"`
}
