# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackups1**](RelationalBackupAPIApi.md#CreateBackups1) | **Post** /v1/backups/create | 
[**DeleteBackups1**](RelationalBackupAPIApi.md#DeleteBackups1) | **Delete** /v1/backups/{backupId}/delete | 
[**GetDetailBackupById1**](RelationalBackupAPIApi.md#GetDetailBackupById1) | **Get** /v1/backups/detail/{backupId} | 
[**GetFreeBackupUsage1**](RelationalBackupAPIApi.md#GetFreeBackupUsage1) | **Get** /v1/backups/free-backup | 
[**GetListBackups1**](RelationalBackupAPIApi.md#GetListBackups1) | **Get** /v1/backups | 
[**GetListBackupsByInstanceId1**](RelationalBackupAPIApi.md#GetListBackupsByInstanceId1) | **Get** /v1/backups/insId/{instanceId} | 
[**RestoreBackup1**](RelationalBackupAPIApi.md#RestoreBackup1) | **Post** /v1/backups/{id}/restore | 

# **CreateBackups1**
> WrapContentCreateBackupResponse CreateBackups1(ctx, body)


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

# **DeleteBackups1**
> WrapContentListDeleteBackupResponse DeleteBackups1(ctx, body, backupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **backupId** | **string**|  | 

### Return type

[**WrapContentListDeleteBackupResponse**](WrapContentListDeleteBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailBackupById1**
> WrapContentBackupInfo GetDetailBackupById1(ctx, backupId)


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

# **GetFreeBackupUsage1**
> WrapContentFreeBackupStorageInfo GetFreeBackupUsage1(ctx, )


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

# **GetListBackups1**
> WrapContentBackupInfoGatewayResponse GetListBackups1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RelationalBackupAPIApiGetListBackups1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupAPIApiGetListBackups1Opts struct
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

# **GetListBackupsByInstanceId1**
> WrapContentListBackupInfo GetListBackupsByInstanceId1(ctx, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListBackupInfo**](WrapContentListBackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackup1**
> WrapContentListOrderResponse RestoreBackup1(ctx, body, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **id** | **string**|  | 
 **optional** | ***RelationalBackupAPIApiRestoreBackup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupAPIApiRestoreBackup1Opts struct
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

