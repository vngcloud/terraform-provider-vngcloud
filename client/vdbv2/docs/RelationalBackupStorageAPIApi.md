# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelationalBackUpStorage**](RelationalBackupStorageAPIApi.md#CreateRelationalBackUpStorage) | **Post** /vdb-relational/v1/payment/backup-storages | 
[**DeleteBackupStorage**](RelationalBackupStorageAPIApi.md#DeleteBackupStorage) | **Post** /vdb-relational/v1/backup-storages/actions/deletions | 
[**GetListBackupStorage**](RelationalBackupStorageAPIApi.md#GetListBackupStorage) | **Get** /vdb-relational/v1/backup-storages/information | 
[**GetListQuotaPackage**](RelationalBackupStorageAPIApi.md#GetListQuotaPackage) | **Get** /vdb-relational/v1/backup-storages | 
[**ResizeBackupStorage**](RelationalBackupStorageAPIApi.md#ResizeBackupStorage) | **Post** /vdb-relational/v1/backup-storages/actions/resize | 

# **CreateRelationalBackUpStorage**
> WrapContentListOrderResponse CreateRelationalBackUpStorage(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBackupStorageRequest**](CreateBackupStorageRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalBackupStorageAPIApiCreateRelationalBackUpStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupStorageAPIApiCreateRelationalBackUpStorageOpts struct
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

# **DeleteBackupStorage**
> WrapContentListActionDbInstancesResponse DeleteBackupStorage(ctx, body, portalUserId)


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

# **GetListBackupStorage**
> WrapContentListBackupStorageDetail GetListBackupStorage(ctx, portalUserId)


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

# **GetListQuotaPackage**
> WrapContentListDbBackupPackageResponse GetListQuotaPackage(ctx, portalUserId)


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

# **ResizeBackupStorage**
> WrapContentListOrderResponse ResizeBackupStorage(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResizeBackupStorageRequest**](ResizeBackupStorageRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalBackupStorageAPIApiResizeBackupStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalBackupStorageAPIApiResizeBackupStorageOpts struct
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

