# \LoadBalancerRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoadBalancerUsingPOST1**](LoadBalancerRestControllerV2Api.md#CreateLoadBalancerUsingPOST1) | **Post** /v2/{projectId}/loadBalancers | Create a new load balancer
[**DeleteLoadBalancerUsingDELETE1**](LoadBalancerRestControllerV2Api.md#DeleteLoadBalancerUsingDELETE1) | **Delete** /v2/{projectId}/loadBalancers/{loadBalancerId} | Delete a load balancer
[**GetHeadersUsingGET1**](LoadBalancerRestControllerV2Api.md#GetHeadersUsingGET1) | **Get** /v2/{projectId}/loadBalancers/headers | Get list headers
[**GetListLoadBalancerWithPagingUsingGET**](LoadBalancerRestControllerV2Api.md#GetListLoadBalancerWithPagingUsingGET) | **Get** /v2/{projectId}/loadBalancers | List LoadBalancers Paging
[**GetLoadBalancerUsingGET1**](LoadBalancerRestControllerV2Api.md#GetLoadBalancerUsingGET1) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId} | Load Balancer By Load Balancer Id
[**GetPackagesUsingGET1**](LoadBalancerRestControllerV2Api.md#GetPackagesUsingGET1) | **Get** /v2/{projectId}/loadBalancers/packages | Get list packages


# **CreateLoadBalancerUsingPOST1**
> DataResponseLoadBalancer CreateLoadBalancerUsingPOST1(ctx, createLoadBalancerRequestV2, projectId)
Create a new load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createLoadBalancerRequestV2** | [**CreateLoadBalancerRequestV2**](CreateLoadBalancerRequestV2.md)| createLoadBalancerRequestV2 | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseLoadBalancer**](DataResponse«LoadBalancer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerUsingDELETE1**
> DeleteLoadBalancerUsingDELETE1(ctx, loadBalancerId, projectId)
Delete a load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| The load balancer id | 
  **projectId** | **string**| The project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeadersUsingGET1**
> DataResponseListstring GetHeadersUsingGET1(ctx, projectId)
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

# **GetListLoadBalancerWithPagingUsingGET**
> PagingLoadBalancerDto GetListLoadBalancerWithPagingUsingGET(ctx, name, page, projectId, size)
List LoadBalancers Paging

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
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

# **GetLoadBalancerUsingGET1**
> DataResponseLoadBalancerDto GetLoadBalancerUsingGET1(ctx, loadBalancerId, projectId)
Load Balancer By Load Balancer Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loadBalancerId** | **string**| The load balancer id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseLoadBalancerDto**](DataResponse«LoadBalancerDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPackagesUsingGET1**
> PagingLoadBalancerPackage GetPackagesUsingGET1(ctx, projectId)
Get list packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**PagingLoadBalancerPackage**](Paging«LoadBalancerPackage».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

