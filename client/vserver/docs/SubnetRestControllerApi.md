# \SubnetRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnetUsingPOST**](SubnetRestControllerApi.md#CreateSubnetUsingPOST) | **Post** /v1/{project_id}/subnets | createSubnet
[**DeleteSubnetUsingDELETE**](SubnetRestControllerApi.md#DeleteSubnetUsingDELETE) | **Delete** /v1/{project_id}/subnets | deleteSubnet
[**GetMpPublicInterfaceUsingGET**](SubnetRestControllerApi.md#GetMpPublicInterfaceUsingGET) | **Get** /v1/{project_id}/subnets/mp-public-interfaces | getMpPublicInterface
[**GetSubnetUsingGET**](SubnetRestControllerApi.md#GetSubnetUsingGET) | **Get** /v1/{project_id}/subnets/{subnet_id} | getSubnet
[**ListSubnetsByNetworkUsingGET**](SubnetRestControllerApi.md#ListSubnetsByNetworkUsingGET) | **Get** /v1/{project_id}/subnets/networks/{network_id} | listSubnetsByNetwork


# **CreateSubnetUsingPOST**
> SubnetResponse CreateSubnetUsingPOST(ctx, createSubnetRequest, projectId)
createSubnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSubnetRequest** | [**CreateSubnetRequest**](CreateSubnetRequest.md)| createSubnetRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**SubnetResponse**](SubnetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubnetUsingDELETE**
> BaseResponse DeleteSubnetUsingDELETE(ctx, deleteSubnetRequest, projectId)
deleteSubnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteSubnetRequest** | [**DeleteSubnetRequest**](DeleteSubnetRequest.md)| deleteSubnetRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMpPublicInterfaceUsingGET**
> SubnetResponse GetMpPublicInterfaceUsingGET(ctx, projectId)
getMpPublicInterface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**SubnetResponse**](SubnetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubnetUsingGET**
> SubnetResponse GetSubnetUsingGET(ctx, projectId, subnetId)
getSubnet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **subnetId** | **string**| subnet_id | 

### Return type

[**SubnetResponse**](SubnetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubnetsByNetworkUsingGET**
> SubnetResponse ListSubnetsByNetworkUsingGET(ctx, networkId, projectId)
listSubnetsByNetwork

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**| network_id | 
  **projectId** | **string**| project_id | 

### Return type

[**SubnetResponse**](SubnetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

