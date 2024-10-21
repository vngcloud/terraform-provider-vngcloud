/*
 * vks-api API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0-SNAPSHOT
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vks

type CreateClusterComboDto struct {
	Name                       string               `json:"name"`
	Description                string               `json:"description,omitempty"`
	Version                    string               `json:"version"`
	EnablePrivateCluster       bool                 `json:"enablePrivateCluster"`
	EnabledServiceEndpoint     bool                 `json:"enabledServiceEndpoint"`
	NetworkType                string               `json:"networkType"`
	VpcId                      string               `json:"vpcId"`
	SubnetId                   string               `json:"subnetId"`
	Cidr                       string               `json:"cidr"`
	EnabledLoadBalancerPlugin  bool                 `json:"enabledLoadBalancerPlugin"`
	EnabledBlockStoreCsiPlugin bool                 `json:"enabledBlockStoreCsiPlugin"`
	NodeGroups                 []CreateNodeGroupDto `json:"nodeGroups"`
	SecondarySubnets           []string             `json:"secondarySubnets"`
	NodeNetmaskSize            int32                `json:"nodeNetmaskSize"`
}
