# \LoadBalancerListenerRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListL7PoliciesUsingGET**](LoadBalancerListenerRestControllerV2Api.md#GetListL7PoliciesUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies | Get list policies of a listener
[**GetListListenersByLoadBalancerUsingGET**](LoadBalancerListenerRestControllerV2Api.md#GetListListenersByLoadBalancerUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners | Get list of listeners of a Load Balancer
[**GetListenersUsingGET**](LoadBalancerListenerRestControllerV2Api.md#GetListenersUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId} | Get a specific listener


# **GetListL7PoliciesUsingGET**
> DataResponseListL7Policy GetListL7PoliciesUsingGET(ctx, listenerId, loadBalancerId, projectId)
Get list policies of a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listenerId** | **string**| The listener id | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListL7Policy**](DataResponse«List«L7Policy»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListListenersByLoadBalancerUsingGET**
> DataResponseListListener GetListListenersByLoadBalancerUsingGET(ctx, loadBalancerId, projectId)
Get list of listeners of a Load Balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| The Load Balancer id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListListener**](DataResponse«List«Listener»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListenersUsingGET**
> DataResponseListener GetListenersUsingGET(ctx, listenerId, loadBalancerId, projectId)
Get a specific listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listenerId** | **string**| The listener id | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListener**](DataResponse«Listener».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

