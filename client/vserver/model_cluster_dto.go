/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type ClusterDto struct {
	AclId                       string    `json:"aclId,omitempty"`
	AutoHealingEnabled          bool      `json:"autoHealingEnabled,omitempty"`
	AutoMonitoringEnabled       bool      `json:"autoMonitoringEnabled,omitempty"`
	AutoScalingEnabled          bool      `json:"autoScalingEnabled,omitempty"`
	BootVolumeSize              int32     `json:"bootVolumeSize,omitempty"`
	BootVolumeTypeId            string    `json:"bootVolumeTypeId,omitempty"`
	CalicoCidr                  string    `json:"calicoCidr,omitempty"`
	ClusterId                   int32     `json:"clusterId,omitempty"`
	CreatedAt                   time.Time `json:"createdAt,omitempty"`
	Description                 string    `json:"description,omitempty"`
	DockerVolumeSize            int32     `json:"dockerVolumeSize,omitempty"`
	DockerVolumeTypeId          string    `json:"dockerVolumeTypeId,omitempty"`
	EnabledLb                   bool      `json:"enabledLb,omitempty"`
	Endpoint                    string    `json:"endpoint,omitempty"`
	EtcdVolumeSize              int32     `json:"etcdVolumeSize,omitempty"`
	EtcdVolumeTypeId            string    `json:"etcdVolumeTypeId,omitempty"`
	IngressControllerEnabled    bool      `json:"ingressControllerEnabled,omitempty"`
	K8sNetworkType              string    `json:"k8sNetworkType,omitempty"`
	K8sNetworkTypeId            string    `json:"k8sNetworkTypeId,omitempty"`
	K8sVersion                  string    `json:"k8sVersion,omitempty"`
	K8sVersionId                string    `json:"k8sVersionId,omitempty"`
	MasterClusterSecGroupIdList []string  `json:"masterClusterSecGroupIdList,omitempty"`
	MasterCount                 int32     `json:"masterCount,omitempty"`
	MasterFlavorId              string    `json:"masterFlavorId,omitempty"`
	MasterFlavorName            string    `json:"masterFlavorName,omitempty"`
	MasterInstanceTypeId        string    `json:"masterInstanceTypeId,omitempty"`
	MaxNodeCount                int32     `json:"maxNodeCount,omitempty"`
	MinNodeCount                int32     `json:"minNodeCount,omitempty"`
	MinionClusterSecGroupIdList []string  `json:"minionClusterSecGroupIdList,omitempty"`
	Name                        string    `json:"name,omitempty"`
	NetworkId                   string    `json:"networkId,omitempty"`
	NodeCount                   int32     `json:"nodeCount,omitempty"`
	NodeFlavorId                string    `json:"nodeFlavorId,omitempty"`
	NodeFlavorName              string    `json:"nodeFlavorName,omitempty"`
	NodeInstanceTypeId          string    `json:"nodeInstanceTypeId,omitempty"`
	NodegroupDefaultId          string    `json:"nodegroupDefaultId,omitempty"`
	ProjectId                   string    `json:"projectId,omitempty"`
	SshKeyId                    string    `json:"sshKeyId,omitempty"`
	SshKeyName                  string    `json:"sshKeyName,omitempty"`
	Status                      string    `json:"status,omitempty"`
	SubnetId                    string    `json:"subnetId,omitempty"`
	Uuid                        string    `json:"uuid,omitempty"`
}
