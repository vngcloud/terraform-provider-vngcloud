/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type NetworkAclModel struct {
	AclPolicyId           string          `json:"aclPolicyId,omitempty"`
	AclPolicyRules        []AclPolicyRule `json:"aclPolicyRules,omitempty"`
	CreatedAt             *Timestamp      `json:"createdAt,omitempty"`
	DefaultAcl            bool            `json:"defaultAcl,omitempty"`
	InterfaceNetworkUuid  string          `json:"interfaceNetworkUuid,omitempty"`
	Name                  string          `json:"name,omitempty"`
	ProjectUuid           string          `json:"projectUuid,omitempty"`
	Status                string          `json:"status,omitempty"`
	SubnetAssociationList []string        `json:"subnetAssociationList,omitempty"`
	Uuid                  string          `json:"uuid,omitempty"`
}
