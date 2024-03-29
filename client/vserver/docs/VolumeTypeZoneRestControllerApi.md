# \VolumeTypeZoneRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVolumeTypeZoneUsingGET**](VolumeTypeZoneRestControllerApi.md#GetVolumeTypeZoneUsingGET) | **Get** /v1/{project_id}/volume_type_zones/{volume_type_zone_id} | getVolumeTypeZone
[**ListVolumeTypeZoneUsingGET**](VolumeTypeZoneRestControllerApi.md#ListVolumeTypeZoneUsingGET) | **Get** /v1/{project_id}/volume_type_zones | listVolumeTypeZone


# **GetVolumeTypeZoneUsingGET**
> VolumeTypeZoneResponse GetVolumeTypeZoneUsingGET(ctx, projectId, volumeTypeZoneId)
getVolumeTypeZone

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

