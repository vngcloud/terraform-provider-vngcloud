# CreateNodeGroupDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**NumNodes** | **int32** |  | [default to null]
**AutoScaleConfig** | [***NodeGroupAutoScaleConfigDto**](NodeGroupAutoScaleConfigDto.md) |  | [optional] [default to null]
**UpgradeConfig** | [***NodeGroupUpgradeConfigDto**](NodeGroupUpgradeConfigDto.md) |  | [default to null]
**ImageId** | **string** |  | [default to null]
**FlavorId** | **string** |  | [default to null]
**DiskSize** | **int32** |  | [default to null]
**DiskType** | **string** |  | [default to null]
**EnablePrivateNodes** | **bool** |  | [default to null]
**SecurityGroups** | **[]string** |  | [default to null]
**SshKeyId** | **string** |  | [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**Taints** | [**[]NodeGroupTaintDto**](NodeGroupTaintDto.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

