# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupNow**](PostgreSQLBackupAPIApi.md#BackupNow) | **Post** /vdb-postgresql/v1/backup/backup-vdb/{clusterId}/backup-now | 
[**GetBackupVDB**](PostgreSQLBackupAPIApi.md#GetBackupVDB) | **Get** /vdb-postgresql/v1/backup/backup-vdb/{clusterId}/detail | 
[**ListBackupLocation**](PostgreSQLBackupAPIApi.md#ListBackupLocation) | **Get** /vdb-postgresql/v1/backup/location | 
[**ListBackupPolicy**](PostgreSQLBackupAPIApi.md#ListBackupPolicy) | **Get** /vdb-postgresql/v1/backup/policy | 
[**ListBackupRestorePoint**](PostgreSQLBackupAPIApi.md#ListBackupRestorePoint) | **Get** /vdb-postgresql/v1/backup/backup-vdb/{clusterId}/restore-point | 
[**ListBackupVDB**](PostgreSQLBackupAPIApi.md#ListBackupVDB) | **Get** /vdb-postgresql/v1/backup/backup-vdb | 

# **BackupNow**
> WrapContentString BackupNow(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**WrapContentString**](WrapContentString.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupVDB**
> WrapContentBackupDatabase GetBackupVDB(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**WrapContentBackupDatabase**](WrapContentBackupDatabase.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupLocation**
> WrapContentListBackupLocation ListBackupLocation(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListBackupLocation**](WrapContentListBackupLocation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupPolicy**
> WrapContentListBackupPolicy ListBackupPolicy(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListBackupPolicy**](WrapContentListBackupPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupRestorePoint**
> WrapContentListBackupPoint ListBackupRestorePoint(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**WrapContentListBackupPoint**](WrapContentListBackupPoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupVDB**
> WrapContentListBackupDatabase ListBackupVDB(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListBackupDatabase**](WrapContentListBackupDatabase.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

