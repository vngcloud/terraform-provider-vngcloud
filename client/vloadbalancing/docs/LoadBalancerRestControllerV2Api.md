# \LoadBalancerRestControllerV2Api

All URIs are relative to *https://localhost:8089*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancerUsingPOST**](LoadBalancerRestControllerV2Api.md#CreateLoadBalancerUsingPOST) | **Post** /v2/{projectId}/loadBalancers | Create a new load balancer
[**DeleteLoadBalancerUsingDELETE**](LoadBalancerRestControllerV2Api.md#DeleteLoadBalancerUsingDELETE) | **Delete** /v2/{projectId}/loadBalancers/{loadBalancerId} | Delete a load balancer
[**GetHeadersUsingGET**](LoadBalancerRestControllerV2Api.md#GetHeadersUsingGET) | **Get** /v2/{projectId}/loadBalancers/headers | Get list headers
[**GetListLoadBalancerWithPagingUsingGET**](LoadBalancerRestControllerV2Api.md#GetListLoadBalancerWithPagingUsingGET) | **Get** /v2/{projectId}/loadBalancers | List LoadBalancers Paging
[**GetLoadBalancerUsingGET**](LoadBalancerRestControllerV2Api.md#GetLoadBalancerUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId} | Load Balancer By Load Balancer Id
[**GetPackagesUsingGET**](LoadBalancerRestControllerV2Api.md#GetPackagesUsingGET) | **Get** /v2/{projectId}/loadBalancers/packages | Get list packages


# **CreateLoadBalancerUsingPOST**
> DataResponseLoadBalancer CreateLoadBalancerUsingPOST(ctx, authorization, createLoadBalancerRequestV2, portalUserId, projectId)
Create a new load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **createLoadBalancerRequestV2** | [**CreateLoadBalancerRequestV2**](CreateLoadBalancerRequestV2.md)| createLoadBalancerRequestV2 | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseLoadBalancer**](DataResponse«LoadBalancer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerUsingDELETE**
> DeleteLoadBalancerUsingDELETE(ctx, authorization, loadBalancerId, portalUserId, projectId)
Delete a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **loadBalancerId** | **string**| The load balancer id | 
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

# **GetHeadersUsingGET**
> DataResponseListstring GetHeadersUsingGET(ctx, authorization, portalUserId, projectId)
Get list headers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListstring**](DataResponse«List«string»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListLoadBalancerWithPagingUsingGET**
> PagingLoadBalancerDto GetListLoadBalancerWithPagingUsingGET(ctx, authorization, name, page, portalUserId, projectId, size)
List LoadBalancers Paging

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| project id | 
  **size** | **string**| size | [default to 10]

### Return type

[**PagingLoadBalancerDto**](Paging«LoadBalancerDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerUsingGET**
> DataResponseLoadBalancerDto GetLoadBalancerUsingGET(ctx, authorization, loadBalancerId, portalUserId, projectId)
Load Balancer By Load Balancer Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **loadBalancerId** | **string**| The load balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseLoadBalancerDto**](DataResponse«LoadBalancerDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackagesUsingGET**
> PagingLoadBalancerPackage GetPackagesUsingGET(ctx, authorization, portalUserId, projectId)
Get list packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**PagingLoadBalancerPackage**](Paging«LoadBalancerPackage».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

