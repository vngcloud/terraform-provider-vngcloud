# \VirtualIpAddressRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAddressPairUsingPOST**](VirtualIpAddressRestControllerV2Api.md#AddAddressPairUsingPOST) | **Post** /v2/{projectId}/virtualIpAddress/{virtualIpAddressId}/addressPairs | Create Address Pair
[**CreateVirtualIpAddressUsingPOST**](VirtualIpAddressRestControllerV2Api.md#CreateVirtualIpAddressUsingPOST) | **Post** /v2/{projectId}/virtualIpAddress | Create Virtual Ip Address
[**DeleteVirtualIpAddressUsingDELETE**](VirtualIpAddressRestControllerV2Api.md#DeleteVirtualIpAddressUsingDELETE) | **Delete** /v2/{projectId}/virtualIpAddress/{virtualIpAddressId} | Delete virtual ip address
[**GetListAddressPairUsingGET**](VirtualIpAddressRestControllerV2Api.md#GetListAddressPairUsingGET) | **Get** /v2/{projectId}/virtualIpAddress/{virtualIpAddressId}/addressPairs | List address pair interfaces 
[**ListInternalNetworkInterfaceUsingGET**](VirtualIpAddressRestControllerV2Api.md#ListInternalNetworkInterfaceUsingGET) | **Get** /v2/{projectId}/virtualIpAddress/internalNetworkInterfaces | List of internal network interface for adding address pair
[**ListVirtualIpAddressUsingGET**](VirtualIpAddressRestControllerV2Api.md#ListVirtualIpAddressUsingGET) | **Get** /v2/{projectId}/virtualIpAddress | List All Virtual Ip Address
[**RemoveAddressPairUsingDELETE**](VirtualIpAddressRestControllerV2Api.md#RemoveAddressPairUsingDELETE) | **Delete** /v2/{projectId}/virtualIpAddress/{virtualIpAddressId}/addressPairs/{addressPairId} | Delete Address Pair
[**UpdateUsingPUT2**](VirtualIpAddressRestControllerV2Api.md#UpdateUsingPUT2) | **Put** /v2/{projectId}/virtualIpAddress/{virtualIpAddressId} | Update virtual ip address


# **AddAddressPairUsingPOST**
> DataResponseVirtualIpAddressNetworkEntity AddAddressPairUsingPOST(ctx, createAddressPairModelRequest, projectId, virtualIpAddressId)
Create Address Pair

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createAddressPairModelRequest** | [**CreateAddressPairRequest**](CreateAddressPairRequest.md)| createAddressPairModelRequest | 
  **projectId** | **string**| The project id | 
  **virtualIpAddressId** | **string**| The Virtual Ip Address id | 

### Return type

[**DataResponseVirtualIpAddressNetworkEntity**](DataResponse«VirtualIpAddressNetworkEntity».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVirtualIpAddressUsingPOST**
> DataResponseVirtualIpAddressDto CreateVirtualIpAddressUsingPOST(ctx, createVirtualIpAddressRequest, projectId)
Create Virtual Ip Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createVirtualIpAddressRequest** | [**CreateVirtualIpAddressRequest**](CreateVirtualIpAddressRequest.md)| createVirtualIpAddressRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseVirtualIpAddressDto**](DataResponse«VirtualIpAddressDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualIpAddressUsingDELETE**
> DeleteVirtualIpAddressUsingDELETE(ctx, projectId, virtualIpAddressId)
Delete virtual ip address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **virtualIpAddressId** | **string**| The ssh-key id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListAddressPairUsingGET**
> DataResponseListAddressPairDto GetListAddressPairUsingGET(ctx, projectId, virtualIpAddressId)
List address pair interfaces 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **virtualIpAddressId** | **string**| The Virtual Ip Address id | 

### Return type

[**DataResponseListAddressPairDto**](DataResponse«List«AddressPairDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInternalNetworkInterfaceUsingGET**
> DataResponseListInternalNetworkInterfaceDto ListInternalNetworkInterfaceUsingGET(ctx, projectId)
List of internal network interface for adding address pair

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListInternalNetworkInterfaceDto**](DataResponse«List«InternalNetworkInterfaceDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualIpAddressUsingGET**
> PagingVirtualIpAddressDto ListVirtualIpAddressUsingGET(ctx, projectId, optional)
List All Virtual Ip Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***VirtualIpAddressRestControllerV2ApiListVirtualIpAddressUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualIpAddressRestControllerV2ApiListVirtualIpAddressUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 

### Return type

[**PagingVirtualIpAddressDto**](Paging«VirtualIpAddressDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAddressPairUsingDELETE**
> RemoveAddressPairUsingDELETE(ctx, addressPairId, projectId, virtualIpAddressId)
Delete Address Pair

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressPairId** | **string**| The Address Pair id | 
  **projectId** | **string**| The project id | 
  **virtualIpAddressId** | **string**| The Virtual Ip Address id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT2**
> DataResponseVirtualIpAddressDto UpdateUsingPUT2(ctx, projectId, updateSecurityGroupRuleRequest, virtualIpAddressId)
Update virtual ip address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateSecurityGroupRuleRequest** | [**UpdateVirtualIpAddressRequest**](UpdateVirtualIpAddressRequest.md)| updateSecurityGroupRuleRequest | 
  **virtualIpAddressId** | **string**| The virtual Ip AddressId id | 

### Return type

[**DataResponseVirtualIpAddressDto**](DataResponse«VirtualIpAddressDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

