# \WanIpRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWanIpUsingDELETE**](WanIpRestControllerV2Api.md#DeleteWanIpUsingDELETE) | **Delete** /v2/{projectId}/wanIps/{wanIpId} | Delete a WAN IP
[**ListWanIpUsingGET**](WanIpRestControllerV2Api.md#ListWanIpUsingGET) | **Get** /v2/{projectId}/wanIps | List WAN IP


# **DeleteWanIpUsingDELETE**
> DeleteWanIpUsingDELETE(ctx, projectId, wanIpId)
Delete a WAN IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **wanIpId** | **string**| The WAN IP id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWanIpUsingGET**
> PagingWanIpDto ListWanIpUsingGET(ctx, projectId, optional)
List WAN IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***WanIpRestControllerV2ApiListWanIpUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WanIpRestControllerV2ApiListWanIpUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 

### Return type

[**PagingWanIpDto**](Paging«WanIpDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

