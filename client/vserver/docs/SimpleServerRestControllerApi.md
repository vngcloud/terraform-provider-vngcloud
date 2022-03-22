# \SimpleServerRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServerUsingDELETE1**](SimpleServerRestControllerApi.md#DeleteServerUsingDELETE1) | **Delete** /v1/{project_id}/simple-servers | Delete simple server
[**GetConsoleServerUsingGET**](SimpleServerRestControllerApi.md#GetConsoleServerUsingGET) | **Get** /v1/{project_id}/simple-servers/{server_id}/console | Get console simple server
[**GetServerUsingGET1**](SimpleServerRestControllerApi.md#GetServerUsingGET1) | **Get** /v1/{project_id}/simple-servers/{server_id} | Get simple server
[**ListServerUsingGET1**](SimpleServerRestControllerApi.md#ListServerUsingGET1) | **Get** /v1/{project_id}/simple-servers | List simple server
[**RebootServerUsingPUT1**](SimpleServerRestControllerApi.md#RebootServerUsingPUT1) | **Put** /v1/{project_id}/simple-servers/reboot | Reboot simple server
[**StartServerUsingPUT1**](SimpleServerRestControllerApi.md#StartServerUsingPUT1) | **Put** /v1/{project_id}/simple-servers/start | Start simple server
[**StopServerUsingPUT1**](SimpleServerRestControllerApi.md#StopServerUsingPUT1) | **Put** /v1/{project_id}/simple-servers/stop | Stop simple server


# **DeleteServerUsingDELETE1**
> DataResponseSimpleServer DeleteServerUsingDELETE1(ctx, deleteServerRequest, projectId)
Delete simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteServerRequest** | [**DeleteServerRequest**](DeleteServerRequest.md)| deleteServerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseSimpleServer**](DataResponseSimpleServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsoleServerUsingGET**
> DataResponsestring GetConsoleServerUsingGET(ctx, projectId, serverId)
Get console simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **serverId** | **string**| server id | 

### Return type

[**DataResponsestring**](DataResponsestring.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerUsingGET1**
> DataResponseSimpleServer GetServerUsingGET1(ctx, projectId, serverId)
Get simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **serverId** | **string**| server id | 

### Return type

[**DataResponseSimpleServer**](DataResponseSimpleServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerUsingGET1**
> PagingSimpleServer ListServerUsingGET1(ctx, projectId, optional)
List simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***SimpleServerRestControllerApiListServerUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimpleServerRestControllerApiListServerUsingGET1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingSimpleServer**](PagingSimpleServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootServerUsingPUT1**
> RebootServerUsingPUT1(ctx, projectId, updateServerRequest)
Reboot simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartServerUsingPUT1**
> StartServerUsingPUT1(ctx, projectId, updateServerRequest)
Start simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopServerUsingPUT1**
> StopServerUsingPUT1(ctx, projectId, updateServerRequest)
Stop simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

