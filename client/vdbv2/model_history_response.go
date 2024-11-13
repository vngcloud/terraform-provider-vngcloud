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

type HistoryResponse struct {
	Id           int32     `json:"id,omitempty"`
	InstanceId   string    `json:"instanceId,omitempty"`
	ProjectId    string    `json:"projectId,omitempty"`
	Action       string    `json:"action,omitempty"`
	Status       string    `json:"status,omitempty"`
	Description  string    `json:"description,omitempty"`
	UserId       int32     `json:"userId,omitempty"`
	CreatedTime  time.Time `json:"createdTime,omitempty"`
	UpdatedTime  time.Time `json:"updatedTime,omitempty"`
	ErrorMessage string    `json:"errorMessage,omitempty"`
}