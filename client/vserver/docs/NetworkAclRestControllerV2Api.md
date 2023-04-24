# \NetworkAclRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAclUsingPOST1**](NetworkAclRestControllerV2Api.md#CreateNetworkAclUsingPOST1) | **Post** /v2/{projectId}/policies | Create Network-acl
[**DeleteNetworkAclUsingDELETE1**](NetworkAclRestControllerV2Api.md#DeleteNetworkAclUsingDELETE1) | **Delete** /v2/{projectId}/policies/{networkAclUuid} | Delete Network-acl
[**GetNetworkAclPolicyRulesUsingGET**](NetworkAclRestControllerV2Api.md#GetNetworkAclPolicyRulesUsingGET) | **Get** /v2/{projectId}/policies/{uuid}/rules | Get Network-Acl Policy rules (inbound/outbound) by uuid
[**GetNetworkAclUsingGET1**](NetworkAclRestControllerV2Api.md#GetNetworkAclUsingGET1) | **Get** /v2/{projectId}/policies/{uuid} | Get Detail Network-Acl by uuid
[**GetNetworkAclsUsingGET**](NetworkAclRestControllerV2Api.md#GetNetworkAclsUsingGET) | **Get** /v2/{projectId}/policies/list | List network-acl by project id and network acl uuid
[**UpdateAssociatedSubnetsUsingPUT1**](NetworkAclRestControllerV2Api.md#UpdateAssociatedSubnetsUsingPUT1) | **Put** /v2/{projectId}/policies/{uuid}/subnets | Update associated Subnets ACL of Network-Acl
[**UpdateRulesUsingPUT1**](NetworkAclRestControllerV2Api.md#UpdateRulesUsingPUT1) | **Put** /v2/{projectId}/policies/{uuid}/rules | Update Inbound/Outbound Rules of Network-acl


# **CreateNetworkAclUsingPOST1**
> DataResponseNetworkAclDto CreateNetworkAclUsingPOST1(ctx, createNetworkAclRequest, projectId)
Create Network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createNetworkAclRequest** | [**CreateNetworkAclRequest**](CreateNetworkAclRequest.md)| createNetworkAclRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseNetworkAclDto**](DataResponse«NetworkAclDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkAclUsingDELETE1**
> DeleteNetworkAclUsingDELETE1(ctx, networkAclUuid, projectId)
Delete Network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkAclUuid** | **string**| The network-acl Uuid | 
  **projectId** | **string**| The project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclPolicyRulesUsingGET**
> DataResponseListNetworkAclPolicyRuleDto GetNetworkAclPolicyRulesUsingGET(ctx, projectId, uuid)
Get Network-Acl Policy rules (inbound/outbound) by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**DataResponseListNetworkAclPolicyRuleDto**](DataResponse«List«NetworkAclPolicyRuleDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclUsingGET1**
> DataResponseNetworkAclDetailDto GetNetworkAclUsingGET1(ctx, projectId, uuid)
Get Detail Network-Acl by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**DataResponseNetworkAclDetailDto**](DataResponse«NetworkAclDetailDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkAclsUsingGET**
> PagingNetworkAclDto GetNetworkAclsUsingGET(ctx, name, page, projectId, size)
List network-acl by project id and network acl uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
  **projectId** | **string**| The project id | 
  **size** | **string**| size | [default to 10]

### Return type

[**PagingNetworkAclDto**](Paging«NetworkAclDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAssociatedSubnetsUsingPUT1**
> DataResponseNetworkAclDto UpdateAssociatedSubnetsUsingPUT1(ctx, projectId, updateSubnetsRequest, uuid)
Update associated Subnets ACL of Network-Acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateSubnetsRequest** | [**UpdateNetworkAclSubnetsRequest**](UpdateNetworkAclSubnetsRequest.md)| updateSubnetsRequest | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**DataResponseNetworkAclDto**](DataResponse«NetworkAclDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRulesUsingPUT1**
> DataResponseNetworkAclModel UpdateRulesUsingPUT1(ctx, projectId, updateAclRulesRequest, uuid)
Update Inbound/Outbound Rules of Network-acl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateAclRulesRequest** | [**UpdateNetworkAclRulesRequest**](UpdateNetworkAclRulesRequest.md)| updateAclRulesRequest | 
  **uuid** | **string**| The network-acl uuid | 

### Return type

[**DataResponseNetworkAclModel**](DataResponse«NetworkAclModel».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

