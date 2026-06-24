# CreateBackupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbInstanceId** | **string** | Db instance Id | [optional] [default to null]
**Name** | **string** | Backup name | [optional] [default to null]
**ParentId** | **string** | Parent backup Id. Required if &#x27;backupType&#x27; is INCREMENTAL | [optional] [default to null]
**BackupType** | **string** | Backup type. Allowed values: FULL, INCREMENTAL | [optional] [default to null]
**Description** | **string** | Backup description | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

