# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersClusterIdNodeGroupsGet**](V1NodeGroupControllerApi.md#V1ClustersClusterIdNodeGroupsGet) | **Get** /v1/clusters/{clusterId}/node-groups | 
[**V1ClustersClusterIdNodeGroupsNodeGroupIdDelete**](V1NodeGroupControllerApi.md#V1ClustersClusterIdNodeGroupsNodeGroupIdDelete) | **Delete** /v1/clusters/{clusterId}/node-groups/{nodeGroupId} | 
[**V1ClustersClusterIdNodeGroupsNodeGroupIdGet**](V1NodeGroupControllerApi.md#V1ClustersClusterIdNodeGroupsNodeGroupIdGet) | **Get** /v1/clusters/{clusterId}/node-groups/{nodeGroupId} | 
[**V1ClustersClusterIdNodeGroupsNodeGroupIdNodesGet**](V1NodeGroupControllerApi.md#V1ClustersClusterIdNodeGroupsNodeGroupIdNodesGet) | **Get** /v1/clusters/{clusterId}/node-groups/{nodeGroupId}/nodes | 
[**V1ClustersClusterIdNodeGroupsNodeGroupIdPut**](V1NodeGroupControllerApi.md#V1ClustersClusterIdNodeGroupsNodeGroupIdPut) | **Put** /v1/clusters/{clusterId}/node-groups/{nodeGroupId} | 
[**V1ClustersClusterIdNodeGroupsPost**](V1NodeGroupControllerApi.md#V1ClustersClusterIdNodeGroupsPost) | **Post** /v1/clusters/{clusterId}/node-groups | 

# **V1ClustersClusterIdNodeGroupsGet**
> PagingResultDtoNodeGroupDto V1ClustersClusterIdNodeGroupsGet(ctx, clusterId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 10]
 **portalUserId** | **optional.Int64**|  | 

### Return type

[**PagingResultDtoNodeGroupDto**](PagingResultDtoNodeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdNodeGroupsNodeGroupIdDelete**
> NodeGroupDto V1ClustersClusterIdNodeGroupsNodeGroupIdDelete(ctx, clusterId, nodeGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **nodeGroupId** | **string**|  | 
 **optional** | ***V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalUserId** | **optional.Int64**|  | 

### Return type

[**NodeGroupDto**](NodeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdNodeGroupsNodeGroupIdGet**
> NodeGroupDetailDto V1ClustersClusterIdNodeGroupsNodeGroupIdGet(ctx, clusterId, nodeGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **nodeGroupId** | **string**|  | 
 **optional** | ***V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portalUserId** | **optional.Int64**|  | 

### Return type

[**NodeGroupDetailDto**](NodeGroupDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdNodeGroupsNodeGroupIdNodesGet**
> PagingResultDtoNodeDto V1ClustersClusterIdNodeGroupsNodeGroupIdNodesGet(ctx, clusterId, nodeGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **nodeGroupId** | **string**|  | 
 **optional** | ***V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdNodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdNodesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 10]
 **portalUserId** | **optional.Int64**|  | 

### Return type

[**PagingResultDtoNodeDto**](PagingResultDtoNodeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdNodeGroupsNodeGroupIdPut**
> NodeGroupDto V1ClustersClusterIdNodeGroupsNodeGroupIdPut(ctx, clusterId, nodeGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **nodeGroupId** | **string**|  | 
 **optional** | ***V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateNodeGroupDto**](UpdateNodeGroupDto.md)|  | 
 **portalUserId** | **optional.**|  | 

### Return type

[**NodeGroupDto**](NodeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClustersClusterIdNodeGroupsPost**
> NodeGroupDto V1ClustersClusterIdNodeGroupsPost(ctx, clusterId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateNodeGroupDto**](CreateNodeGroupDto.md)|  | 
 **portalUserId** | **optional.**|  | 

### Return type

[**NodeGroupDto**](NodeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

