/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type RootDiskLimit struct {
	MaxIOPS int32 `json:"maxIOPS,omitempty"`
	MaxSize int32 `json:"maxSize,omitempty"`
	MinSize int32 `json:"minSize,omitempty"`
}