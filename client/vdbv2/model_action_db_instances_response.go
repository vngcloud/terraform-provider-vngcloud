/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type ActionDbInstancesResponse struct {
	DatabaseInstances string `json:"databaseInstances,omitempty"`
	Action            string `json:"action,omitempty"`
	Status            string `json:"status,omitempty"`
	Success           bool   `json:"success,omitempty"`
	ErrorMsg          string `json:"errorMsg,omitempty"`
	Code              int32  `json:"code,omitempty"`
}