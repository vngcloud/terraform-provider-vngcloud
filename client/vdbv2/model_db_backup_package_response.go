/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type DbBackupPackageResponse struct {
	EngineGroup int32                   `json:"engineGroup,omitempty"`
	Packages    []DbBackupPackageDetail `json:"packages,omitempty"`
}