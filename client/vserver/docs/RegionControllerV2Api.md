# \RegionControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetALlRegionUsingGET**](RegionControllerV2Api.md#GetALlRegionUsingGET) | **Get** /v2/{projectId}/region | Get All Region
[**ListQuotaUsedUsingGET2**](RegionControllerV2Api.md#ListQuotaUsedUsingGET2) | **Get** /v2/{projectId}/region/{regionId}/users/validation | Check userId belong to Region


# **GetALlRegionUsingGET**
> DataResponseListRegionDto GetALlRegionUsingGET(ctx, projectId)
Get All Region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**DataResponseListRegionDto**](DataResponse«List«RegionDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotaUsedUsingGET2**
> DataResponseboolean ListQuotaUsedUsingGET2(ctx, projectId, regionId)
Check userId belong to Region

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **regionId** | **string**| regionId | 

### Return type

[**DataResponseboolean**](DataResponse«boolean».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

