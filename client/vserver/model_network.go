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

type Network struct {
	Cidr              string            `json:"cidr,omitempty"`
	CreatedAt         time.Time         `json:"createdAt,omitempty"`
	DisplayCreatedAt  string            `json:"displayCreatedAt,omitempty"`
	ElasticIpEntities []ElasticIpEntity `json:"elasticIpEntities,omitempty"`
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Status            string            `json:"status,omitempty"`
	Subnets           []Subnet          `json:"subnets,omitempty"`
}
