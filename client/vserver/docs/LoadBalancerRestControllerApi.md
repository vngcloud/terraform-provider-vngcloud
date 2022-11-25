# \LoadBalancerRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListenerUsingPOST**](LoadBalancerRestControllerApi.md#CreateListenerUsingPOST) | **Post** /v1/{project_id}/listeners | Create a listener
[**CreateLoadBalancerUsingPOST**](LoadBalancerRestControllerApi.md#CreateLoadBalancerUsingPOST) | **Post** /v1/{project_id}/load-balancers | Create a new load balancer
[**CreatePolicyUsingPOST**](LoadBalancerRestControllerApi.md#CreatePolicyUsingPOST) | **Post** /v1/{project_id}/l7-policies | Create a l7 policy
[**CreatePoolUsingPOST**](LoadBalancerRestControllerApi.md#CreatePoolUsingPOST) | **Post** /v1/{project_id}/pools | Create a new pool
[**DeleteListenerUsingDELETE**](LoadBalancerRestControllerApi.md#DeleteListenerUsingDELETE) | **Delete** /v1/{project_id}/listeners | Delete a listener
[**DeleteLoadBalancerUsingDELETE**](LoadBalancerRestControllerApi.md#DeleteLoadBalancerUsingDELETE) | **Delete** /v1/{project_id}/load-balancers | Delete a load balancer
[**DeletePolicyUsingDELETE**](LoadBalancerRestControllerApi.md#DeletePolicyUsingDELETE) | **Delete** /v1/{project_id}/l7-policies | Delete a policy
[**DeletePoolUsingDELETE**](LoadBalancerRestControllerApi.md#DeletePoolUsingDELETE) | **Delete** /v1/{project_id}/pools | Delete a pool
[**GetHeadersUsingGET**](LoadBalancerRestControllerApi.md#GetHeadersUsingGET) | **Get** /v1/{project_id}/load-balancers/headers | Get list headers
[**GetHealthMonitorFromPoolUsingGET**](LoadBalancerRestControllerApi.md#GetHealthMonitorFromPoolUsingGET) | **Get** /v1/{project_id}/pools/{pool_id}/health-monitors | Get the health monitor associated with a specific pool
[**GetHealthMonitorUsingGET**](LoadBalancerRestControllerApi.md#GetHealthMonitorUsingGET) | **Get** /v1/{project_id}/health-monitors/{health_monitor_id} | getHealthMonitor
[**GetL7PoliciesUsingGET**](LoadBalancerRestControllerApi.md#GetL7PoliciesUsingGET) | **Get** /v1/{project_id}/listeners/{listener_id}/l7policies | Get list policies of a listener
[**GetListListenersUsingGET**](LoadBalancerRestControllerApi.md#GetListListenersUsingGET) | **Get** /v1/{project_id}/load-balancers/{load_balancer_id}/listeners | Get list of listeners of a load balancer
[**GetListLoadBalancersUsingGET**](LoadBalancerRestControllerApi.md#GetListLoadBalancersUsingGET) | **Get** /v1/{project_id}/load-balancers | Get list of load balancers
[**GetListPoolsUsingGET**](LoadBalancerRestControllerApi.md#GetListPoolsUsingGET) | **Get** /v1/{project_id}/load-balancers/{load_balancer_id}/pools | Get list of pools of a load balancer
[**GetListenerUsingGET**](LoadBalancerRestControllerApi.md#GetListenerUsingGET) | **Get** /v1/{project_id}/listeners/{listener_id} | Get a specific listener
[**GetLoadBalancerUsingGET**](LoadBalancerRestControllerApi.md#GetLoadBalancerUsingGET) | **Get** /v1/{project_id}/load-balancers/{load_balancer_id} | Get a specific load balancer
[**GetMemberFromPoolUsingGET**](LoadBalancerRestControllerApi.md#GetMemberFromPoolUsingGET) | **Get** /v1/{project_id}/pools/{pool_id}/members | Get members from a pool
[**GetPackagesUsingGET**](LoadBalancerRestControllerApi.md#GetPackagesUsingGET) | **Get** /v1/{project_id}/load-balancers/packages | Get list packages
[**GetPoolUsingGET**](LoadBalancerRestControllerApi.md#GetPoolUsingGET) | **Get** /v1/{project_id}/pools/{pool_id} | Get a specific pool
[**ReorderPoliciesUsingPUT**](LoadBalancerRestControllerApi.md#ReorderPoliciesUsingPUT) | **Put** /v1/{project_id}/reorder-l7-policies | Reorder the L7 policies
[**UpdateAPolicyUsingPUT**](LoadBalancerRestControllerApi.md#UpdateAPolicyUsingPUT) | **Put** /v1/{project_id}/l7-policies | Update a policy
[**UpdateListenerUsingPUT**](LoadBalancerRestControllerApi.md#UpdateListenerUsingPUT) | **Put** /v1/{project_id}/listeners | Update a listener
[**UpdateMembersUsingPUT**](LoadBalancerRestControllerApi.md#UpdateMembersUsingPUT) | **Put** /v1/{project_id}/pools/{pool_id}/members | Update members
[**UpdatePoolUsingPUT**](LoadBalancerRestControllerApi.md#UpdatePoolUsingPUT) | **Put** /v1/{project_id}/pools | Update a pool


# **CreateListenerUsingPOST**
> DataResponseListener CreateListenerUsingPOST(ctx, createListenerRequest, projectId)
Create a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createListenerRequest** | [**CreateListenerRequest**](CreateListenerRequest.md)| createListenerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListener**](DataResponse«Listener».md)

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

[**DataResponseLoadBalancer**](DataResponse«LoadBalancer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyUsingPOST**
> DataResponseL7Policy CreatePolicyUsingPOST(ctx, createL7PolicyRequest, projectId)
Create a l7 policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createL7PolicyRequest** | [**CreateL7PolicyRequest**](CreateL7PolicyRequest.md)| createL7PolicyRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseL7Policy**](DataResponse«L7Policy».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePoolUsingPOST**
> DataResponsePool CreatePoolUsingPOST(ctx, createPoolRequest, projectId)
Create a new pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPoolRequest** | [**CreatePoolRequest**](CreatePoolRequest.md)| createPoolRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponsePool**](DataResponse«Pool».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListenerUsingDELETE**
> DeleteListenerUsingDELETE(ctx, deleteListenerRequest, projectId)
Delete a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListenerRequest** | [**DeleteListenerRequest**](DeleteListenerRequest.md)| deleteListenerRequest | 
  **projectId** | **string**| project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerUsingDELETE**
> DeleteLoadBalancerUsingDELETE(ctx, deleteLoadBalancerRequest, projectId)
Delete a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteLoadBalancerRequest** | [**DeleteLoadBalancerRequest**](DeleteLoadBalancerRequest.md)| deleteLoadBalancerRequest | 
  **projectId** | **string**| project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyUsingDELETE**
> DeletePolicyUsingDELETE(ctx, deleteL7PolicyRequest, projectId)
Delete a policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteL7PolicyRequest** | [**DeleteL7PolicyRequest**](DeleteL7PolicyRequest.md)| deleteL7PolicyRequest | 
  **projectId** | **string**| project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePoolUsingDELETE**
> DeletePoolUsingDELETE(ctx, deletePoolRequest, projectId)
Delete a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deletePoolRequest** | [**DeletePoolRequest**](DeletePoolRequest.md)| deletePoolRequest | 
  **projectId** | **string**| project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeadersUsingGET**
> DataResponseListstring GetHeadersUsingGET(ctx, projectId)
Get list headers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListstring**](DataResponse«List«string»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHealthMonitorFromPoolUsingGET**
> DataResponseHealthMonitor GetHealthMonitorFromPoolUsingGET(ctx, poolId, projectId)
Get the health monitor associated with a specific pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseHealthMonitor**](DataResponse«HealthMonitor».md)

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

[**DataResponseHealthMonitor**](DataResponse«HealthMonitor».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL7PoliciesUsingGET**
> PagingL7Policy GetL7PoliciesUsingGET(ctx, listenerId, projectId, optional)
Get list policies of a listener

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

[**PagingL7Policy**](Paging«L7Policy».md)

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

[**PagingListener**](Paging«Listener».md)

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

[**PagingLoadBalancer**](Paging«LoadBalancer».md)

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

[**PagingPool**](Paging«Pool».md)

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

[**DataResponseListener**](DataResponse«Listener».md)

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

[**DataResponseLoadBalancer**](DataResponse«LoadBalancer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMemberFromPoolUsingGET**
> PagingMember GetMemberFromPoolUsingGET(ctx, poolId, projectId, optional)
Get members from a pool

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

[**PagingMember**](Paging«Member».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackagesUsingGET**
> PagingLoadBalancerPackage GetPackagesUsingGET(ctx, projectId)
Get list packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 

### Return type

[**PagingLoadBalancerPackage**](Paging«LoadBalancerPackage».md)

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

[**DataResponsePool**](DataResponse«Pool».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderPoliciesUsingPUT**
> ReorderPoliciesUsingPUT(ctx, projectId, reorderPoliciesRequest)
Reorder the L7 policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **reorderPoliciesRequest** | [**ReorderPoliciesRequest**](ReorderPoliciesRequest.md)| reorderPoliciesRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAPolicyUsingPUT**
> UpdateAPolicyUsingPUT(ctx, projectId, updateL7PolicyRequest)
Update a policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateL7PolicyRequest** | [**UpdateL7PolicyRequest**](UpdateL7PolicyRequest.md)| updateL7PolicyRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateListenerUsingPUT**
> UpdateListenerUsingPUT(ctx, projectId, updateListenerRequest)
Update a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateListenerRequest** | [**UpdateListenerRequest**](UpdateListenerRequest.md)| updateListenerRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMembersUsingPUT**
> UpdateMembersUsingPUT(ctx, poolId, projectId, updateMembersRequest)
Update members

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
> UpdatePoolUsingPUT(ctx, projectId, updatePoolRequest)
Update a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updatePoolRequest** | [**UpdatePoolRequest**](UpdatePoolRequest.md)| updatePoolRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

