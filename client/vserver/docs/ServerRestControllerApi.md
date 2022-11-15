# \ServerRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerUsingPOST**](ServerRestControllerApi.md#CreateServerUsingPOST) | **Post** /v1/{project_id}/servers | Create server
[**DeleteServerUsingDELETE**](ServerRestControllerApi.md#DeleteServerUsingDELETE) | **Delete** /v1/{project_id}/servers | Delete Server
[**GetServerUsingGET**](ServerRestControllerApi.md#GetServerUsingGET) | **Get** /v1/{project_id}/servers/{server_id} | Get server by id
[**ListServerUsingGET**](ServerRestControllerApi.md#ListServerUsingGET) | **Get** /v1/{project_id}/servers | List server
[**RebootServerUsingPUT**](ServerRestControllerApi.md#RebootServerUsingPUT) | **Put** /v1/{project_id}/servers/reboot | Reboot server
[**ResizeServerUsingPUT**](ServerRestControllerApi.md#ResizeServerUsingPUT) | **Put** /v1/{project_id}/servers/resize | Change flavor of server
[**ServerLimitUsingGET**](ServerRestControllerApi.md#ServerLimitUsingGET) | **Get** /v1/{project_id}/servers/limit | Server limit
[**StartServerUsingPUT**](ServerRestControllerApi.md#StartServerUsingPUT) | **Put** /v1/{project_id}/servers/start | Start server
[**StopServerUsingPUT**](ServerRestControllerApi.md#StopServerUsingPUT) | **Put** /v1/{project_id}/servers/stop | Stop server
[**UpdateSecGroupServerUsingPUT**](ServerRestControllerApi.md#UpdateSecGroupServerUsingPUT) | **Put** /v1/{project_id}/servers/update_sec_group | Update SecGroups of server


# **CreateServerUsingPOST**
> ServerResponse CreateServerUsingPOST(ctx, createServerRequest, projectId)
Create server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createServerRequest** | [**CreateServerRequest**](CreateServerRequest.md)| createServerRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerUsingDELETE**
> ServerResponse DeleteServerUsingDELETE(ctx, deleteServerRequest, projectId)
Delete Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteServerRequest** | [**DeleteServerRequest**](DeleteServerRequest.md)| deleteServerRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerUsingGET**
> ServerResponse GetServerUsingGET(ctx, projectId, serverId)
Get server by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerUsingGET**
> ServerResponse ListServerUsingGET(ctx, projectId)
List server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootServerUsingPUT**
> ServerResponse RebootServerUsingPUT(ctx, projectId, updateServerRequest)
Reboot server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeServerUsingPUT**
> ServerResponse ResizeServerUsingPUT(ctx, projectId, resizeServerRequest)
Change flavor of server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **resizeServerRequest** | [**ResizeServerRequest**](ResizeServerRequest.md)| resizeServerRequest | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerLimitUsingGET**
> ServerLimit ServerLimitUsingGET(ctx, projectId)
Server limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**ServerLimit**](ServerLimit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartServerUsingPUT**
> ServerResponse StartServerUsingPUT(ctx, projectId, updateServerRequest)
Start server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopServerUsingPUT**
> ServerResponse StopServerUsingPUT(ctx, projectId, updateServerRequest)
Stop server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateServerRequest** | [**UpdateServerRequest**](UpdateServerRequest.md)| updateServerRequest | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecGroupServerUsingPUT**
> ServerResponse UpdateSecGroupServerUsingPUT(ctx, changeSecGroupRequest, projectId)
Update SecGroups of server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeSecGroupRequest** | [**ChangeSecGroupRequest**](ChangeSecGroupRequest.md)| changeSecGroupRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**ServerResponse**](ServerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

