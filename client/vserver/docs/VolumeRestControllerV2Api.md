# \VolumeRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolumeUsingPUT1**](VolumeRestControllerV2Api.md#AttachVolumeUsingPUT1) | **Put** /v2/{projectId}/volumes/{volumeId}/servers/{serverId}/attach | Attach volume into the server
[**CreateVolumeUsingPOST1**](VolumeRestControllerV2Api.md#CreateVolumeUsingPOST1) | **Post** /v2/{projectId}/volumes | Create volume
[**DeleteVolumeUsingDELETE1**](VolumeRestControllerV2Api.md#DeleteVolumeUsingDELETE1) | **Delete** /v2/{projectId}/volumes/{volumeId} | Delete volume
[**DetachVolumeUsingPUT1**](VolumeRestControllerV2Api.md#DetachVolumeUsingPUT1) | **Put** /v2/{projectId}/volumes/{volumeId}/servers/{serverId}/detach | Detach volume into the server
[**GetBootVolumeByInstanceIdUsingGET1**](VolumeRestControllerV2Api.md#GetBootVolumeByInstanceIdUsingGET1) | **Get** /v2/{projectId}/volumes/servers/{serverId}/boot | Get boot volume by server id
[**GetVolumeByInstanceIdUsingGET1**](VolumeRestControllerV2Api.md#GetVolumeByInstanceIdUsingGET1) | **Get** /v2/{projectId}/volumes/servers/{serverId} | Get volume by server id
[**GetVolumeUsingGET2**](VolumeRestControllerV2Api.md#GetVolumeUsingGET2) | **Get** /v2/{projectId}/volumes/{volumeId} | Get volume by id
[**ListVolumeHistoryUsingGET**](VolumeRestControllerV2Api.md#ListVolumeHistoryUsingGET) | **Get** /v2/{projectId}/volumes/{volumeId}/history | List volume
[**ListVolumeUsingGET2**](VolumeRestControllerV2Api.md#ListVolumeUsingGET2) | **Get** /v2/{projectId}/volumes | List volume
[**ResizeVolumeUsingPUT1**](VolumeRestControllerV2Api.md#ResizeVolumeUsingPUT1) | **Put** /v2/{projectId}/volumes/{volumeId}/resize | Resize volume


# **AttachVolumeUsingPUT1**
> DataResponseVolume AttachVolumeUsingPUT1(ctx, attachVolumeRequest, projectId, serverId, volumeId)
Attach volume into the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachVolumeRequest** | [**AttachVolumeRequest**](AttachVolumeRequest.md)| attachVolumeRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 
  **volumeId** | **string**| The volume id | 

### Return type

[**DataResponseVolume**](DataResponse«Volume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVolumeUsingPOST1**
> DataResponse CreateVolumeUsingPOST1(ctx, createVolumeRequest, projectId)
Create volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createVolumeRequest** | [**CreateVolumeRequest**](CreateVolumeRequest.md)| createVolumeRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVolumeUsingDELETE1**
> DataResponseVolume DeleteVolumeUsingDELETE1(ctx, projectId, volumeId)
Delete volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **volumeId** | **string**| The volume id | 

### Return type

[**DataResponseVolume**](DataResponse«Volume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachVolumeUsingPUT1**
> DataResponseVolume DetachVolumeUsingPUT1(ctx, detachVolumeRequest, projectId, serverId, volumeId)
Detach volume into the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **detachVolumeRequest** | [**DetachVolumeRequest**](DetachVolumeRequest.md)| detachVolumeRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 
  **volumeId** | **string**| The volume id | 

### Return type

[**DataResponseVolume**](DataResponse«Volume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBootVolumeByInstanceIdUsingGET1**
> VolumeResponse GetBootVolumeByInstanceIdUsingGET1(ctx, projectId, serverId)
Get boot volume by server id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The volume id | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeByInstanceIdUsingGET1**
> VolumeResponse GetVolumeByInstanceIdUsingGET1(ctx, projectId, serverId)
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

# **GetVolumeUsingGET2**
> DataResponseVolume GetVolumeUsingGET2(ctx, projectId, volumeId)
Get volume by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **volumeId** | **string**| The volume id | 

### Return type

[**DataResponseVolume**](DataResponse«Volume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeHistoryUsingGET**
> []VolumeAction ListVolumeHistoryUsingGET(ctx, projectId, volumeId)
List volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **volumeId** | **string**| The volume id | 

### Return type

[**[]VolumeAction**](VolumeAction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeUsingGET2**
> PagingVolume ListVolumeUsingGET2(ctx, projectId, optional)
List volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***VolumeRestControllerV2ApiListVolumeUsingGET2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeRestControllerV2ApiListVolumeUsingGET2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingVolume**](Paging«Volume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeVolumeUsingPUT1**
> DataResponseVolume ResizeVolumeUsingPUT1(ctx, projectId, resizeVolumeRequest, volumeId)
Resize volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **resizeVolumeRequest** | [**ResizeVolumeRequest**](ResizeVolumeRequest.md)| resizeVolumeRequest | 
  **volumeId** | **string**| The volume id | 

### Return type

[**DataResponseVolume**](DataResponse«Volume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

