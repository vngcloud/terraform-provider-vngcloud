# \VolumeTypeZoneRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlavorZoneUsingGET1**](VolumeTypeZoneRestControllerApi.md#GetFlavorZoneUsingGET1) | **Get** /v1/{project_id}/volume_type_zones/{volume_type_zone_id} | getFlavorZone
[**ListVolumeTypeZoneUsingGET**](VolumeTypeZoneRestControllerApi.md#ListVolumeTypeZoneUsingGET) | **Get** /v1/{project_id}/volume_type_zones | listVolumeTypeZone


# **GetFlavorZoneUsingGET1**
> VolumeTypeZoneResponse GetFlavorZoneUsingGET1(ctx, projectId, volumeTypeZoneId)
getFlavorZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **volumeTypeZoneId** | **string**| volume_type_zone_id | 

### Return type

[**VolumeTypeZoneResponse**](VolumeTypeZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVolumeTypeZoneUsingGET**
> VolumeTypeZoneResponse ListVolumeTypeZoneUsingGET(ctx, projectId)
listVolumeTypeZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**VolumeTypeZoneResponse**](VolumeTypeZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

