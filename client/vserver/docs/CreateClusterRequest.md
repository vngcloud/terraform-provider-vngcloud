# CreateClusterRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoHealingEnabled** | **bool** | Auto-healing feature is enable | [optional] [default to null]
**AutoMonitoringEnabled** | **bool** | Auto monitoring feature is enable | [optional] [default to null]
**AutoScalingEnabled** | **bool** | Auto scaling feature is enable | [optional] [default to null]
**BootVolumeSize** | **int32** | Size of boot volume | [default to null]
**BootVolumeTypeId** | **string** | Id of boot volume type | [default to null]
**CalicoCidr** | **string** | Cidr of Calico  | [default to null]
**Description** | **string** | Description of the Cluster | [optional] [default to null]
**DockerVolumeSize** | **int32** | Size of docker volume | [default to null]
**DockerVolumeTypeId** | **string** | Id of docker volume type | [default to null]
**EnabledLb** | **bool** | Enable Load Balancer or Not | [default to null]
**EtcdVolumeSize** | **int32** | Size of etcd volume | [default to null]
**EtcdVolumeTypeId** | **string** | Id of etcd volume type | [default to null]
**IngressControllerEnabled** | **bool** | Ingress Controller feature is enable | [optional] [default to null]
**IpipMode** | **string** | IpIp Mode | [default to null]
**K8sVersion** | **string** | Id of K8s Version of the Cluster | [default to null]
**MasterCount** | **int32** | Number of Master Node | [default to null]
**MasterFlavorId** | **string** | Id of the flavor of master | [default to null]
**MasterInstanceTypeId** | **string** | Id of the image of master | [default to null]
**MaxNodeCount** | **int32** | Maximum number of node count | [default to null]
**MinNodeCount** | **int32** | Minimum number of node Count | [default to null]
**Name** | **string** | Name of the Cluster | [default to null]
**NetworkId** | **string** | Id of the network | [default to null]
**NetworkType** | **string** | Type of the network | [default to null]
**NodeCount** | **int32** | Number of Minion node | [default to null]
**NodeFlavorId** | **string** | Id of the flavor of minion | [default to null]
**NodeGroupRequestModelList** | [**[]NodeGroupRequestModel**](NodeGroupRequestModel.md) | List of Node Group want to create | [optional] [default to null]
**NodeInstanceTypeId** | **string** | Id of the image of minion | [default to null]
**SecGroupIds** | **[]string** | List of Id of the secGroup | [optional] [default to null]
**SshKeyId** | **string** | Id of the sshKey | [optional] [default to null]
**SubnetId** | **string** | Id of the subnet | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


