# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersClusterIdDelete**](V1ClusterControllerApi.md#V1ClustersClusterIdDelete) | **Delete** /v1/clusters/{clusterId} | 
[**V1ClustersClusterIdGet**](V1ClusterControllerApi.md#V1ClustersClusterIdGet) | **Get** /v1/clusters/{clusterId} | 
[**V1ClustersClusterIdKubeconfigGet**](V1ClusterControllerApi.md#V1ClustersClusterIdKubeconfigGet) | **Get** /v1/clusters/{clusterId}/kubeconfig | 
[**V1ClustersClusterIdPut**](V1ClusterControllerApi.md#V1ClustersClusterIdPut) | **Put** /v1/clusters/{clusterId} | 
[**V1ClustersGet**](V1ClusterControllerApi.md#V1ClustersGet) | **Get** /v1/clusters | 
[**V1ClustersPost**](V1ClusterControllerApi.md#V1ClustersPost) | **Post** /v1/clusters | 

# **V1ClustersClusterIdDelete**
> ClusterDto V1ClustersClusterIdDelete(ctx, clusterId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***V1ClusterControllerApiV1ClustersClusterIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ClusterControllerApiV1ClustersClusterIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalUserId** | **optional.Int64**|  | 

### Return type

[**ClusterDto**](ClusterDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdGet**
> ClusterDetailDto V1ClustersClusterIdGet(ctx, clusterId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***V1ClusterControllerApiV1ClustersClusterIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ClusterControllerApiV1ClustersClusterIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalUserId** | **optional.Int64**|  | 

### Return type

[**ClusterDetailDto**](ClusterDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdKubeconfigGet**
> string V1ClustersClusterIdKubeconfigGet(ctx, clusterId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***V1ClusterControllerApiV1ClustersClusterIdKubeconfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ClusterControllerApiV1ClustersClusterIdKubeconfigGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalUserId** | **optional.Int64**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdPut**
> ClusterDto V1ClustersClusterIdPut(ctx, clusterId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***V1ClusterControllerApiV1ClustersClusterIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ClusterControllerApiV1ClustersClusterIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateClusterDto**](UpdateClusterDto.md)|  | 
 **portalUserId** | **optional.**|  | 

### Return type

[**ClusterDto**](ClusterDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersGet**
> PagingResultDtoClusterDto V1ClustersGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1ClusterControllerApiV1ClustersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ClusterControllerApiV1ClustersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**|  | [default to {}]
 **page** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 10]
 **portalUserId** | **optional.Int64**|  | 

### Return type

[**PagingResultDtoClusterDto**](PagingResultDtoClusterDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersPost**
> ClusterDto V1ClustersPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1ClusterControllerApiV1ClustersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ClusterControllerApiV1ClustersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateClusterComboDto**](CreateClusterComboDto.md)|  | 
 **poc** | **optional.**|  | [default to false]
 **portalUserId** | **optional.**|  | 

### Return type

[**ClusterDto**](ClusterDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

