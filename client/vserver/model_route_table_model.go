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

type RouteTableModel struct {
	CreatedAt time.Time    `json:"createdAt,omitempty"`
	Name      string       `json:"name,omitempty"`
	Routes    []RouteModel `json:"routes,omitempty"`
	Status    string       `json:"status,omitempty"`
	Subnets   []Subnet     `json:"subnets,omitempty"`
	System    bool         `json:"system,omitempty"`
	Uuid      string       `json:"uuid,omitempty"`
	VpcUuid   string       `json:"vpcUuid,omitempty"`
}
