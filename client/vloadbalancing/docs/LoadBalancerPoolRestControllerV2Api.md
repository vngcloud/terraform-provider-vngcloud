# \LoadBalancerPoolRestControllerV2Api

All URIs are relative to *https://localhost:8089*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePoolUsingPOST**](LoadBalancerPoolRestControllerV2Api.md#CreatePoolUsingPOST) | **Post** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools | Create a new pool
[**DeletePoolUsingDELETE**](LoadBalancerPoolRestControllerV2Api.md#DeletePoolUsingDELETE) | **Delete** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId} | Delete a pool
[**GetHealthMonitorFromPoolUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetHealthMonitorFromPoolUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId}/healthMonitor | Get the health monitor associated with a specific pool
[**GetListPoolsUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetListPoolsUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools | Get list of pools of a load balancer
[**GetMembersFromPoolUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetMembersFromPoolUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId}/members | Get members from a pool
[**GetPoolUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetPoolUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId} | Get a specific pool
[**UpdateMembersUsingPUT**](LoadBalancerPoolRestControllerV2Api.md#UpdateMembersUsingPUT) | **Put** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId}/members | Update list members
[**UpdatePoolUsingPUT**](LoadBalancerPoolRestControllerV2Api.md#UpdatePoolUsingPUT) | **Put** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId} | Update a pool


# **CreatePoolUsingPOST**
> DataResponsePool CreatePoolUsingPOST(ctx, authorization, createPoolRequestV2, loadBalancerId, portalUserId, projectId)
Create a new pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **createPoolRequestV2** | [**CreatePoolRequestV2**](CreatePoolRequestV2.md)| createPoolRequestV2 | 
  **loadBalancerId** | **string**| The load balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponsePool**](DataResponse«Pool».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePoolUsingDELETE**
> DeletePoolUsingDELETE(ctx, authorization, poolId, portalUserId, projectId)
Delete a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **poolId** | **string**| The pool id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHealthMonitorFromPoolUsingGET**
> DataResponseHealthMonitor GetHealthMonitorFromPoolUsingGET(ctx, authorization, poolId, portalUserId, projectId)
Get the health monitor associated with a specific pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **poolId** | **string**| pool id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseHealthMonitor**](DataResponse«HealthMonitor».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListPoolsUsingGET**
> DataResponseListPool GetListPoolsUsingGET(ctx, authorization, loadBalancerId, portalUserId, projectId)
Get list of pools of a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListPool**](DataResponse«List«Pool»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMembersFromPoolUsingGET**
> DataResponseListMember GetMembersFromPoolUsingGET(ctx, authorization, poolId, portalUserId, projectId)
Get members from a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **poolId** | **string**| pool id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListMember**](DataResponse«List«Member»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolUsingGET**
> DataResponsePool GetPoolUsingGET(ctx, authorization, loadBalancerId, poolId, portalUserId, projectId)
Get a specific pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **poolId** | **string**| The listener id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponsePool**](DataResponse«Pool».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMembersUsingPUT**
> map[string]string UpdateMembersUsingPUT(ctx, authorization, poolId, portalUserId, projectId, updateMembersRequestV2)
Update list members

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **poolId** | **string**| The pool id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 
  **updateMembersRequestV2** | [**UpdateMembersRequestV2**](UpdateMembersRequestV2.md)| updateMembersRequestV2 | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePoolUsingPUT**
> DataResponsePool UpdatePoolUsingPUT(ctx, authorization, poolId, portalUserId, projectId, updatePoolRequestV2)
Update a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **poolId** | **string**| The pool id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 
  **updatePoolRequestV2** | [**UpdatePoolRequestV2**](UpdatePoolRequestV2.md)| updatePoolRequestV2 | 

### Return type

[**DataResponsePool**](DataResponse«Pool».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

