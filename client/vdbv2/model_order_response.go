/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type OrderResponse struct {
	OrderUrl string `json:"orderUrl,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
}
