# \VolumeRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolumeUsingPUT**](VolumeRestControllerApi.md#AttachVolumeUsingPUT) | **Put** /v1/{project_id}/volumes/attach | Attach volume into the server
[**CreateVolumeUsingPOST**](VolumeRestControllerApi.md#CreateVolumeUsingPOST) | **Post** /v1/{project_id}/volumes | Create volume
[**DeleteVolumeUsingDELETE**](VolumeRestControllerApi.md#DeleteVolumeUsingDELETE) | **Delete** /v1/{project_id}/volumes | Delete volume
[**DetachVolumeUsingPUT**](VolumeRestControllerApi.md#DetachVolumeUsingPUT) | **Put** /v1/{project_id}/volumes/detach | Detach volume into the server
[**GetVolumeUsingGET**](VolumeRestControllerApi.md#GetVolumeUsingGET) | **Get** /v1/{project_id}/volumes/{volume_id} | Get volume by id
[**ListVolumeUsingGET**](VolumeRestControllerApi.md#ListVolumeUsingGET) | **Get** /v1/{project_id}/volumes | List volume
[**ResizeVolumeUsingPUT**](VolumeRestControllerApi.md#ResizeVolumeUsingPUT) | **Put** /v1/{project_id}/volumes/resize | Resize volume


# **AttachVolumeUsingPUT**
> VolumeResponse AttachVolumeUsingPUT(ctx, attachVolumeRequest, projectId)
Attach volume into the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachVolumeRequest** | [**AttachVolumeRequest**](AttachVolumeRequest.md)| attachVolumeRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVolumeUsingPOST**
> VolumeResponse CreateVolumeUsingPOST(ctx, createVolumeRequest, projectId)
Create volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createVolumeRequest** | [**CreateVolumeRequest**](CreateVolumeRequest.md)| createVolumeRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVolumeUsingDELETE**
> VolumeResponse DeleteVolumeUsingDELETE(ctx, deleteVolumeRequest, projectId)
Delete volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteVolumeRequest** | [**DeleteVolumeRequest**](DeleteVolumeRequest.md)| deleteVolumeRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachVolumeUsingPUT**
> VolumeResponse DetachVolumeUsingPUT(ctx, detachVolumeRequest, projectId)
Detach volume into the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **detachVolumeRequest** | [**DetachVolumeRequest**](DetachVolumeRequest.md)| detachVolumeRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeUsingGET**
> VolumeResponse GetVolumeUsingGET(ctx, projectId, volumeId)
Get volume by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **volumeId** | **string**| The project id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeUsingGET**
> VolumeResponse ListVolumeUsingGET(ctx, projectId)
List volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeVolumeUsingPUT**
> VolumeResponse ResizeVolumeUsingPUT(ctx, projectId, resizeVolumeRequest)
Resize volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **resizeVolumeRequest** | [**ResizeVolumeRequest**](ResizeVolumeRequest.md)| resizeVolumeRequest | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

