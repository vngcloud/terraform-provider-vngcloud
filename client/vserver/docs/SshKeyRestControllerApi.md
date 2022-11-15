# \SshKeyRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHKeyUsingPOST**](SshKeyRestControllerApi.md#CreateSSHKeyUsingPOST) | **Post** /v1/{project_id}/ssh_keys | Create SSH key
[**DeleteSSHKeyUsingDELETE**](SshKeyRestControllerApi.md#DeleteSSHKeyUsingDELETE) | **Delete** /v1/{project_id}/ssh_keys | Delete SSH key
[**GetSSHKeyUsingGET**](SshKeyRestControllerApi.md#GetSSHKeyUsingGET) | **Get** /v1/{project_id}/ssh_keys/{ssh_key_id} | Get SSH key
[**ImportSSHKeyUsingPOST**](SshKeyRestControllerApi.md#ImportSSHKeyUsingPOST) | **Post** /v1/{project_id}/ssh_keys/import | Import SSH key
[**ListSSHKeyUsingGET**](SshKeyRestControllerApi.md#ListSSHKeyUsingGET) | **Get** /v1/{project_id}/ssh_keys | List SSH key


# **CreateSSHKeyUsingPOST**
> SshKeyResponse CreateSSHKeyUsingPOST(ctx, createSSHKeyRequest, projectId)
Create SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSSHKeyRequest** | [**CreateSshKeyRequest**](CreateSshKeyRequest.md)| createSSHKeyRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSSHKeyUsingDELETE**
> BaseResponse DeleteSSHKeyUsingDELETE(ctx, deleteSSHKeyRequest, projectId)
Delete SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteSSHKeyRequest** | [**DeleteSshKeyRequest**](DeleteSshKeyRequest.md)| deleteSSHKeyRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSSHKeyUsingGET**
> SshKeyResponse GetSSHKeyUsingGET(ctx, projectId, sshKeyId)
Get SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **sshKeyId** | **string**| The ssh key id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSSHKeyUsingPOST**
> SshKeyResponse ImportSSHKeyUsingPOST(ctx, importSSHKeyRequest, projectId)
Import SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importSSHKeyRequest** | [**ImportSshKeyRequest**](ImportSshKeyRequest.md)| importSSHKeyRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSSHKeyUsingGET**
> SshKeyResponse ListSSHKeyUsingGET(ctx, projectId)
List SSH key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

