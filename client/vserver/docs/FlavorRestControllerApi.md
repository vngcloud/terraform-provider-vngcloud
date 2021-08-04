# \FlavorRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlavorUsingGET**](FlavorRestControllerApi.md#GetFlavorUsingGET) | **Get** /v1/{project_id}/flavors/{flavor_id} | getFlavor
[**ListFlavorUsingGET**](FlavorRestControllerApi.md#ListFlavorUsingGET) | **Get** /v1/{project_id}/{flavor_zone_id}/flavors | listFlavor


# **GetFlavorUsingGET**
> FlavorResponse GetFlavorUsingGET(ctx, flavorId, projectId)
getFlavor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flavorId** | **string**| flavor_id | 
  **projectId** | **string**| project_id | 

### Return type

[**FlavorResponse**](FlavorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFlavorUsingGET**
> FlavorResponse ListFlavorUsingGET(ctx, flavorZoneId, projectId)
listFlavor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flavorZoneId** | **string**| flavor_zone_id | 
  **projectId** | **string**| project_id | 

### Return type

[**FlavorResponse**](FlavorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

