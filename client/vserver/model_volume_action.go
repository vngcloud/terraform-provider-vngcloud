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

type VolumeAction struct {
	Iops  int32     `json:"iops,omitempty"`
	Size  int32     `json:"size,omitempty"`
	Start time.Time `json:"start,omitempty"`
	Type_ string    `json:"type,omitempty"`
}
