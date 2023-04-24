# \NetworkInterfaceElasticRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkInterfaceElasticUsingPOST**](NetworkInterfaceElasticRestControllerV2Api.md#CreateNetworkInterfaceElasticUsingPOST) | **Post** /v2/{projectId}/network-interfaces-elastic | Create Network Interface Elastic
[**DeleteNetworkInterfaceElasticUsingDELETE**](NetworkInterfaceElasticRestControllerV2Api.md#DeleteNetworkInterfaceElasticUsingDELETE) | **Delete** /v2/{projectId}/network-interfaces-elastic/{networkInterfaceId} | Delete Network Interface Elastic
[**GetNetworkInterfaceElasticUsingGET**](NetworkInterfaceElasticRestControllerV2Api.md#GetNetworkInterfaceElasticUsingGET) | **Get** /v2/{projectId}/network-interfaces-elastic/{networkInterfaceId} | Get Network Interface Elastic
[**ListNetworkInterfaceElasticUsingGET**](NetworkInterfaceElasticRestControllerV2Api.md#ListNetworkInterfaceElasticUsingGET) | **Get** /v2/{projectId}/network-interfaces-elastic | List Network Interface Elastic
[**RenameNetworkInterfaceElasticUsingPUT**](NetworkInterfaceElasticRestControllerV2Api.md#RenameNetworkInterfaceElasticUsingPUT) | **Put** /v2/{projectId}/network-interfaces-elastic/{networkInterfaceId}/rename | Rename Network Interface Elastic


# **CreateNetworkInterfaceElasticUsingPOST**
> DataResponse CreateNetworkInterfaceElasticUsingPOST(ctx, createNetworkInterfaceRequest, projectId)
Create Network Interface Elastic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createNetworkInterfaceRequest** | [**CreateNetworkInterfaceRequest**](CreateNetworkInterfaceRequest.md)| createNetworkInterfaceRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkInterfaceElasticUsingDELETE**
> DeleteNetworkInterfaceElasticUsingDELETE(ctx, networkInterfaceId, projectId)
Delete Network Interface Elastic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkInterfaceId** | **string**| The network interface id | 
  **projectId** | **string**| The project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkInterfaceElasticUsingGET**
> DataResponse GetNetworkInterfaceElasticUsingGET(ctx, networkInterfaceId, projectId)
Get Network Interface Elastic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkInterfaceId** | **string**| The network interface id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkInterfaceElasticUsingGET**
> PagingNetworkInterfaceElastic ListNetworkInterfaceElasticUsingGET(ctx, projectId, optional)
List Network Interface Elastic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***NetworkInterfaceElasticRestControllerV2ApiListNetworkInterfaceElasticUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkInterfaceElasticRestControllerV2ApiListNetworkInterfaceElasticUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingNetworkInterfaceElastic**](Paging«NetworkInterfaceElastic».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameNetworkInterfaceElasticUsingPUT**
> DataResponse RenameNetworkInterfaceElasticUsingPUT(ctx, networkInterfaceId, projectId, renameNetworkInterfaceRequest)
Rename Network Interface Elastic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkInterfaceId** | **string**| The network interface id | 
  **projectId** | **string**| The project id | 
  **renameNetworkInterfaceRequest** | [**RenameNetworkInterfaceRequest**](RenameNetworkInterfaceRequest.md)| renameNetworkInterfaceRequest | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

