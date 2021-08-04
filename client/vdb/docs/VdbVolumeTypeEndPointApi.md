# \VdbVolumeTypeEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListActiveVolumeTypeInZoneUsingGET**](VdbVolumeTypeEndPointApi.md#GetListActiveVolumeTypeInZoneUsingGET) | **Get** /v1/trv/{projectId}/volume-type/{volumeTypeZoneId} | getListActiveVolumeTypeInZone
[**GetListActiveVolumeTypeUsingGET**](VdbVolumeTypeEndPointApi.md#GetListActiveVolumeTypeUsingGET) | **Get** /v1/trv/{projectId}/volume-type | getListActiveVolumeType


# **GetListActiveVolumeTypeInZoneUsingGET**
> BaseResponse GetListActiveVolumeTypeInZoneUsingGET(ctx, projectId, volumeTypeZoneId)
getListActiveVolumeTypeInZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **volumeTypeZoneId** | **string**| volumeTypeZoneId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListActiveVolumeTypeUsingGET**
> VolumeTypeResponse GetListActiveVolumeTypeUsingGET(ctx, projectId)
getListActiveVolumeType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**VolumeTypeResponse**](VolumeTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

