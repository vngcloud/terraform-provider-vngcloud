# NodeGroupDetailDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**ClusterId** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**NumNodes** | **int64** |  | [optional] [default to null]
**ImageId** | **string** |  | [optional] [default to null]
**CreatedAt** | **string** |  | [optional] [default to null]
**UpdatedAt** | **string** |  | [optional] [default to null]
**FlavorId** | **string** |  | [optional] [default to null]
**DiskSize** | **int32** |  | [optional] [default to null]
**DiskType** | **string** |  | [optional] [default to null]
**EnablePrivateNodes** | **bool** |  | [optional] [default to null]
**SshKeyId** | **string** |  | [optional] [default to null]
**SecurityGroups** | **[]string** |  | [optional] [default to null]
**AutoScaleConfig** | [***NodeGroupAutoScaleConfigDto**](NodeGroupAutoScaleConfigDto.md) |  | [optional] [default to null]
**UpgradeConfig** | [***NodeGroupUpgradeConfigDto**](NodeGroupUpgradeConfigDto.md) |  | [optional] [default to null]
**Labels** | **map[string]string** |  | [optional] [default to null]
**Taints** | [**[]NodeGroupTaintDto**](NodeGroupTaintDto.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

