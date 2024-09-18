# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackups**](MemoryStoreBackupAPIApi.md#CreateBackups) | **Post** /vdb-memory/v1/backups/create | 
[**DeleteBackups**](MemoryStoreBackupAPIApi.md#DeleteBackups) | **Post** /vdb-memory/v1/backups/delete | 
[**GetDetailBackupById**](MemoryStoreBackupAPIApi.md#GetDetailBackupById) | **Get** /vdb-memory/v1/backups/{backupId}/detail | 
[**GetFreeBackupUsage**](MemoryStoreBackupAPIApi.md#GetFreeBackupUsage) | **Get** /vdb-memory/v1/backups/free-backup | 
[**GetListBackups**](MemoryStoreBackupAPIApi.md#GetListBackups) | **Get** /vdb-memory/v1/backups | 
[**RestoreBackup**](MemoryStoreBackupAPIApi.md#RestoreBackup) | **Post** /vdb-memory/v1/backups/{backupId}/restore | 

# **CreateBackups**
> WrapContentCreateBackupResponse CreateBackups(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

[**WrapContentCreateBackupResponse**](WrapContentCreateBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackups**
> WrapContentListDeleteBackupResponse DeleteBackups(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

[**WrapContentListDeleteBackupResponse**](WrapContentListDeleteBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailBackupById**
> WrapContentBackupInfo GetDetailBackupById(ctx, backupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 

### Return type

[**WrapContentBackupInfo**](WrapContentBackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreeBackupUsage**
> WrapContentFreeBackupStorageInfo GetFreeBackupUsage(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentFreeBackupStorageInfo**](WrapContentFreeBackupStorageInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackups**
> WrapContentBackupInfoGatewayResponse GetListBackups(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemoryStoreBackupAPIApiGetListBackupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupAPIApiGetListBackupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 10]

### Return type

[**WrapContentBackupInfoGatewayResponse**](WrapContentBackupInfoGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackup**
> WrapContentListOrderResponse RestoreBackup(ctx, body, backupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **backupId** | **string**|  | 
 **optional** | ***MemoryStoreBackupAPIApiRestoreBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupAPIApiRestoreBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userType** | **optional.**|  | [default to ROOT_USER]

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

