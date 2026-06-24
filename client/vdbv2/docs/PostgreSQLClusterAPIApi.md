# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrderPostgreCluster**](PostgreSQLClusterAPIApi.md#CreateOrderPostgreCluster) | **Post** /vdb-postgresql/v1/cluster | 
[**GetDatastoresPostgreCluster**](PostgreSQLClusterAPIApi.md#GetDatastoresPostgreCluster) | **Get** /vdb-postgresql/v1/cluster/datastore | 
[**GetPostgreClusterVolumeUsed**](PostgreSQLClusterAPIApi.md#GetPostgreClusterVolumeUsed) | **Get** /vdb-postgresql/v1/cluster/{clusterId}/volume-used | 
[**GetVolumeTypesPostgreCluster**](PostgreSQLClusterAPIApi.md#GetVolumeTypesPostgreCluster) | **Get** /vdb-postgresql/v1/cluster/volume-types | 
[**ListFlavorPostgreCluster**](PostgreSQLClusterAPIApi.md#ListFlavorPostgreCluster) | **Get** /vdb-postgresql/v1/cluster/flavors | 
[**ResizePostgreCluster**](PostgreSQLClusterAPIApi.md#ResizePostgreCluster) | **Put** /vdb-postgresql/v1/cluster/{clusterId}/resize | 
[**UpdateConfigGroupPostgreCluster**](PostgreSQLClusterAPIApi.md#UpdateConfigGroupPostgreCluster) | **Put** /vdb-postgresql/v1/cluster/{clusterId}/config-group | 
[**UpdateSettingsPostgreCluster**](PostgreSQLClusterAPIApi.md#UpdateSettingsPostgreCluster) | **Put** /vdb-postgresql/v1/cluster/{clusterId}/settings | 

# **CreateOrderPostgreCluster**
> WrapContentListOrderResponse CreateOrderPostgreCluster(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePostgreClusterRequest**](CreatePostgreClusterRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***PostgreSQLClusterAPIApiCreateOrderPostgreClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostgreSQLClusterAPIApiCreateOrderPostgreClusterOpts struct
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

# **GetDatastoresPostgreCluster**
> WrapContentListPostgreVersion GetDatastoresPostgreCluster(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListPostgreVersion**](WrapContentListPostgreVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPostgreClusterVolumeUsed**
> WrapContentListString GetPostgreClusterVolumeUsed(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**WrapContentListString**](WrapContentListString.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeTypesPostgreCluster**
> WrapContentListPostgreStorageType GetVolumeTypesPostgreCluster(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***PostgreSQLClusterAPIApiGetVolumeTypesPostgreClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostgreSQLClusterAPIApiGetVolumeTypesPostgreClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneId** | **optional.String**|  | [default to HCM03-1A]

### Return type

[**WrapContentListPostgreStorageType**](WrapContentListPostgreStorageType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFlavorPostgreCluster**
> WrapContentListPostgrePackage ListFlavorPostgreCluster(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***PostgreSQLClusterAPIApiListFlavorPostgreClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostgreSQLClusterAPIApiListFlavorPostgreClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneId** | **optional.String**|  | [default to HCM03-1A]

### Return type

[**WrapContentListPostgrePackage**](WrapContentListPostgrePackage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizePostgreCluster**
> WrapContentListOrderResponse ResizePostgreCluster(ctx, body, clusterId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResizePostgreClusterRequest**](ResizePostgreClusterRequest.md)|  | 
  **clusterId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***PostgreSQLClusterAPIApiResizePostgreClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostgreSQLClusterAPIApiResizePostgreClusterOpts struct
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

# **UpdateConfigGroupPostgreCluster**
> WrapContentInstanceActionResult UpdateConfigGroupPostgreCluster(ctx, body, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePostgreClusterConfigGroupRequest**](UpdatePostgreClusterConfigGroupRequest.md)|  | 
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**WrapContentInstanceActionResult**](WrapContentInstanceActionResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsPostgreCluster**
> WrapContentInstanceActionResult UpdateSettingsPostgreCluster(ctx, body, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePostgreClusterSettingsRequest**](UpdatePostgreClusterSettingsRequest.md)|  | 
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**WrapContentInstanceActionResult**](WrapContentInstanceActionResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

