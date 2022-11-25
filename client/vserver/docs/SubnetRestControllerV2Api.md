# \SubnetRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnetUsingPOST1**](SubnetRestControllerV2Api.md#CreateSubnetUsingPOST1) | **Post** /v2/{projectId}/networks/{networkId}/subnets | createSubnet
[**DeleteNetworkUsingDELETE2**](SubnetRestControllerV2Api.md#DeleteNetworkUsingDELETE2) | **Delete** /v2/{projectId}/networks/{networkId}/subnets/{subnetId} | Delete subnet
[**EditSubnetUsingPATCH**](SubnetRestControllerV2Api.md#EditSubnetUsingPATCH) | **Patch** /v2/{projectId}/networks/{networkId}/subnets/{subnetId} | Edit subnet
[**GetSubnetByIdUsingGET**](SubnetRestControllerV2Api.md#GetSubnetByIdUsingGET) | **Get** /v2/{projectId}/networks/{networkId}/subnets/{subnetId} | List subnet by id
[**ListSubnetOfNetworkUsingGET**](SubnetRestControllerV2Api.md#ListSubnetOfNetworkUsingGET) | **Get** /v2/{projectId}/networks/{networkId}/subnets | List subnet of network


# **CreateSubnetUsingPOST1**
> DataResponseSubnetDto CreateSubnetUsingPOST1(ctx, createSubnetRequest, networkId, projectId)
createSubnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSubnetRequest** | [**CreateSubnetRequest**](CreateSubnetRequest.md)| createSubnetRequest | 
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 

### Return type

[**DataResponseSubnetDto**](DataResponse«SubnetDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkUsingDELETE2**
> DeleteNetworkUsingDELETE2(ctx, projectId, subnetId, networkId)
Delete subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **subnetId** | **string**| subnetId | 
  **networkId** | **string**| networkId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditSubnetUsingPATCH**
> DataResponseSubnetDto EditSubnetUsingPATCH(ctx, projectId, subnetId, updateSubnetRequest, networkId)
Edit subnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **subnetId** | **string**| subnetId | 
  **updateSubnetRequest** | [**UpdateSubnetRequest**](UpdateSubnetRequest.md)| updateSubnetRequest | 
  **networkId** | **string**| networkId | 

### Return type

[**DataResponseSubnetDto**](DataResponse«SubnetDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubnetByIdUsingGET**
> SubnetDto GetSubnetByIdUsingGET(ctx, networkId, projectId, subnetId)
List subnet by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 
  **subnetId** | **string**| subnetId | 

### Return type

[**SubnetDto**](SubnetDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubnetOfNetworkUsingGET**
> []SubnetDto ListSubnetOfNetworkUsingGET(ctx, networkId, projectId)
List subnet of network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| networkId | 
  **projectId** | **string**| projectId | 

### Return type

[**[]SubnetDto**](SubnetDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

