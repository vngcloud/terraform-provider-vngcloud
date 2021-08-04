# \SecgroupRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecgroupUsingPOST**](SecgroupRestControllerApi.md#CreateSecgroupUsingPOST) | **Post** /v1/{project_id}/secgroups | createSecgroup
[**DeleteSecgroupUsingDELETE**](SecgroupRestControllerApi.md#DeleteSecgroupUsingDELETE) | **Delete** /v1/{project_id}/secgroups | deleteSecgroup
[**EditSecgroupUsingPUT**](SecgroupRestControllerApi.md#EditSecgroupUsingPUT) | **Put** /v1/{project_id}/secgroups | editSecgroup
[**GetSecgroupUsingGET**](SecgroupRestControllerApi.md#GetSecgroupUsingGET) | **Get** /v1/{project_id}/secgroups/{secgroup_id} | getSecgroup
[**ListSecgroupByInstanceUsingGET**](SecgroupRestControllerApi.md#ListSecgroupByInstanceUsingGET) | **Get** /v1/{project_id}/secgroups/vm-ids/{vm_id} | listSecgroupByInstance
[**ListSecgroupUsingGET**](SecgroupRestControllerApi.md#ListSecgroupUsingGET) | **Get** /v1/{project_id}/secgroups | listSecgroup


# **CreateSecgroupUsingPOST**
> SecgroupResponse CreateSecgroupUsingPOST(ctx, createSecurityGroupRequest, projectId)
createSecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSecurityGroupRequest** | [**CreateSecurityGroupRequest**](CreateSecurityGroupRequest.md)| createSecurityGroupRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecgroupUsingDELETE**
> BaseResponse DeleteSecgroupUsingDELETE(ctx, deleteSecurityGroupRequest, projectId)
deleteSecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteSecurityGroupRequest** | [**DeleteSecurityGroupRequest**](DeleteSecurityGroupRequest.md)| deleteSecurityGroupRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditSecgroupUsingPUT**
> SecgroupResponse EditSecgroupUsingPUT(ctx, editSecurityGroupRequest, projectId)
editSecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **editSecurityGroupRequest** | [**EditSecurityGroupRequest**](EditSecurityGroupRequest.md)| editSecurityGroupRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecgroupUsingGET**
> SecgroupResponse GetSecgroupUsingGET(ctx, projectId, secgroupId)
getSecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **secgroupId** | **string**| secgroup_id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecgroupByInstanceUsingGET**
> SecgroupResponse ListSecgroupByInstanceUsingGET(ctx, projectId, vmId)
listSecgroupByInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **vmId** | **string**| vm_id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecgroupUsingGET**
> SecgroupResponse ListSecgroupUsingGET(ctx, projectId)
listSecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

