# \NetworkRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkUsingPOST1**](NetworkRestControllerV2Api.md#CreateNetworkUsingPOST1) | **Post** /v2/{projectId}/networks | Create network
[**DeleteNetworkUsingDELETE1**](NetworkRestControllerV2Api.md#DeleteNetworkUsingDELETE1) | **Delete** /v2/{projectId}/networks/{networkId} | Delete network
[**EditNetworkUsingPATCH**](NetworkRestControllerV2Api.md#EditNetworkUsingPATCH) | **Patch** /v2/{projectId}/networks/{networkId} | Edit network
[**GetNetworkUsingGET1**](NetworkRestControllerV2Api.md#GetNetworkUsingGET1) | **Get** /v2/{projectId}/networks/{networkId} | Get network
[**ListNetworkActiveUsingGET**](NetworkRestControllerV2Api.md#ListNetworkActiveUsingGET) | **Get** /v2/{projectId}/networks/active | List network active
[**ListNetworkUsingGET1**](NetworkRestControllerV2Api.md#ListNetworkUsingGET1) | **Get** /v2/{projectId}/networks | List network


# **CreateNetworkUsingPOST1**
> DataResponseNetworkDto CreateNetworkUsingPOST1(ctx, createNetworkRequest, projectId)
Create network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createNetworkRequest** | [**CreateNetworkRequest**](CreateNetworkRequest.md)| createNetworkRequest | 
  **projectId** | **string**| projectId | 

### Return type

[**DataResponseNetworkDto**](DataResponse«NetworkDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkUsingDELETE1**
> DeleteNetworkUsingDELETE1(ctx, networkId, projectId)
Delete network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditNetworkUsingPATCH**
> DataResponseNetworkDto EditNetworkUsingPATCH(ctx, networkId, projectId, updateNetworkRequest)
Edit network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 
  **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md)| updateNetworkRequest | 

### Return type

[**DataResponseNetworkDto**](DataResponse«NetworkDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkUsingGET1**
> NetworkDto GetNetworkUsingGET1(ctx, networkId, projectId)
Get network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 

### Return type

[**NetworkDto**](NetworkDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkActiveUsingGET**
> PagingNetworkDto ListNetworkActiveUsingGET(ctx, projectId, optional)
List network active

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***NetworkRestControllerV2ApiListNetworkActiveUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkRestControllerV2ApiListNetworkActiveUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingNetworkDto**](Paging«NetworkDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkUsingGET1**
> PagingNetworkDto ListNetworkUsingGET1(ctx, projectId, optional)
List network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***NetworkRestControllerV2ApiListNetworkUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkRestControllerV2ApiListNetworkUsingGET1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingNetworkDto**](Paging«NetworkDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

