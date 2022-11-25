# \SimpleVolumeRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVolumeUsingGET**](SimpleVolumeRestControllerApi.md#GetVolumeUsingGET) | **Get** /v1/{project_id}/simple-volumes/{volume_id} | Get simple volume
[**ListVolumeUsingGET**](SimpleVolumeRestControllerApi.md#ListVolumeUsingGET) | **Get** /v1/{project_id}/simple-volumes | List simple volume


# **GetVolumeUsingGET**
> DataResponseSimpleServer GetVolumeUsingGET(ctx, projectId, volumeId)
Get simple volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **volumeId** | **string**| volume id | 

### Return type

[**DataResponseSimpleServer**](DataResponse«SimpleServer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeUsingGET**
> PagingSimpleVolume ListVolumeUsingGET(ctx, projectId, optional)
List simple volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***SimpleVolumeRestControllerApiListVolumeUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimpleVolumeRestControllerApiListVolumeUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingSimpleVolume**](Paging«SimpleVolume».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

