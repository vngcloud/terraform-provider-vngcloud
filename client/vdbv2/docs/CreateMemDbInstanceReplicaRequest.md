# CreateMemDbInstanceReplicaRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | New DB instance name | [optional] [default to null]
**DatastoreType** | **string** | Datastore type | [optional] [default to null]
**DatastoreVersion** | **string** | Version of datastore type | [optional] [default to null]
**NetIds** | **[]string** | Subnet ID | [optional] [default to null]
**ConfigId** | **string** | Optional. Config group ID attached to DB instance | [optional] [default to null]
**PublicAccess** | **bool** | Set to &#x27;true&#x27; if you want to allow public access to the DB instance | [optional] [default to null]
**BackupAuto** | **bool** | Set to &#x27;true&#x27; if you want to enable daily auto backup | [optional] [default to null]
**BackupDuration** | **int32** | Required if &#x27;backupAuto&#x27; is true. Determine how many days your backup will be retained. Minimum 2 days and maximum 14 days | [optional] [default to null]
**BackupTime** | **string** | Required if &#x27;backupAuto&#x27; is true. The time of the day to backup your DB instance. | [optional] [default to null]
**PackageId** | **string** | Package ID | [optional] [default to null]
**ReplicaSourceId** | **string** | ID of the source DB instance used to create this replica (slave). Required when creating a read replica. | [optional] [default to null]
**Poc** | **bool** | Set to &#x27;true&#x27; if paying with PoC Credit. Auto Payment only. | [optional] [default to null]
**LocateZoneId** | **string** | Zone ID | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

