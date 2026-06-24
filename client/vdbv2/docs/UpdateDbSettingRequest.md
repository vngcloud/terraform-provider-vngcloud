# UpdateDbSettingRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbInstanceId** | **string** | DB instance ID | [optional] [default to null]
**Password** | **string** | Optional. New master password. | [optional] [default to null]
**PublicAccess** | **bool** | Optional. Set to &#x27;true&#x27; if you want to allow public access to the DB instance. | [optional] [default to null]
**BackupAuto** | **bool** | Optional. Set to &#x27;true&#x27; if you want to enable daily auto backup. | [optional] [default to null]
**BackupDuration** | **int32** | Required if &#x27;backupAuto&#x27; is true. Determine how many days your backup will be retained. Minimum 2 days and maximum 14 days. | [optional] [default to null]
**BackupTime** | **string** | Required if &#x27;backupAuto&#x27; is true. The time of the day to backup your DB instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

