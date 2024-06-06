# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1WorkspaceGet**](V1WorkspaceControllerApi.md#V1WorkspaceGet) | **Get** /v1/workspace | 
[**V1WorkspacePost**](V1WorkspaceControllerApi.md#V1WorkspacePost) | **Post** /v1/workspace | 
[**V1WorkspaceResetServiceAccountPost**](V1WorkspaceControllerApi.md#V1WorkspaceResetServiceAccountPost) | **Post** /v1/workspace/reset-service-account | 

# **V1WorkspaceGet**
> WorkspaceDto V1WorkspaceGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1WorkspaceControllerApiV1WorkspaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1WorkspaceControllerApiV1WorkspaceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalUserId** | **optional.Int64**|  | 

### Return type

[**WorkspaceDto**](WorkspaceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WorkspacePost**
> WorkspaceDto V1WorkspacePost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1WorkspaceControllerApiV1WorkspacePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1WorkspaceControllerApiV1WorkspacePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalUserId** | **optional.Int64**|  | 

### Return type

[**WorkspaceDto**](WorkspaceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WorkspaceResetServiceAccountPost**
> WorkspaceDto V1WorkspaceResetServiceAccountPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1WorkspaceControllerApiV1WorkspaceResetServiceAccountPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1WorkspaceControllerApiV1WorkspaceResetServiceAccountPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalUserId** | **optional.Int64**|  | 

### Return type

[**WorkspaceDto**](WorkspaceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

