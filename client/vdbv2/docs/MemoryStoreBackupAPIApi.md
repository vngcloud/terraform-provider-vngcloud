# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackups1**](MemoryStoreBackupAPIApi.md#CreateBackups1) | **Post** /vdb-memory/v1/backups/create | 
[**DeleteBackups**](MemoryStoreBackupAPIApi.md#DeleteBackups) | **Post** /vdb-memory/v1/backups/delete | 
[**GetDetailBackupById1**](MemoryStoreBackupAPIApi.md#GetDetailBackupById1) | **Get** /vdb-memory/v1/backups/{backupId}/detail | 
[**GetFreeBackupUsage1**](MemoryStoreBackupAPIApi.md#GetFreeBackupUsage1) | **Get** /vdb-memory/v1/backups/free-backup | 
[**GetListBackups1**](MemoryStoreBackupAPIApi.md#GetListBackups1) | **Get** /vdb-memory/v1/backups | 
[**RestoreBackup1**](MemoryStoreBackupAPIApi.md#RestoreBackup1) | **Post** /vdb-memory/v1/backups/{backupId}/restore | 

# **CreateBackups1**
> WrapContentCreateBackupResponse CreateBackups1(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBackupRequest**](CreateBackupRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentCreateBackupResponse**](WrapContentCreateBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackups**
> WrapContentListDeleteBackupResponse DeleteBackups(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DeleteBackupRequest**](DeleteBackupRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListDeleteBackupResponse**](WrapContentListDeleteBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailBackupById1**
> WrapContentBackupInfo GetDetailBackupById1(ctx, portalUserId, backupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **backupId** | **string**|  | 

### Return type

[**WrapContentBackupInfo**](WrapContentBackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreeBackupUsage1**
> WrapContentFreeBackupStorageInfo GetFreeBackupUsage1(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentFreeBackupStorageInfo**](WrapContentFreeBackupStorageInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackups1**
> WrapContentBackupInfoGatewayResponse GetListBackups1(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreBackupAPIApiGetListBackups1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupAPIApiGetListBackups1Opts struct
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

# **RestoreBackup1**
> WrapContentListOrderResponse RestoreBackup1(ctx, body, backupId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestoreMemBackupRequest**](RestoreMemBackupRequest.md)|  | 
  **backupId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreBackupAPIApiRestoreBackup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupAPIApiRestoreBackup1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userType** | **optional.**| ROOT_USER for Checkout flow or IAM_USER for Auto Payment flow. Available values: ROOT_USER, IAM_USER. Default value: ROOT_USER.  | 

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

