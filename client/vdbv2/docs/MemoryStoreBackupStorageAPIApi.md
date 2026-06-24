# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMemoryStoreBackUpStorage**](MemoryStoreBackupStorageAPIApi.md#CreateMemoryStoreBackUpStorage) | **Post** /vdb-memory/v1/payment/backup-storages | 
[**DeleteBackupStorage1**](MemoryStoreBackupStorageAPIApi.md#DeleteBackupStorage1) | **Post** /vdb-memory/v1/backup-storages/actions/delete | 
[**GetListBackupStorage1**](MemoryStoreBackupStorageAPIApi.md#GetListBackupStorage1) | **Get** /vdb-memory/v1/backup-storages | 
[**GetListQuotaPackage1**](MemoryStoreBackupStorageAPIApi.md#GetListQuotaPackage1) | **Get** /vdb-memory/v1/backup-storages/packages | 
[**ResizeBackupStorage1**](MemoryStoreBackupStorageAPIApi.md#ResizeBackupStorage1) | **Post** /vdb-memory/v1/backup-storages/actions/resize | 

# **CreateMemoryStoreBackUpStorage**
> WrapContentListOrderResponse CreateMemoryStoreBackUpStorage(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBackupStorageRequest**](CreateBackupStorageRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreBackupStorageAPIApiCreateMemoryStoreBackUpStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupStorageAPIApiCreateMemoryStoreBackUpStorageOpts struct
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

# **DeleteBackupStorage1**
> WrapContentListActionDbInstancesResponse DeleteBackupStorage1(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteBackupStorageRequest**](DeleteBackupStorageRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackupStorage1**
> WrapContentListBackupStorageDetail GetListBackupStorage1(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListBackupStorageDetail**](WrapContentListBackupStorageDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListQuotaPackage1**
> WrapContentListDbBackupPackageResponse GetListQuotaPackage1(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListDbBackupPackageResponse**](WrapContentListDbBackupPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeBackupStorage1**
> WrapContentListOrderResponse ResizeBackupStorage1(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResizeBackupStorageRequest**](ResizeBackupStorageRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreBackupStorageAPIApiResizeBackupStorage1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupStorageAPIApiResizeBackupStorage1Opts struct
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

