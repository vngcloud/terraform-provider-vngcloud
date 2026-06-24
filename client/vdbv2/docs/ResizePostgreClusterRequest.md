# ResizePostgreClusterRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Resize Type. Allowed values:  VOLUME-SIZE, VOLUME-TYPE, NUMBER-OF-NODES | [optional] [default to null]
**NumberOfNodes** | **int32** | New Number of Nodes. Use with type NUMBER-OF-NODES | [optional] [default to null]
**VolumeTypeId** | **string** | New Volume Type Id. Use with type VOLUME-TYPE | [optional] [default to null]
**VolumeSize** | **int32** | New Volume Size. Use with type VOLUME-SIZE | [optional] [default to null]
**IsPoc** | **bool** | Set to &#x27;true&#x27; if paying with PoC Credit. Auto Payment only. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

