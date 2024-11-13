# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListResource**](ResourceControllerApi.md#GetListResource) | **Get** /v1/resources | 
[**UpdateAutoRenew**](ResourceControllerApi.md#UpdateAutoRenew) | **Put** /v1/resources/{artifactId}/autoRenew | 

# **GetListResource**
> WrapContentResourcesBillingInfo GetListResource(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***ResourceControllerApiGetListResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResourceControllerApiGetListResourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | **optional.String**|  | 
 **renewType** | **optional.String**|  | 
 **status** | **optional.String**|  | 

### Return type

[**WrapContentResourcesBillingInfo**](WrapContentResourcesBillingInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAutoRenew**
> UpdateAutoRenew(ctx, body, artifactId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AutoRenewRequest**](AutoRenewRequest.md)|  | 
  **artifactId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

