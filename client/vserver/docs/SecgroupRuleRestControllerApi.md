# \SecgroupRuleRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecgroupRuleUsingPOST**](SecgroupRuleRestControllerApi.md#CreateSecgroupRuleUsingPOST) | **Post** /v1/{project_id}/secgroup-rules | Create security group rule
[**DeleteSecgroupRuleUsingDELETE**](SecgroupRuleRestControllerApi.md#DeleteSecgroupRuleUsingDELETE) | **Delete** /v1/{project_id}/secgroup-rules | Delete security group rule
[**GetSecgroupRuleUsingGET**](SecgroupRuleRestControllerApi.md#GetSecgroupRuleUsingGET) | **Get** /v1/{project_id}/secgroup-rules/{secgroup_rule_id} | getSecgroupRule
[**ListSecgroupRuleBySecgroupUsingGET**](SecgroupRuleRestControllerApi.md#ListSecgroupRuleBySecgroupUsingGET) | **Get** /v1/{project_id}/secgroup-rules/secgroup-ids/{secgroup_id} | listSecgroupRuleBySecgroup
[**UpdateUsingPUT**](SecgroupRuleRestControllerApi.md#UpdateUsingPUT) | **Put** /v1/{project_id}/secgroup-rules | Update security group rule


# **CreateSecgroupRuleUsingPOST**
> SecgroupRuleResponse CreateSecgroupRuleUsingPOST(ctx, createSecurityGroupRuleRequest, projectId)
Create security group rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSecurityGroupRuleRequest** | [**CreateSecurityGroupRuleRequest**](CreateSecurityGroupRuleRequest.md)| createSecurityGroupRuleRequest | 
  **projectId** | **string**| The project id | 

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
Delete security group rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteSecurityGroupRuleRequest** | [**DeleteSecurityGroupRuleRequest**](DeleteSecurityGroupRuleRequest.md)| deleteSecurityGroupRuleRequest | 
  **projectId** | **string**| The project id | 

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
  **projectId** | **string**| The project id | 
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
  **projectId** | **string**| The project id | 
  **secgroupId** | **string**| secgroup_id | 

### Return type

[**SecgroupRuleResponse**](SecgroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT**
> SecgroupRuleResponse UpdateUsingPUT(ctx, projectId, updateSecurityGroupRuleRequest)
Update security group rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateSecurityGroupRuleRequest** | [**UpdateSecurityGroupRuleRequest**](UpdateSecurityGroupRuleRequest.md)| updateSecurityGroupRuleRequest | 

### Return type

[**SecgroupRuleResponse**](SecgroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

