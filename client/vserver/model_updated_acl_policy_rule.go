/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdatedAclPolicyRule struct {
	Action                 string `json:"action,omitempty"`
	InterfaceAclPolicyUuid string `json:"interfaceAclPolicyUuid,omitempty"`
	Port                   string `json:"port,omitempty"`
	Protocol               string `json:"protocol,omitempty"`
	SeqNumber              int32  `json:"seqNumber,omitempty"`
	Source                 string `json:"source,omitempty"`
	System                 bool   `json:"system,omitempty"`
	Type_                  string `json:"type,omitempty"`
}
