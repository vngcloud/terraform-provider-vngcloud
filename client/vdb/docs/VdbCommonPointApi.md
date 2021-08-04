# \VdbCommonPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDbConfigMapUsingGET**](VdbCommonPointApi.md#GetDbConfigMapUsingGET) | **Get** /v1/trv/{projectId}/database-config-map | getDbConfigMap
[**GetDbStatusToActionUsingGET**](VdbCommonPointApi.md#GetDbStatusToActionUsingGET) | **Get** /v1/trv/{projectId}/database-status-to-actions | getDbStatusToAction


# **GetDbConfigMapUsingGET**
> BaseResponse GetDbConfigMapUsingGET(ctx, projectId)
getDbConfigMap

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDbStatusToActionUsingGET**
> BaseResponse GetDbStatusToActionUsingGET(ctx, projectId)
getDbStatusToAction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

