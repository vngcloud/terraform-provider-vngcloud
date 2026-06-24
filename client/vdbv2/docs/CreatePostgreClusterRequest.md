# CreatePostgreClusterRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Cluster name | [optional] [default to null]
**LocateZoneId** | **string** | Zone ID | [optional] [default to null]
**PackageId** | **string** | Flavor ID. Can only be compatible with a PostgreSQL Cluster Flavor. | [optional] [default to null]
**VolumeTypeId** | **string** | Volume Type ID. Can only be compatible with a PostgreSQL Cluster Volume Type. | [optional] [default to null]
**VolumeSize** | **int32** | Volume Size | [optional] [default to null]
**NumberOfNodes** | **int32** | Number of nodes. Minimum 2 nodes and maximum 10 nodes. | [optional] [default to null]
**User** | [***UserRequest2**](UserRequest_2.md) |  | [optional] [default to null]
**Databases** | [**[]DatabaseRequest2**](DatabaseRequest_2.md) | Database name. Currently support 1 database name at creation. | [optional] [default to null]
**DatastoreVersion** | **string** | Datastore version. Can only be compatible with a PostgreSQL Cluster Datastore Version. | [optional] [default to null]
**NetIds** | **[]string** | Subnet ID | [optional] [default to null]
**ConfigId** | **string** | Config Group ID. Can only be compatible with a config group with a &#x27;cluster&#x27; deploy type. | [optional] [default to null]
**PublicAccess** | **bool** | Public access | [optional] [default to null]
**BackupLocationId** | **string** | Backup Location ID. Managed by vBackup. | [optional] [default to null]
**BackupPolicyId** | **string** | Backup Policy ID. Managed by vBackup. | [optional] [default to null]
**BackupPointId** | **string** | Backup Point ID. Managed by vBackup. Required if Restoring cluster from Backup. | [optional] [default to null]
**IsPoc** | **bool** | Set to &#x27;true&#x27; if paying with PoC Credit. Auto Payment only. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

