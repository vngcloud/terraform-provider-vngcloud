# \ImageRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGPUImageUsingGET**](ImageRestControllerApi.md#ListGPUImageUsingGET) | **Get** /v1/{project_id}/images/gpu | listGPUImage
[**ListOSImageDefaultUsingGET**](ImageRestControllerApi.md#ListOSImageDefaultUsingGET) | **Get** /v1/{project_id}/images/os_default | listOSImageDefault
[**ListOSImageUsingGET**](ImageRestControllerApi.md#ListOSImageUsingGET) | **Get** /v1/{project_id}/images/os | listOSImage


# **ListGPUImageUsingGET**
> OsImageBaseResponse ListGPUImageUsingGET(ctx, projectId)
listGPUImage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**OsImageBaseResponse**](OSImageBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSImageDefaultUsingGET**
> string ListOSImageDefaultUsingGET(ctx, projectId)
listOSImageDefault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSImageUsingGET**
> OsImageBaseResponse ListOSImageUsingGET(ctx, projectId)
listOSImage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**OsImageBaseResponse**](OSImageBaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

