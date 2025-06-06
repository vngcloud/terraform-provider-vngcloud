/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

import (
	"time"
)

type NetworkResponseV2 struct {
	Uuid        string           `json:"uuid,omitempty"`
	CreatedAt   time.Time        `json:"createdAt,omitempty"`
	Status      string           `json:"status,omitempty"`
	DisplayName string           `json:"displayName,omitempty"`
	NetworkId   int32            `json:"networkId,omitempty"`
	Subnets     []SubnetResponse `json:"subnets,omitempty"`
}
