# \FlavorZoneRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlavorZoneUsingGET**](FlavorZoneRestControllerApi.md#GetFlavorZoneUsingGET) | **Get** /v1/{project_id}/flavor_zones/{flavor_zone_id} | getFlavorZone
[**ListFlavorZoneUsingGET**](FlavorZoneRestControllerApi.md#ListFlavorZoneUsingGET) | **Get** /v1/{project_id}/flavor_zones | listFlavorZone


# **GetFlavorZoneUsingGET**
> FlavorZoneResponse GetFlavorZoneUsingGET(ctx, flavorZoneId, projectId)
getFlavorZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flavorZoneId** | **string**| flavor_zone_id | 
  **projectId** | **string**| project_id | 

### Return type

[**FlavorZoneResponse**](FlavorZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFlavorZoneUsingGET**
> FlavorZoneResponse ListFlavorZoneUsingGET(ctx, projectId)
listFlavorZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**FlavorZoneResponse**](FlavorZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

