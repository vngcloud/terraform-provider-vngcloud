# \SshKeyRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHKeyUsingPOST**](SshKeyRestControllerApi.md#CreateSSHKeyUsingPOST) | **Post** /v1/{project_id}/ssh_keys | createSSHKey
[**DeleteSSHKeyUsingDELETE**](SshKeyRestControllerApi.md#DeleteSSHKeyUsingDELETE) | **Delete** /v1/{project_id}/ssh_keys | deleteSSHKey
[**GetSSHKeyUsingGET**](SshKeyRestControllerApi.md#GetSSHKeyUsingGET) | **Get** /v1/{project_id}/ssh_keys/{ssh_key_id} | getSSHKey
[**ImportSSHKeyUsingPOST**](SshKeyRestControllerApi.md#ImportSSHKeyUsingPOST) | **Post** /v1/{project_id}/ssh_keys/import | importSSHKey
[**ListSSHKeyUsingGET**](SshKeyRestControllerApi.md#ListSSHKeyUsingGET) | **Get** /v1/{project_id}/ssh_keys | listSSHKey


# **CreateSSHKeyUsingPOST**
> SshKeyResponse CreateSSHKeyUsingPOST(ctx, createSSHKeyRequest, projectId)
createSSHKey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSSHKeyRequest** | [**CreateSshKeyRequest**](CreateSshKeyRequest.md)| createSSHKeyRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSSHKeyUsingDELETE**
> SshKeyResponse DeleteSSHKeyUsingDELETE(ctx, deleteRequest, projectId)
deleteSSHKey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteRequest** | [**SdnSshKeyDeleteRequest**](SdnSshKeyDeleteRequest.md)| deleteRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSSHKeyUsingGET**
> SshKeyResponse GetSSHKeyUsingGET(ctx, projectId, sshKeyId)
getSSHKey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **sshKeyId** | **string**| ssh_key_id | 

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
importSSHKey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importSSHKeyRequest** | [**ImportSshKeyRequest**](ImportSshKeyRequest.md)| importSSHKeyRequest | 
  **projectId** | **string**| project_id | 

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
listSSHKey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**SshKeyResponse**](SSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

