# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelationalBackUpStorage**](RelationalBackupStorageAPIApi.md#CreateRelationalBackUpStorage) | **Post** /v1/payment/backup-storages | 
[**DeleteBackupStorage1**](RelationalBackupStorageAPIApi.md#DeleteBackupStorage1) | **Post** /v1/backup-storages/actions/deletions | 
[**GetListBackupStorage1**](RelationalBackupStorageAPIApi.md#GetListBackupStorage1) | **Get** /v1/backup-storages/information | 
[**GetListQuotaPackage1**](RelationalBackupStorageAPIApi.md#GetListQuotaPackage1) | **Get** /v1/backup-storages | 
[**RenewBackupStorage1**](RelationalBackupStorageAPIApi.md#RenewBackupStorage1) | **Post** /v1/backup-storages/actions/renew | 
[**ResizeBackupStorage1**](RelationalBackupStorageAPIApi.md#ResizeBackupStorage1) | **Post** /v1/backup-storages/actions/resize | 

# **CreateRelationalBackUpStorage**
> WrapContentListOrderResponse CreateRelationalBackUpStorage(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***RelationalBackupStorageAPIApiCreateRelationalBackUpStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupStorageAPIApiCreateRelationalBackUpStorageOpts struct
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

# **DeleteBackupStorage1**
> WrapContentListActionDbInstancesResponse DeleteBackupStorage1(ctx, body)


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

# **GetListBackupStorage1**
> WrapContentListBackupStorageDetail GetListBackupStorage1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RelationalBackupStorageAPIApiGetListBackupStorage1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupStorageAPIApiGetListBackupStorage1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineGroup** | **optional.Int32**|  | [default to 1]

### Return type

[**WrapContentListBackupStorageDetail**](WrapContentListBackupStorageDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListQuotaPackage1**
> WrapContentListDbBackupPackageResponse GetListQuotaPackage1(ctx, )


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

# **RenewBackupStorage1**
> WrapContentListOrderResponse RenewBackupStorage1(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***RelationalBackupStorageAPIApiRenewBackupStorage1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupStorageAPIApiRenewBackupStorage1Opts struct
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

# **ResizeBackupStorage1**
> WrapContentListOrderResponse ResizeBackupStorage1(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***RelationalBackupStorageAPIApiResizeBackupStorage1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupStorageAPIApiResizeBackupStorage1Opts struct
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

