# \SshKeyRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHKeyUsingPOST1**](SshKeyRestControllerV2Api.md#CreateSSHKeyUsingPOST1) | **Post** /v2/{projectId}/sshKeys | Create SSH key
[**DeleteSSHKeyUsingDELETE1**](SshKeyRestControllerV2Api.md#DeleteSSHKeyUsingDELETE1) | **Delete** /v2/{projectId}/sshKeys/{sshKeyId} | Delete SSH key
[**GetListSecgroupWithPagingUsingGET1**](SshKeyRestControllerV2Api.md#GetListSecgroupWithPagingUsingGET1) | **Get** /v2/{projectId}/sshKeys | List  SSH key Paging
[**GetSSHKeyUsingGET1**](SshKeyRestControllerV2Api.md#GetSSHKeyUsingGET1) | **Get** /v2/{projectId}/sshKeys/{sshKeyId} | Get SSH key
[**ImportSSHKeyUsingPOST1**](SshKeyRestControllerV2Api.md#ImportSSHKeyUsingPOST1) | **Post** /v2/{projectId}/sshKeys/import | Import SSH key


# **CreateSSHKeyUsingPOST1**
> DataResponseSshKeyDto CreateSSHKeyUsingPOST1(ctx, createSSHKeyRequest, projectId)
Create SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSSHKeyRequest** | [**CreateSshKeyRequest**](CreateSshKeyRequest.md)| createSSHKeyRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseSshKeyDto**](DataResponse«SSHKeyDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSSHKeyUsingDELETE1**
> DeleteSSHKeyUsingDELETE1(ctx, projectId, sshKeyId)
Delete SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **sshKeyId** | **string**| The ssh-key id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListSecgroupWithPagingUsingGET1**
> PagingSshKeyDto GetListSecgroupWithPagingUsingGET1(ctx, name, page, projectId, size)
List  SSH key Paging

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
  **projectId** | **string**| project id | 
  **size** | **string**| size | [default to 10]

### Return type

[**PagingSshKeyDto**](Paging«SSHKeyDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSSHKeyUsingGET1**
> DataResponseSshKeyDto GetSSHKeyUsingGET1(ctx, projectId, sshKeyId)
Get SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **sshKeyId** | **string**| The ssh key id | 

### Return type

[**DataResponseSshKeyDto**](DataResponse«SSHKeyDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSSHKeyUsingPOST1**
> DataResponseSshKeyDto ImportSSHKeyUsingPOST1(ctx, importSSHKeyRequest, projectId)
Import SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importSSHKeyRequest** | [**ImportSshKeyRequest**](ImportSshKeyRequest.md)| importSSHKeyRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseSshKeyDto**](DataResponse«SSHKeyDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

