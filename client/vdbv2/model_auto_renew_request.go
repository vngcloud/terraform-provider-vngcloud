/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type AutoRenewRequest struct {
	Product       string         `json:"product,omitempty"`
	ArtifactType  string         `json:"artifactType,omitempty"`
	ArtifactId    string         `json:"artifactId,omitempty"`
	AutoRenewInfo *AutoRenewInfo `json:"autoRenewInfo,omitempty"`
}
