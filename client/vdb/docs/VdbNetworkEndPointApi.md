# \VdbNetworkEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListNetworkUsingGET**](VdbNetworkEndPointApi.md#GetListNetworkUsingGET) | **Get** /v1/neutron/{projectId}/networksv2 | getListNetwork
[**GetListSubnetUsingGET**](VdbNetworkEndPointApi.md#GetListSubnetUsingGET) | **Get** /v1/neutron/{projectId}/subnet/networkId/{networkId} | getListSubnet


# **GetListNetworkUsingGET**
> ListNetworkResponse GetListNetworkUsingGET(ctx, projectId)
getListNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**ListNetworkResponse**](ListNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListSubnetUsingGET**
> ListSubnetResponse GetListSubnetUsingGET(ctx, networkId, projectId)
getListSubnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 

### Return type

[**ListSubnetResponse**](ListSubnetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

