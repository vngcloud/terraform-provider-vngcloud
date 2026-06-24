# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackups**](RelationalBackupAPIApi.md#CreateBackups) | **Post** /vdb-relational/v1/backups/create | 
[**DeleteBackups1**](RelationalBackupAPIApi.md#DeleteBackups1) | **Delete** /vdb-relational/v1/backups/{backupId}/delete | 
[**GetDetailBackupById**](RelationalBackupAPIApi.md#GetDetailBackupById) | **Get** /vdb-relational/v1/backups/detail/{backupId} | 
[**GetFreeBackupUsage**](RelationalBackupAPIApi.md#GetFreeBackupUsage) | **Get** /vdb-relational/v1/backups/free-backup | 
[**GetListBackups**](RelationalBackupAPIApi.md#GetListBackups) | **Get** /vdb-relational/v1/backups | 
[**GetListBackupsByInstanceId**](RelationalBackupAPIApi.md#GetListBackupsByInstanceId) | **Get** /vdb-relational/v1/backups/insId/{instanceId} | 
[**RestoreBackup**](RelationalBackupAPIApi.md#RestoreBackup) | **Post** /vdb-relational/v1/backups/{id}/restore | 

# **CreateBackups**
> WrapContentCreateBackupResponse CreateBackups(ctx, body, portalUserId)


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

# **DeleteBackups1**
> WrapContentListDeleteBackupResponse DeleteBackups1(ctx, body, backupId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DeleteBackupRequest**](DeleteBackupRequest.md)|  | 
  **backupId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListDeleteBackupResponse**](WrapContentListDeleteBackupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailBackupById**
> WrapContentBackupInfo GetDetailBackupById(ctx, portalUserId, backupId)


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

# **GetFreeBackupUsage**
> WrapContentFreeBackupStorageInfo GetFreeBackupUsage(ctx, portalUserId)


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

# **GetListBackups**
> WrapContentBackupInfoGatewayResponse GetListBackups(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalBackupAPIApiGetListBackupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupAPIApiGetListBackupsOpts struct
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

# **GetListBackupsByInstanceId**
> WrapContentListBackupInfo GetListBackupsByInstanceId(ctx, portalUserId, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListBackupInfo**](WrapContentListBackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackup**
> WrapContentListOrderResponse RestoreBackup(ctx, body, id, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RestoreBackupRequest**](RestoreBackupRequest.md)|  | 
  **id** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalBackupAPIApiRestoreBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupAPIApiRestoreBackupOpts struct
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

