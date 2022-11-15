# \AttachmentControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAttachmentByServerUsingGET**](AttachmentControllerApi.md#ListAttachmentByServerUsingGET) | **Get** /v1/{project_id}/multiVolumeAttachment/server/{server_uuid} | Get list attached instance of server 
[**ListAttachmentUsingGET**](AttachmentControllerApi.md#ListAttachmentUsingGET) | **Get** /v1/{project_id}/multiVolumeAttachment/{uuid} | Get list attached instance of multi attach volume


# **ListAttachmentByServerUsingGET**
> AttachmentResponse ListAttachmentByServerUsingGET(ctx, projectId, serverUuid)
Get list attached instance of server 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverUuid** | **string**| The server id | 

### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAttachmentUsingGET**
> AttachmentResponse ListAttachmentUsingGET(ctx, projectId, uuid)
Get list attached instance of multi attach volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The volume id | 

### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

