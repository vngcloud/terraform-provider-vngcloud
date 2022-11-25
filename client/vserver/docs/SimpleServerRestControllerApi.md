# \SimpleServerRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServerUsingDELETE2**](SimpleServerRestControllerApi.md#DeleteServerUsingDELETE2) | **Delete** /v1/{project_id}/simple-servers | Delete simple server
[**GetConsoleServerUsingGET**](SimpleServerRestControllerApi.md#GetConsoleServerUsingGET) | **Get** /v1/{project_id}/simple-servers/{server_id}/console | Get console simple server
[**GetServerUsingGET2**](SimpleServerRestControllerApi.md#GetServerUsingGET2) | **Get** /v1/{project_id}/simple-servers/{server_id} | Get simple server
[**ListServerUsingGET2**](SimpleServerRestControllerApi.md#ListServerUsingGET2) | **Get** /v1/{project_id}/simple-servers | List simple server
[**RebootServerUsingPUT2**](SimpleServerRestControllerApi.md#RebootServerUsingPUT2) | **Put** /v1/{project_id}/simple-servers/reboot | Reboot simple server
[**StartServerUsingPUT2**](SimpleServerRestControllerApi.md#StartServerUsingPUT2) | **Put** /v1/{project_id}/simple-servers/start | Start simple server
[**StopServerUsingPUT2**](SimpleServerRestControllerApi.md#StopServerUsingPUT2) | **Put** /v1/{project_id}/simple-servers/stop | Stop simple server


# **DeleteServerUsingDELETE2**
> DataResponseSimpleServer DeleteServerUsingDELETE2(ctx, deleteServerRequest, projectId)
Delete simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteServerRequest** | [**DeleteServerRequest**](DeleteServerRequest.md)| deleteServerRequest | 
  **projectId** | **string**| project id | 

### Return type

[**DataResponseSimpleServer**](DataResponse«SimpleServer».md)

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

[**DataResponsestring**](DataResponse«string».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerUsingGET2**
> DataResponseSimpleServer GetServerUsingGET2(ctx, projectId, serverId)
Get simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **serverId** | **string**| server id | 

### Return type

[**DataResponseSimpleServer**](DataResponse«SimpleServer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerUsingGET2**
> PagingSimpleServer ListServerUsingGET2(ctx, projectId, optional)
List simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***SimpleServerRestControllerApiListServerUsingGET2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SimpleServerRestControllerApiListServerUsingGET2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingSimpleServer**](Paging«SimpleServer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootServerUsingPUT2**
> DataResponseSimpleServer RebootServerUsingPUT2(ctx, projectId, updateServerRequest)
Reboot simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

[**DataResponseSimpleServer**](DataResponse«SimpleServer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartServerUsingPUT2**
> DataResponseSimpleServer StartServerUsingPUT2(ctx, projectId, updateServerRequest)
Start simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

[**DataResponseSimpleServer**](DataResponse«SimpleServer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopServerUsingPUT2**
> DataResponseSimpleServer StopServerUsingPUT2(ctx, projectId, updateServerRequest)
Stop simple server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

[**DataResponseSimpleServer**](DataResponse«SimpleServer».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

