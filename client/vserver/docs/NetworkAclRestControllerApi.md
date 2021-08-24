# \NetworkAclRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAclUsingPOST**](NetworkAclRestControllerApi.md#CreateNetworkAclUsingPOST) | **Post** /v1/{project_id}/policies | Create network-acl
[**DeleteNetworkAclUsingDELETE**](NetworkAclRestControllerApi.md#DeleteNetworkAclUsingDELETE) | **Delete** /v1/{project_id}/policies/{uuid} | Delete Network-acl
[**GetNetworkAclUsingGET**](NetworkAclRestControllerApi.md#GetNetworkAclUsingGET) | **Get** /v1/{project_id}/policies/{uuid} | Get network-acl by uuid
[**ListNetworkAclUsingGET**](NetworkAclRestControllerApi.md#ListNetworkAclUsingGET) | **Get** /v1/{project_id}/policies | List network-acl
[**UpdateAssociatedSubnetsUsingPUT**](NetworkAclRestControllerApi.md#UpdateAssociatedSubnetsUsingPUT) | **Put** /v1/{project_id}/policies/{uuid}/subnets | Update ACL Rule of Network-acl
[**UpdateRulesUsingPUT**](NetworkAclRestControllerApi.md#UpdateRulesUsingPUT) | **Put** /v1/{project_id}/policies/{uuid}/rules | Update Inbound/Outbound Rules of Network-acl


# **CreateNetworkAclUsingPOST**
> NetworkAclResponse CreateNetworkAclUsingPOST(ctx, createNetworkAclRequest, projectId)
Create network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createNetworkAclRequest** | [**CreateNetworkAclRequest**](CreateNetworkAclRequest.md)| createNetworkAclRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**NetworkAclResponse**](NetworkAclResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkAclUsingDELETE**
> BaseResponse DeleteNetworkAclUsingDELETE(ctx, projectId, uuid)
Delete Network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclUsingGET**
> NetworkAclResponse GetNetworkAclUsingGET(ctx, projectId, uuid)
Get network-acl by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**NetworkAclResponse**](NetworkAclResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkAclUsingGET**
> NetworkAclListResponse ListNetworkAclUsingGET(ctx, projectId)
List network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**NetworkAclListResponse**](NetworkAclListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAssociatedSubnetsUsingPUT**
> NetworkAclResponse UpdateAssociatedSubnetsUsingPUT(ctx, projectId, updateSubnetsRequest, uuid)
Update ACL Rule of Network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateSubnetsRequest** | [**UpdateNetworkAclSubnetsRequest**](UpdateNetworkAclSubnetsRequest.md)| updateSubnetsRequest | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**NetworkAclResponse**](NetworkAclResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRulesUsingPUT**
> NetworkAclResponse UpdateRulesUsingPUT(ctx, projectId, updateAclRulesRequest, uuid)
Update Inbound/Outbound Rules of Network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateAclRulesRequest** | [**UpdateNetworkAclRulesRequest**](UpdateNetworkAclRulesRequest.md)| updateAclRulesRequest | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**NetworkAclResponse**](NetworkAclResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

