# \NetworkRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkUsingPOST**](NetworkRestControllerApi.md#CreateNetworkUsingPOST) | **Post** /v1/{project_id}/networks | createNetwork
[**DeleteNetworkUsingDELETE**](NetworkRestControllerApi.md#DeleteNetworkUsingDELETE) | **Delete** /v1/{project_id}/networks | deleteNetwork
[**EditNetworkUsingPUT**](NetworkRestControllerApi.md#EditNetworkUsingPUT) | **Put** /v1/{project_id}/networks | editNetwork
[**GetNetworkUsingGET**](NetworkRestControllerApi.md#GetNetworkUsingGET) | **Get** /v1/{project_id}/networks/{network_id} | getNetwork
[**ListNetworkUsingGET**](NetworkRestControllerApi.md#ListNetworkUsingGET) | **Get** /v1/{project_id}/networks | listNetwork


# **CreateNetworkUsingPOST**
> NetworkResponse CreateNetworkUsingPOST(ctx, createNetworkRequest, projectId)
createNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createNetworkRequest** | [**CreateNetworkRequest**](CreateNetworkRequest.md)| createNetworkRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**NetworkResponse**](NetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkUsingDELETE**
> BaseResponse DeleteNetworkUsingDELETE(ctx, deleteNetworkRequest, projectId)
deleteNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteNetworkRequest** | [**DeleteNetworkRequest**](DeleteNetworkRequest.md)| deleteNetworkRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditNetworkUsingPUT**
> NetworkResponse EditNetworkUsingPUT(ctx, projectId, updateNetworkRequest)
editNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md)| updateNetworkRequest | 

### Return type

[**NetworkResponse**](NetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkUsingGET**
> NetworkResponse GetNetworkUsingGET(ctx, networkId, projectId)
getNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| network_id | 
  **projectId** | **string**| project_id | 

### Return type

[**NetworkResponse**](NetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkUsingGET**
> NetworkResponse ListNetworkUsingGET(ctx, projectId)
listNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**NetworkResponse**](NetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

