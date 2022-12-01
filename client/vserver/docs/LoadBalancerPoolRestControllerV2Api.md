# \LoadBalancerPoolRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthMonitorFromPoolUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetHealthMonitorFromPoolUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId}/healthMonitor | Get the health monitor associated with a specific pool
[**GetListPoolsUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetListPoolsUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools | Get list of pools of a load balancer
[**GetMembersFromPoolUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetMembersFromPoolUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId}/members | Get members from a pool
[**GetPoolUsingGET**](LoadBalancerPoolRestControllerV2Api.md#GetPoolUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/pools/{poolId} | Get a specific pool


# **GetHealthMonitorFromPoolUsingGET**
> DataResponseHealthMonitor GetHealthMonitorFromPoolUsingGET(ctx, poolId, projectId)
Get the health monitor associated with a specific pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
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
> DataResponseListPool GetListPoolsUsingGET(ctx, loadBalancerId, projectId)
Get list of pools of a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| The Load Balancer id | 
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
> DataResponseListMember GetMembersFromPoolUsingGET(ctx, poolId, projectId)
Get members from a pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| pool id | 
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
> DataResponsePool GetPoolUsingGET(ctx, loadBalancerId, poolId, projectId)
Get a specific pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| The Load Balancer id | 
  **poolId** | **string**| The listener id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponsePool**](DataResponse«Pool».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

