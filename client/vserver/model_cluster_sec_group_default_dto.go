/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ClusterSecGroupDefaultDto struct {
	Master *SecGroupDefault `json:"master,omitempty"`
	Minion *SecGroupDefault `json:"minion,omitempty"`
}
