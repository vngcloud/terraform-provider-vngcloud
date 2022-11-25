# \VolumeRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolumeUsingPUT**](VolumeRestControllerApi.md#AttachVolumeUsingPUT) | **Put** /v1/{project_id}/volumes/attach | Attach volume into the server
[**CreateVolumeUsingPOST**](VolumeRestControllerApi.md#CreateVolumeUsingPOST) | **Post** /v1/{project_id}/volumes | Create volume
[**DeleteVolumeUsingDELETE**](VolumeRestControllerApi.md#DeleteVolumeUsingDELETE) | **Delete** /v1/{project_id}/volumes | Delete volume
[**DetachVolumeUsingPUT**](VolumeRestControllerApi.md#DetachVolumeUsingPUT) | **Put** /v1/{project_id}/volumes/detach | Detach volume into the server
[**GetBootVolumeByInstanceIdUsingGET**](VolumeRestControllerApi.md#GetBootVolumeByInstanceIdUsingGET) | **Get** /v1/{project_id}/volumes/servers/{server_id}/boot | Get boot volume by server id
[**GetVolumeByInstanceIdUsingGET**](VolumeRestControllerApi.md#GetVolumeByInstanceIdUsingGET) | **Get** /v1/{project_id}/volumes/servers/{server_id} | Get volume by server id
[**GetVolumeUsingGET1**](VolumeRestControllerApi.md#GetVolumeUsingGET1) | **Get** /v1/{project_id}/volumes/{volume_id} | Get volume by id
[**ListEncryptionUsingGET**](VolumeRestControllerApi.md#ListEncryptionUsingGET) | **Get** /v1/{project_id}/volumes/encryption_types | Get list encryption type
[**ListVolumeUsingGET1**](VolumeRestControllerApi.md#ListVolumeUsingGET1) | **Get** /v1/{project_id}/volumes | List volume
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

# **GetBootVolumeByInstanceIdUsingGET**
> VolumeResponse GetBootVolumeByInstanceIdUsingGET(ctx, projectId, serverId)
Get boot volume by server id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeByInstanceIdUsingGET**
> VolumeResponse GetVolumeByInstanceIdUsingGET(ctx, projectId, serverId)
Get volume by server id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeUsingGET1**
> VolumeResponse GetVolumeUsingGET1(ctx, projectId, volumeId)
Get volume by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **volumeId** | **string**| The volume id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEncryptionUsingGET**
> []EncryptionType ListEncryptionUsingGET(ctx, projectId)
Get list encryption type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**[]EncryptionType**](EncryptionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeUsingGET1**
> VolumeResponse ListVolumeUsingGET1(ctx, projectId)
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

