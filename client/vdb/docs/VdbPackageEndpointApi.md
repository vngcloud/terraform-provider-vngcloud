# \VdbPackageEndpointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListPackageByEngineTypeAndVersionUsingGET**](VdbPackageEndpointApi.md#GetListPackageByEngineTypeAndVersionUsingGET) | **Get** /v1/trv/{projectId}/package | getListPackageByEngineTypeAndVersion


# **GetListPackageByEngineTypeAndVersionUsingGET**
> PackageResponse GetListPackageByEngineTypeAndVersionUsingGET(ctx, engineType, engineVersion, projectId)
getListPackageByEngineTypeAndVersion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineType** | **string**| engineType | 
  **engineVersion** | **string**| engineVersion | 
  **projectId** | **string**| projectId | 

### Return type

[**PackageResponse**](PackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

