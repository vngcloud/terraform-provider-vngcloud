# \FlavorZoneRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlavorZoneUsingGET**](FlavorZoneRestControllerApi.md#GetFlavorZoneUsingGET) | **Get** /v1/{project_id}/flavor_zones/{flavor_zone_id} | getFlavorZone
[**ListCodeUsingGET**](FlavorZoneRestControllerApi.md#ListCodeUsingGET) | **Get** /v1/{project_id}/flavor_zones/codes | listCode
[**ListCustomUsingGET**](FlavorZoneRestControllerApi.md#ListCustomUsingGET) | **Get** /v1/{project_id}/flavor_zones/customs | listCustom
[**ListFamilyUsingGET**](FlavorZoneRestControllerApi.md#ListFamilyUsingGET) | **Get** /v1/{project_id}/flavor_zones/families | listFamily
[**ListFlavorZoneUsingGET**](FlavorZoneRestControllerApi.md#ListFlavorZoneUsingGET) | **Get** /v1/{project_id}/flavor_zones/product | listFlavorZone
[**ListFlavorZoneUsingGET1**](FlavorZoneRestControllerApi.md#ListFlavorZoneUsingGET1) | **Get** /v1/{project_id}/flavor_zones/product/{product} | listFlavorZone


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

# **ListCodeUsingGET**
> []Code ListCodeUsingGET(ctx, projectId)
listCode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**[]Code**](Code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomUsingGET**
> []FlavorZone ListCustomUsingGET(ctx, projectId)
listCustom

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**[]FlavorZone**](FlavorZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFamilyUsingGET**
> []Family ListFamilyUsingGET(ctx, projectId)
listFamily

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**[]Family**](Family.md)

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

# **ListFlavorZoneUsingGET1**
> FlavorZoneResponse ListFlavorZoneUsingGET1(ctx, product, projectId)
listFlavorZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **product** | **string**| product | 
  **projectId** | **string**| project_id | 

### Return type

[**FlavorZoneResponse**](FlavorZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

