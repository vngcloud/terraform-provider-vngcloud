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

type ServerAction struct {
	Action     string    `json:"action,omitempty"`
	StartTime  time.Time `json:"startTime,omitempty"`
	UserAction string    `json:"userAction,omitempty"`
}
