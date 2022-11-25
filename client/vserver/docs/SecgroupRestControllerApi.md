# \SecgroupRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecgroupUsingPOST**](SecgroupRestControllerApi.md#CreateSecgroupUsingPOST) | **Post** /v1/{project_id}/secgroups | Create security group
[**DeleteSecgroupUsingDELETE**](SecgroupRestControllerApi.md#DeleteSecgroupUsingDELETE) | **Delete** /v1/{project_id}/secgroups | Delete security group
[**GetForSimpleServerUsingGET**](SecgroupRestControllerApi.md#GetForSimpleServerUsingGET) | **Get** /v1/{project_id}/secgroups/simple-servers/{server_id} | getForSimpleServer
[**GetSecgroupUsingGET**](SecgroupRestControllerApi.md#GetSecgroupUsingGET) | **Get** /v1/{project_id}/secgroups/{secgroup_id} | getSecgroup
[**ListSecgroupByInstanceUsingGET**](SecgroupRestControllerApi.md#ListSecgroupByInstanceUsingGET) | **Get** /v1/{project_id}/secgroups/vm-ids/{vm_id} | listSecgroupByInstance
[**ListSecgroupUsingGET**](SecgroupRestControllerApi.md#ListSecgroupUsingGET) | **Get** /v1/{project_id}/secgroups | listSecgroup
[**UpdateSecgroupUsingPUT**](SecgroupRestControllerApi.md#UpdateSecgroupUsingPUT) | **Put** /v1/{project_id}/secgroups | Update security group


# **CreateSecgroupUsingPOST**
> SecgroupResponse CreateSecgroupUsingPOST(ctx, createSecurityGroupRequest, projectId)
Create security group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSecurityGroupRequest** | [**CreateSecurityGroupRequest**](CreateSecurityGroupRequest.md)| createSecurityGroupRequest | 
  **projectId** | **string**| The project id | 

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
Delete security group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteSecurityGroupRequest** | [**DeleteSecurityGroupRequest**](DeleteSecurityGroupRequest.md)| deleteSecurityGroupRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForSimpleServerUsingGET**
> SecgroupResponse GetForSimpleServerUsingGET(ctx, projectId, serverId)
getForSimpleServer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **serverId** | **string**| server_id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecgroupUsingGET**
> SecgroupResponse GetSecgroupUsingGET(ctx, projectId, secgroupId)
getSecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
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
  **projectId** | **string**| The project id | 
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
  **projectId** | **string**| The project id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecgroupUsingPUT**
> SecgroupResponse UpdateSecgroupUsingPUT(ctx, editSecurityGroupRequest, projectId)
Update security group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **editSecurityGroupRequest** | [**EditSecurityGroupRequest**](EditSecurityGroupRequest.md)| editSecurityGroupRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**SecgroupResponse**](SecgroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

