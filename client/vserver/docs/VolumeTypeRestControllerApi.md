# \VolumeTypeRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVolumeTypeUsingGET**](VolumeTypeRestControllerApi.md#GetVolumeTypeUsingGET) | **Get** /v1/{project_id}/volume_types/{volume_type_id} | getVolumeType
[**ListVolumeTypeUsingGET**](VolumeTypeRestControllerApi.md#ListVolumeTypeUsingGET) | **Get** /v1/{project_id}/{volume_type_zone_id}/volume_types | listVolumeType


# **GetVolumeTypeUsingGET**
> VolumeTypeResponse GetVolumeTypeUsingGET(ctx, projectId, volumeTypeId)
getVolumeType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **volumeTypeId** | **string**| volume_type_id | 

### Return type

[**VolumeTypeResponse**](VolumeTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeTypeUsingGET**
> VolumeTypeResponse ListVolumeTypeUsingGET(ctx, projectId, volumeTypeZoneId)
listVolumeType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **volumeTypeZoneId** | **string**| volume_type_zone_id | 

### Return type

[**VolumeTypeResponse**](VolumeTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

