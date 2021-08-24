# \ServerGroupRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerGroupUsingPOST**](ServerGroupRestControllerApi.md#CreateServerGroupUsingPOST) | **Post** /v1/{project_id}/server-groups | Create server group
[**DeleteServerGroupUsingDELETE**](ServerGroupRestControllerApi.md#DeleteServerGroupUsingDELETE) | **Delete** /v1/{project_id}/server-groups | Delete server group
[**GetServerGroupUsingGET**](ServerGroupRestControllerApi.md#GetServerGroupUsingGET) | **Get** /v1/{project_id}/server-groups/{server_group_id} | Get server group
[**ListServerGroupPolicyUsingGET**](ServerGroupRestControllerApi.md#ListServerGroupPolicyUsingGET) | **Get** /v1/{project_id}/server-groups/policies | List server group policy
[**ListServerGroupUsingGET**](ServerGroupRestControllerApi.md#ListServerGroupUsingGET) | **Get** /v1/{project_id}/server-groups | List server group


# **CreateServerGroupUsingPOST**
> ServerGroupResponse CreateServerGroupUsingPOST(ctx, createServerGroupRequest, projectId)
Create server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createServerGroupRequest** | [**CreateServerGroupRequest**](CreateServerGroupRequest.md)| createServerGroupRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerGroupUsingDELETE**
> ServerGroupResponse DeleteServerGroupUsingDELETE(ctx, deleteServerRequest, projectId)
Delete server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteServerRequest** | [**DeleteServerGroupRequest**](DeleteServerGroupRequest.md)| deleteServerRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupUsingGET**
> ServerGroupResponse GetServerGroupUsingGET(ctx, projectId, serverGroupId)
Get server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverGroupId** | **string**| The server group id | 

### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerGroupPolicyUsingGET**
> ServerGroupPolicyResponse ListServerGroupPolicyUsingGET(ctx, projectId)
List server group policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**ServerGroupPolicyResponse**](ServerGroupPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerGroupUsingGET**
> ServerGroupResponse ListServerGroupUsingGET(ctx, projectId)
List server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

