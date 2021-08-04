# \SecgroupRuleRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecgroupRuleUsingPOST**](SecgroupRuleRestControllerApi.md#CreateSecgroupRuleUsingPOST) | **Post** /v1/{project_id}/secgroup-rules | createSecgroupRule
[**DeleteSecgroupRuleUsingDELETE**](SecgroupRuleRestControllerApi.md#DeleteSecgroupRuleUsingDELETE) | **Delete** /v1/{project_id}/secgroup-rules | deleteSecgroupRule
[**GetSecgroupRuleUsingGET**](SecgroupRuleRestControllerApi.md#GetSecgroupRuleUsingGET) | **Get** /v1/{project_id}/secgroup-rules/{secgroup_rule_id} | getSecgroupRule
[**ListSecgroupRuleBySecgroupUsingGET**](SecgroupRuleRestControllerApi.md#ListSecgroupRuleBySecgroupUsingGET) | **Get** /v1/{project_id}/secgroup-rules/secgroup-ids/{secgroup_id} | listSecgroupRuleBySecgroup


# **CreateSecgroupRuleUsingPOST**
> SecgroupRuleResponse CreateSecgroupRuleUsingPOST(ctx, createSecurityGroupRuleRequest, projectId)
createSecgroupRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSecurityGroupRuleRequest** | [**CreateSecurityGroupRuleRequest**](CreateSecurityGroupRuleRequest.md)| createSecurityGroupRuleRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**SecgroupRuleResponse**](SecgroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecgroupRuleUsingDELETE**
> BaseResponse DeleteSecgroupRuleUsingDELETE(ctx, deleteSecurityGroupRuleRequest, projectId)
deleteSecgroupRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteSecurityGroupRuleRequest** | [**DeleteSecurityGroupRuleRequest**](DeleteSecurityGroupRuleRequest.md)| deleteSecurityGroupRuleRequest | 
  **projectId** | **string**| project_id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecgroupRuleUsingGET**
> SecgroupRuleResponse GetSecgroupRuleUsingGET(ctx, projectId, secgroupRuleId)
getSecgroupRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **secgroupRuleId** | **string**| secgroup_rule_id | 

### Return type

[**SecgroupRuleResponse**](SecgroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecgroupRuleBySecgroupUsingGET**
> SecgroupRuleResponse ListSecgroupRuleBySecgroupUsingGET(ctx, projectId, secgroupId)
listSecgroupRuleBySecgroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **secgroupId** | **string**| secgroup_id | 

### Return type

[**SecgroupRuleResponse**](SecgroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

