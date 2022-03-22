# \LoadBalancerRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListenerUsingPOST**](LoadBalancerRestControllerApi.md#CreateListenerUsingPOST) | **Post** /v1/{project_id}/listeners | createListener
[**CreateLoadBalancerUsingPOST**](LoadBalancerRestControllerApi.md#CreateLoadBalancerUsingPOST) | **Post** /v1/{project_id}/load-balancers | Create a new load balancer
[**CreatePolicyUsingPOST**](LoadBalancerRestControllerApi.md#CreatePolicyUsingPOST) | **Post** /v1/{project_id}/l7-policies | createPolicy
[**CreatePoolUsingPOST**](LoadBalancerRestControllerApi.md#CreatePoolUsingPOST) | **Post** /v1/{project_id}/pools | createPool
[**DeleteListenerUsingDELETE**](LoadBalancerRestControllerApi.md#DeleteListenerUsingDELETE) | **Delete** /v1/{project_id}/listeners | deleteListener
[**DeleteLoadBalancerUsingDELETE**](LoadBalancerRestControllerApi.md#DeleteLoadBalancerUsingDELETE) | **Delete** /v1/{project_id}/load-balancers | deleteLoadBalancer
[**DeletePolicyUsingDELETE**](LoadBalancerRestControllerApi.md#DeletePolicyUsingDELETE) | **Delete** /v1/{project_id}/l7-policies | deletePolicy
[**DeletePoolUsingDELETE**](LoadBalancerRestControllerApi.md#DeletePoolUsingDELETE) | **Delete** /v1/{project_id}/pools | deletePool
[**GetHealthMonitorFromPoolUsingGET**](LoadBalancerRestControllerApi.md#GetHealthMonitorFromPoolUsingGET) | **Get** /v1/{project_id}/pools/{pool_id}/health-monitors | getHealthMonitorFromPool
[**GetHealthMonitorUsingGET**](LoadBalancerRestControllerApi.md#GetHealthMonitorUsingGET) | **Get** /v1/{project_id}/health-monitors/{health_monitor_id} | getHealthMonitor
[**GetL7PoliciesUsingGET**](LoadBalancerRestControllerApi.md#GetL7PoliciesUsingGET) | **Get** /v1/{project_id}/listeners/{listener_id}/l7policies | getL7Policies
[**GetListListenersUsingGET**](LoadBalancerRestControllerApi.md#GetListListenersUsingGET) | **Get** /v1/{project_id}/load-balancers/{load_balancer_id}/listeners | Get list of listeners of a load balancer
[**GetListLoadBalancersUsingGET**](LoadBalancerRestControllerApi.md#GetListLoadBalancersUsingGET) | **Get** /v1/{project_id}/load-balancers | Get list of load balancers
[**GetListPoolsUsingGET**](LoadBalancerRestControllerApi.md#GetListPoolsUsingGET) | **Get** /v1/{project_id}/load-balancers/{load_balancer_id}/pools | Get list of pools of a load balancer
[**GetListenerUsingGET**](LoadBalancerRestControllerApi.md#GetListenerUsingGET) | **Get** /v1/{project_id}/listeners/{listener_id} | Get a specific listener
[**GetLoadBalancerUsingGET**](LoadBalancerRestControllerApi.md#GetLoadBalancerUsingGET) | **Get** /v1/{project_id}/load-balancers/{load_balancer_id} | Get a specific load balancer
[**GetMemberFromPoolUsingGET**](LoadBalancerRestControllerApi.md#GetMemberFromPoolUsingGET) | **Get** /v1/{project_id}/pools/{pool_id}/members | getMemberFromPool
[**GetPackagesUsingGET**](LoadBalancerRestControllerApi.md#GetPackagesUsingGET) | **Get** /v1/{project_id}/load-balancers/packages | getPackages
[**GetPoolUsingGET**](LoadBalancerRestControllerApi.md#GetPoolUsingGET) | **Get** /v1/{project_id}/pools/{pool_id} | Get a specific pool
[**ReorderPoliciesUsingPUT**](LoadBalancerRestControllerApi.md#ReorderPoliciesUsingPUT) | **Put** /v1/{project_id}/reorder-l7-policies | reorderPolicies
[**UpdateAPolicyUsingPUT**](LoadBalancerRestControllerApi.md#UpdateAPolicyUsingPUT) | **Put** /v1/{project_id}/l7-policies | updateAPolicy
[**UpdateListenerUsingPUT**](LoadBalancerRestControllerApi.md#UpdateListenerUsingPUT) | **Put** /v1/{project_id}/listeners | updateListener
[**UpdateMembersUsingPUT**](LoadBalancerRestControllerApi.md#UpdateMembersUsingPUT) | **Put** /v1/{project_id}/pools/{pool_id}/members | updateMembers
[**UpdatePoolUsingPUT**](LoadBalancerRestControllerApi.md#UpdatePoolUsingPUT) | **Put** /v1/{project_id}/pools | updatePool


# **CreateListenerUsingPOST**
> DataResponseListener CreateListenerUsingPOST(ctx, createListenerRequest, projectId)
createListener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createListenerRequest** | [**CreateListenerRequest**](CreateListenerRequest.md)| createListenerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListener**](DataResponseListener.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerUsingPOST**
> DataResponseLoadBalancer CreateLoadBalancerUsingPOST(ctx, createLoadBalancerRequest, projectId)
Create a new load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createLoadBalancerRequest** | [**CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md)| createLoadBalancerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseLoadBalancer**](DataResponseLoadBalancer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyUsingPOST**
> DataResponseL7Policy CreatePolicyUsingPOST(ctx, createL7PolicyRequest, projectId)
createPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createL7PolicyRequest** | [**CreateL7PolicyRequest**](CreateL7PolicyRequest.md)| createL7PolicyRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseL7Policy**](DataResponseL7Policy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePoolUsingPOST**
> DataResponsePool CreatePoolUsingPOST(ctx, createPoolRequest, projectId)
createPool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPoolRequest** | [**CreatePoolRequest**](CreatePoolRequest.md)| createPoolRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponsePool**](DataResponsePool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListenerUsingDELETE**
> DataResponseListener DeleteListenerUsingDELETE(ctx, deleteListenerRequest, projectId)
deleteListener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListenerRequest** | [**DeleteListenerRequest**](DeleteListenerRequest.md)| deleteListenerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListener**](DataResponseListener.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerUsingDELETE**
> DataResponseLoadBalancer DeleteLoadBalancerUsingDELETE(ctx, deleteLoadBalancerRequest, projectId)
deleteLoadBalancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteLoadBalancerRequest** | [**DeleteLoadBalancerRequest**](DeleteLoadBalancerRequest.md)| deleteLoadBalancerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseLoadBalancer**](DataResponseLoadBalancer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyUsingDELETE**
> DataResponseL7Policy DeletePolicyUsingDELETE(ctx, deleteL7PolicyRequest, projectId)
deletePolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteL7PolicyRequest** | [**DeleteL7PolicyRequest**](DeleteL7PolicyRequest.md)| deleteL7PolicyRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseL7Policy**](DataResponseL7Policy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePoolUsingDELETE**
> DataResponsePool DeletePoolUsingDELETE(ctx, deletePoolRequest, projectId)
deletePool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deletePoolRequest** | [**DeletePoolRequest**](DeletePoolRequest.md)| deletePoolRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponsePool**](DataResponsePool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHealthMonitorFromPoolUsingGET**
> DataResponseHealthMonitor GetHealthMonitorFromPoolUsingGET(ctx, poolId, projectId)
getHealthMonitorFromPool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseHealthMonitor**](DataResponseHealthMonitor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHealthMonitorUsingGET**
> DataResponseHealthMonitor GetHealthMonitorUsingGET(ctx, healthMonitorId, projectId)
getHealthMonitor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **healthMonitorId** | **string**| health monitor id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseHealthMonitor**](DataResponseHealthMonitor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL7PoliciesUsingGET**
> PagingL7Policy GetL7PoliciesUsingGET(ctx, listenerId, projectId, optional)
getL7Policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listenerId** | **string**| listener id | 
  **projectId** | **string**| project id | 
 **optional** | ***LoadBalancerRestControllerApiGetL7PoliciesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerRestControllerApiGetL7PoliciesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.String**| page | 
 **size** | **optional.String**| size | 

### Return type

[**PagingL7Policy**](PagingL7Policy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListListenersUsingGET**
> PagingListener GetListListenersUsingGET(ctx, loadBalancerId, projectId, optional)
Get list of listeners of a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| load balancer id | 
  **projectId** | **string**| project id | 
 **optional** | ***LoadBalancerRestControllerApiGetListListenersUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerRestControllerApiGetListListenersUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.String**| page | 
 **size** | **optional.String**| size | 

### Return type

[**PagingListener**](PagingListener.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListLoadBalancersUsingGET**
> PagingLoadBalancer GetListLoadBalancersUsingGET(ctx, projectId, optional)
Get list of load balancers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***LoadBalancerRestControllerApiGetListLoadBalancersUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerRestControllerApiGetListLoadBalancersUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| page | 
 **size** | **optional.String**| size | 

### Return type

[**PagingLoadBalancer**](PagingLoadBalancer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListPoolsUsingGET**
> PagingPool GetListPoolsUsingGET(ctx, loadBalancerId, projectId, optional)
Get list of pools of a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| load balancer id | 
  **projectId** | **string**| project id | 
 **optional** | ***LoadBalancerRestControllerApiGetListPoolsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerRestControllerApiGetListPoolsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.String**| page | 
 **size** | **optional.String**| size | 

### Return type

[**PagingPool**](PagingPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListenerUsingGET**
> DataResponseListener GetListenerUsingGET(ctx, listenerId, projectId)
Get a specific listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listenerId** | **string**| listener id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListener**](DataResponseListener.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerUsingGET**
> DataResponseLoadBalancer GetLoadBalancerUsingGET(ctx, loadBalancerId, projectId)
Get a specific load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| load balancer id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseLoadBalancer**](DataResponseLoadBalancer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMemberFromPoolUsingGET**
> PagingMember GetMemberFromPoolUsingGET(ctx, poolId, projectId, optional)
getMemberFromPool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
  **projectId** | **string**| project id | 
 **optional** | ***LoadBalancerRestControllerApiGetMemberFromPoolUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoadBalancerRestControllerApiGetMemberFromPoolUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.String**| page | 
 **size** | **optional.String**| size | 

### Return type

[**PagingMember**](PagingMember.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackagesUsingGET**
> PagingLoadBalancerPackage GetPackagesUsingGET(ctx, projectId)
getPackages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 

### Return type

[**PagingLoadBalancerPackage**](PagingLoadBalancerPackage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolUsingGET**
> DataResponsePool GetPoolUsingGET(ctx, poolId, projectId)
Get a specific pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponsePool**](DataResponsePool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderPoliciesUsingPUT**
> DataResponseL7Policy ReorderPoliciesUsingPUT(ctx, projectId, reorderPoliciesRequest)
reorderPolicies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **reorderPoliciesRequest** | [**ReorderPoliciesRequest**](ReorderPoliciesRequest.md)| reorderPoliciesRequest | 

### Return type

[**DataResponseL7Policy**](DataResponseL7Policy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAPolicyUsingPUT**
> DataResponseL7Policy UpdateAPolicyUsingPUT(ctx, projectId, updateL7PolicyRequest)
updateAPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateL7PolicyRequest** | [**UpdateL7PolicyRequest**](UpdateL7PolicyRequest.md)| updateL7PolicyRequest | 

### Return type

[**DataResponseL7Policy**](DataResponseL7Policy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateListenerUsingPUT**
> DataResponseListener UpdateListenerUsingPUT(ctx, projectId, updateListenerRequest)
updateListener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateListenerRequest** | [**UpdateListenerRequest**](UpdateListenerRequest.md)| updateListenerRequest | 

### Return type

[**DataResponseListener**](DataResponseListener.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMembersUsingPUT**
> UpdateMembersUsingPUT(ctx, poolId, projectId, updateMembersRequest)
updateMembers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
  **projectId** | **string**| project id | 
  **updateMembersRequest** | [**UpdateMembersRequest**](UpdateMembersRequest.md)| updateMembersRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePoolUsingPUT**
> DataResponsePool UpdatePoolUsingPUT(ctx, projectId, updatePoolRequest)
updatePool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updatePoolRequest** | [**UpdatePoolRequest**](UpdatePoolRequest.md)| updatePoolRequest | 

### Return type

[**DataResponsePool**](DataResponsePool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

