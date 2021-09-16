# \ProjectRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectUsingGET**](ProjectRestControllerApi.md#GetProjectUsingGET) | **Get** /v1/projects/{project_id} | getProject
[**ListProjectUsingGET**](ProjectRestControllerApi.md#ListProjectUsingGET) | **Get** /v1/projects | listProject


# **GetProjectUsingGET**
> ProjectResponse GetProjectUsingGET(ctx, projectId)
getProject

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjectUsingGET**
> ProjectResponse ListProjectUsingGET(ctx, )
listProject

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

