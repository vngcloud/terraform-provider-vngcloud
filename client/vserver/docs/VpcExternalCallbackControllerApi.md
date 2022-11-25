# \VpcExternalCallbackControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateStatusUsingPOST**](VpcExternalCallbackControllerApi.md#UpdateStatusUsingPOST) | **Post** /v2/external/callback | Update status in task table


# **UpdateStatusUsingPOST**
> CallbackResponse UpdateStatusUsingPOST(ctx, updateTaskStatusRequest)
Update status in task table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateTaskStatusRequest** | [**UpdateTaskStatusRequest**](UpdateTaskStatusRequest.md)| updateTaskStatusRequest | 

### Return type

[**CallbackResponse**](CallbackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

