# \K8SClusterRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachToLbUsingPOST**](K8SClusterRestControllerV2Api.md#AttachToLbUsingPOST) | **Post** /v2/{projectId}/clusters/{clusterId}/attach-load-balancer | Attach Load Balancer
[**CreateClusterNodeGroupUsingPOST**](K8SClusterRestControllerV2Api.md#CreateClusterNodeGroupUsingPOST) | **Post** /v2/{projectId}/clusters/{clusterId}/cluster-node-groups | Create Cluster Node Group
[**CreateClusterUsingPOST**](K8SClusterRestControllerV2Api.md#CreateClusterUsingPOST) | **Post** /v2/{projectId}/clusters | Create cluster
[**DeleteClusterNodeGroupUsingDELETE**](K8SClusterRestControllerV2Api.md#DeleteClusterNodeGroupUsingDELETE) | **Delete** /v2/{projectId}/clusters/{clusterId}/cluster-node-groups | Delete Cluster Node Group
[**DeleteClusterUsingDELETE**](K8SClusterRestControllerV2Api.md#DeleteClusterUsingDELETE) | **Delete** /v2/{projectId}/clusters/{clusterId} | Delete Cluster
[**GetClusterNodeGroupUsingGET**](K8SClusterRestControllerV2Api.md#GetClusterNodeGroupUsingGET) | **Get** /v2/{projectId}/clusters/cluster-node-groups/{nodegroupId} | Get Cluster Node Group
[**GetClusterUsingGET**](K8SClusterRestControllerV2Api.md#GetClusterUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId} | Get cluster
[**GetConfigUsingGET**](K8SClusterRestControllerV2Api.md#GetConfigUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId}/config | Get Config
[**GetConsoleUsingGET**](K8SClusterRestControllerV2Api.md#GetConsoleUsingGET) | **Get** /v2/{projectId}/clusters/nodes/{serverId}/console | Get Console
[**GetMaxNodeCountUsingGET**](K8SClusterRestControllerV2Api.md#GetMaxNodeCountUsingGET) | **Get** /v2/{projectId}/clusters/max-node-count | Get max node count
[**ListClusterNodeGroupUsingGET**](K8SClusterRestControllerV2Api.md#ListClusterNodeGroupUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId}/cluster-node-groups | List Cluster Node Group
[**ListClusterPoolByClusterUsingGET**](K8SClusterRestControllerV2Api.md#ListClusterPoolByClusterUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId}/pools | List cluster pools
[**ListClusterUsingGET**](K8SClusterRestControllerV2Api.md#ListClusterUsingGET) | **Get** /v2/{projectId}/clusters | List cluster
[**ListK8sNetworkTypeUsingGET**](K8SClusterRestControllerV2Api.md#ListK8sNetworkTypeUsingGET) | **Get** /v2/{projectId}/clusters/network-types | List K8s Networks Type
[**ListK8sVersionUsingGET**](K8SClusterRestControllerV2Api.md#ListK8sVersionUsingGET) | **Get** /v2/{projectId}/clusters/versions | List K8s Version
[**ListNodeUsingGET**](K8SClusterRestControllerV2Api.md#ListNodeUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId}/nodes | List node
[**ListSecGroupUsingGET**](K8SClusterRestControllerV2Api.md#ListSecGroupUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId}/sec-group/{type} | Get Cluster Sec Group Detail
[**ListSecgroupDefaultUsingGET**](K8SClusterRestControllerV2Api.md#ListSecgroupDefaultUsingGET) | **Get** /v2/{projectId}/clusters/{clusterId}/secgroups-default | List Secgroup Default
[**ListVolumeUsingGET3**](K8SClusterRestControllerV2Api.md#ListVolumeUsingGET3) | **Get** /v2/{projectId}/clusters/{clusterId}/volumes | List volume
[**ScaleMinionUsingPOST**](K8SClusterRestControllerV2Api.md#ScaleMinionUsingPOST) | **Post** /v2/{projectId}/clusters/{clusterId}/scale-minion | Scale Minion
[**UpdateSecGroupUsingPUT**](K8SClusterRestControllerV2Api.md#UpdateSecGroupUsingPUT) | **Put** /v2/{projectId}/clusters/sec-group | Update sec group


# **AttachToLbUsingPOST**
> DataResponseClusterDto AttachToLbUsingPOST(ctx, attachToLoadBalancerBackendRequest, clusterId, projectId)
Attach Load Balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachToLoadBalancerBackendRequest** | [**AttachToLoadBalancerBackendRequest**](AttachToLoadBalancerBackendRequest.md)| attachToLoadBalancerBackendRequest | 
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseClusterDto**](DataResponse«ClusterDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterNodeGroupUsingPOST**
> DataResponseClusterNodeGroupDto CreateClusterNodeGroupUsingPOST(ctx, clusterId, createNodeGroupBackendRequest, projectId)
Create Cluster Node Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **createNodeGroupBackendRequest** | [**CreateNodeGroupBackendRequest**](CreateNodeGroupBackendRequest.md)| createNodeGroupBackendRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseClusterNodeGroupDto**](DataResponse«ClusterNodeGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterUsingPOST**
> DataResponseClusterDto CreateClusterUsingPOST(ctx, createClusterRequest, projectId)
Create cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md)| createClusterRequest | 
  **projectId** | **string**| The project Id | 

### Return type

[**DataResponseClusterDto**](DataResponse«ClusterDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterNodeGroupUsingDELETE**
> DataResponseClusterNodeGroupDto DeleteClusterNodeGroupUsingDELETE(ctx, clusterId, deleteNodeGroupBackendRequest, projectId)
Delete Cluster Node Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **deleteNodeGroupBackendRequest** | [**DeleteNodeGroupBackendRequest**](DeleteNodeGroupBackendRequest.md)| deleteNodeGroupBackendRequest | 
  **projectId** | **string**| THe project id | 

### Return type

[**DataResponseClusterNodeGroupDto**](DataResponse«ClusterNodeGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterUsingDELETE**
> DataResponseClusterDto DeleteClusterUsingDELETE(ctx, clusterId, projectId)
Delete Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseClusterDto**](DataResponse«ClusterDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNodeGroupUsingGET**
> ClusterNodeGroupDto GetClusterNodeGroupUsingGET(ctx, nodegroupId, projectId)
Get Cluster Node Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodegroupId** | **string**| The node group id | 
  **projectId** | **string**| The project id | 

### Return type

[**ClusterNodeGroupDto**](ClusterNodeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterUsingGET**
> DataResponseClusterDto GetClusterUsingGET(ctx, clusterId, projectId)
Get cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project Id | 

### Return type

[**DataResponseClusterDto**](DataResponse«ClusterDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigUsingGET**
> ClusterConfigResultDto GetConfigUsingGET(ctx, clusterId, projectId)
Get Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**ClusterConfigResultDto**](ClusterConfigResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsoleUsingGET**
> NodeConsoleResultDto GetConsoleUsingGET(ctx, projectId, serverId)
Get Console

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**NodeConsoleResultDto**](NodeConsoleResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMaxNodeCountUsingGET**
> int32 GetMaxNodeCountUsingGET(ctx, projectId)
Get max node count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project Id | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterNodeGroupUsingGET**
> []ClusterNodeGroupDto ListClusterNodeGroupUsingGET(ctx, clusterId, projectId)
List Cluster Node Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**[]ClusterNodeGroupDto**](ClusterNodeGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterPoolByClusterUsingGET**
> []ClusterPoolDto ListClusterPoolByClusterUsingGET(ctx, clusterId, projectId)
List cluster pools

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**[]ClusterPoolDto**](ClusterPoolDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterUsingGET**
> PagingClusterDto ListClusterUsingGET(ctx, projectId, optional)
List cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***K8SClusterRestControllerV2ApiListClusterUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8SClusterRestControllerV2ApiListClusterUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingClusterDto**](Paging«ClusterDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListK8sNetworkTypeUsingGET**
> []K8sNetworkTypeDto ListK8sNetworkTypeUsingGET(ctx, projectId)
List K8s Networks Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project Id | 

### Return type

[**[]K8sNetworkTypeDto**](K8sNetworkTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListK8sVersionUsingGET**
> []K8sVersionDto ListK8sVersionUsingGET(ctx, projectId)
List K8s Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project Id | 

### Return type

[**[]K8sVersionDto**](K8sVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeUsingGET**
> []ClusterNodeDto ListNodeUsingGET(ctx, clusterId, projectId)
List node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project Id | 

### Return type

[**[]ClusterNodeDto**](ClusterNodeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecGroupUsingGET**
> []ClusterSecGroupDto ListSecGroupUsingGET(ctx, clusterId, projectId, type_)
Get Cluster Sec Group Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 
  **type_** | **bool**| The type | 

### Return type

[**[]ClusterSecGroupDto**](ClusterSecGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecgroupDefaultUsingGET**
> []ClusterSecGroupDefaultDto ListSecgroupDefaultUsingGET(ctx, clusterId, projectId)
List Secgroup Default

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**[]ClusterSecGroupDefaultDto**](ClusterSecGroupDefaultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeUsingGET3**
> []ClusterVolumeDto ListVolumeUsingGET3(ctx, clusterId, projectId)
List volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 

### Return type

[**[]ClusterVolumeDto**](ClusterVolumeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScaleMinionUsingPOST**
> DataResponseClusterDto ScaleMinionUsingPOST(ctx, clusterId, projectId, scaleMinionBackendRequest)
Scale Minion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**| The cluster id | 
  **projectId** | **string**| The project id | 
  **scaleMinionBackendRequest** | [**ScaleMinionBackendRequest**](ScaleMinionBackendRequest.md)| scaleMinionBackendRequest | 

### Return type

[**DataResponseClusterDto**](DataResponse«ClusterDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecGroupUsingPUT**
> DataResponseListClusterSecGroupDto UpdateSecGroupUsingPUT(ctx, projectId, updateClusterSecGroupRequest)
Update sec group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateClusterSecGroupRequest** | [**UpdateClusterSecGroupRequest**](UpdateClusterSecGroupRequest.md)| updateClusterSecGroupRequest | 

### Return type

[**DataResponseListClusterSecGroupDto**](DataResponse«List«ClusterSecGroupDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

