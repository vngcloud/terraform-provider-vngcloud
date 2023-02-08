# \ProjectRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectUsingPOST**](ProjectRestControllerApi.md#CreateProjectUsingPOST) | **Post** /v1/projects | createProject
[**GetProjectUsingGET**](ProjectRestControllerApi.md#GetProjectUsingGET) | **Get** /v1/projects/{project_id} | getProject
[**ListAllProjectUsingGET**](ProjectRestControllerApi.md#ListAllProjectUsingGET) | **Get** /v1/projects/all | listAllProject
[**ListProjectUsingGET**](ProjectRestControllerApi.md#ListProjectUsingGET) | **Get** /v1/projects | listProject


# **CreateProjectUsingPOST**
> ProjectResponse CreateProjectUsingPOST(ctx, )
createProject

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **ListAllProjectUsingGET**
> []ProjectInfo ListAllProjectUsingGET(ctx, )
listAllProject

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectInfo**](ProjectInfo.md)

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

