# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMemoryStoreBackUpStorage**](MemoryStoreBackupStorageAPIApi.md#CreateMemoryStoreBackUpStorage) | **Post** /vdb-memory/v1/payment/backup-storages | 
[**DeleteBackupStorage**](MemoryStoreBackupStorageAPIApi.md#DeleteBackupStorage) | **Post** /vdb-memory/v1/backup-storages/actions/delete | 
[**GetListBackupStorage**](MemoryStoreBackupStorageAPIApi.md#GetListBackupStorage) | **Get** /vdb-memory/v1/backup-storages | 
[**GetListQuotaPackage**](MemoryStoreBackupStorageAPIApi.md#GetListQuotaPackage) | **Get** /vdb-memory/v1/backup-storages/packages | 
[**RenewBackupStorage**](MemoryStoreBackupStorageAPIApi.md#RenewBackupStorage) | **Post** /vdb-memory/v1/backup-storages/actions/renew | 
[**ResizeBackupStorage**](MemoryStoreBackupStorageAPIApi.md#ResizeBackupStorage) | **Post** /vdb-memory/v1/backup-storages/actions/resize | 

# **CreateMemoryStoreBackUpStorage**
> WrapContentListOrderResponse CreateMemoryStoreBackUpStorage(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***MemoryStoreBackupStorageAPIApiCreateMemoryStoreBackUpStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupStorageAPIApiCreateMemoryStoreBackUpStorageOpts struct
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

# **DeleteBackupStorage**
> WrapContentListActionDbInstancesResponse DeleteBackupStorage(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackupStorage**
> WrapContentListBackupStorageDetail GetListBackupStorage(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemoryStoreBackupStorageAPIApiGetListBackupStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupStorageAPIApiGetListBackupStorageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineGroup** | **optional.Int32**|  | [default to 2]

### Return type

[**WrapContentListBackupStorageDetail**](WrapContentListBackupStorageDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListQuotaPackage**
> WrapContentListDbBackupPackageResponse GetListQuotaPackage(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListDbBackupPackageResponse**](WrapContentListDbBackupPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewBackupStorage**
> WrapContentListOrderResponse RenewBackupStorage(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***MemoryStoreBackupStorageAPIApiRenewBackupStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupStorageAPIApiRenewBackupStorageOpts struct
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

# **ResizeBackupStorage**
> WrapContentListOrderResponse ResizeBackupStorage(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***MemoryStoreBackupStorageAPIApiResizeBackupStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreBackupStorageAPIApiResizeBackupStorageOpts struct
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

