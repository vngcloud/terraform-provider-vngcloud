/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type Engine struct {
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Image          string   `json:"image,omitempty"`
	EngineLicenses []string `json:"engineLicenses,omitempty"`
}