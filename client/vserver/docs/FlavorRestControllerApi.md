# \FlavorRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlavorUsingGET**](FlavorRestControllerApi.md#GetFlavorUsingGET) | **Get** /v1/{project_id}/flavors/{flavor_id} | getFlavor
[**ListFlavorByFamilyAndPlatformUsingGET**](FlavorRestControllerApi.md#ListFlavorByFamilyAndPlatformUsingGET) | **Get** /v1/{project_id}/flavors/families/{familyKey}/platforms/{platformKey} | List flavor custom for project
[**ListFlavorCustomUsingGET**](FlavorRestControllerApi.md#ListFlavorCustomUsingGET) | **Get** /v1/{project_id}/flavors/customs | List flavor custom for project
[**ListFlavorForSimpleServerUsingGET**](FlavorRestControllerApi.md#ListFlavorForSimpleServerUsingGET) | **Get** /v1/{project_id}/flavors/simple-servers | List flavor for simple server
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

# **ListFlavorByFamilyAndPlatformUsingGET**
> []Flavor ListFlavorByFamilyAndPlatformUsingGET(ctx, familyKey, platformKey, projectId)
List flavor custom for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **familyKey** | **string**| The family key | 
  **platformKey** | **string**| The platform key | 
  **projectId** | **string**| The project id | 

### Return type

[**[]Flavor**](Flavor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFlavorCustomUsingGET**
> []Flavor ListFlavorCustomUsingGET(ctx, projectId)
List flavor custom for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**[]Flavor**](Flavor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFlavorForSimpleServerUsingGET**
> FlavorResponse ListFlavorForSimpleServerUsingGET(ctx, projectId)
List flavor for simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

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

