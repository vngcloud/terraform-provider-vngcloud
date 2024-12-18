/*
 * vks-api API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0-SNAPSHOT
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vks

type CreateNodeGroupDto struct {
	Name                    string                       `json:"name"`
	NumNodes                int32                        `json:"numNodes"`
	AutoScaleConfig         *NodeGroupAutoScaleConfigDto `json:"autoScaleConfig,omitempty"`
	UpgradeConfig           NodeGroupUpgradeConfigDto    `json:"upgradeConfig,omitempty"`
	ImageId                 string                       `json:"imageId"`
	FlavorId                string                       `json:"flavorId"`
	DiskSize                int32                        `json:"diskSize"`
	DiskType                string                       `json:"diskType"`
	EnablePrivateNodes      bool                         `json:"enablePrivateNodes"`
	SecurityGroups          []string                     `json:"securityGroups"`
	SshKeyId                string                       `json:"sshKeyId"`
	Labels                  map[string]string            `json:"labels,omitempty"`
	Taints                  []NodeGroupTaintDto          `json:"taints,omitempty"`
	SecondarySubnets        []string                     `json:"secondarySubnets"`
	SubnetId                string                       `json:"subnetId,omitempty"`
	EnabledEncryptionVolume bool                         `json:"enabledEncryptionVolume"`
}
