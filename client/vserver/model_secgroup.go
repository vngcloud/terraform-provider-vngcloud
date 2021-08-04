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

type Secgroup struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayCreatedAt string `json:"displayCreatedAt,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	System bool `json:"system,omitempty"`
}
