# \VdbBackupEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupUsingPOST**](VdbBackupEndPointApi.md#CreateBackupUsingPOST) | **Post** /v1/trv/{projectId}/backups | createBackup
[**DeleteBackupUsingDELETE**](VdbBackupEndPointApi.md#DeleteBackupUsingDELETE) | **Delete** /v1/trv/{projectId}/backups/{backupId} | deleteBackup
[**GetBackupByIdUsingGET**](VdbBackupEndPointApi.md#GetBackupByIdUsingGET) | **Get** /v1/trv/{projectId}/backups/{backupId} | getBackupById
[**GetListBackupsByInstanceUsingGET**](VdbBackupEndPointApi.md#GetListBackupsByInstanceUsingGET) | **Get** /v1/trv/{projectId}/backups/instance/{dbInstanceId} | getListBackupsByInstance
[**GetListBackupsUsingGET**](VdbBackupEndPointApi.md#GetListBackupsUsingGET) | **Get** /v1/trv/{projectId}/backups | getListBackups
[**RestoreBackupUsingPOST**](VdbBackupEndPointApi.md#RestoreBackupUsingPOST) | **Post** /v1/trv/{projectId}/backups/restore | restoreBackup


# **CreateBackupUsingPOST**
> BackupResponse CreateBackupUsingPOST(ctx, projectId, request)
createBackup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**BackupRequest**](BackupRequest.md)| request | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupUsingDELETE**
> BackupResponse DeleteBackupUsingDELETE(ctx, backupId, projectId)
deleteBackup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**| backupId | 
  **projectId** | **string**| projectId | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupByIdUsingGET**
> BackupResponse GetBackupByIdUsingGET(ctx, backupId, projectId)
getBackupById

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**| backupId | 
  **projectId** | **string**| projectId | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackupsByInstanceUsingGET**
> ListBackupResponse GetListBackupsByInstanceUsingGET(ctx, dbInstanceId, projectId)
getListBackupsByInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**ListBackupResponse**](ListBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackupsUsingGET**
> ListBackupResponse GetListBackupsUsingGET(ctx, projectId)
getListBackups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**ListBackupResponse**](ListBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackupUsingPOST**
> DbInstanceResponse RestoreBackupUsingPOST(ctx, projectId, request)
restoreBackup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**RestoreRequest**](RestoreRequest.md)| request | 

### Return type

[**DbInstanceResponse**](DbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

