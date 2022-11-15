# \SecgroupRuleRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecgroupRuleUsingPOST1**](SecgroupRuleRestControllerV2Api.md#CreateSecgroupRuleUsingPOST1) | **Post** /v2/{projectId}/secgroups/{secgroupId}/secgroupRules | Create security group rule
[**DeleteSecgroupRuleUsingDELETE1**](SecgroupRuleRestControllerV2Api.md#DeleteSecgroupRuleUsingDELETE1) | **Delete** /v2/{projectId}/secgroups/{secgroupId}/secgroupRules/{secgroupRuleId} | Delete security group rule
[**GetSecGroupRuleSamplesUsingGET**](SecgroupRuleRestControllerV2Api.md#GetSecGroupRuleSamplesUsingGET) | **Get** /v2/{projectId}/secgroups/{secgroupId}/secgroupRules/samples | List SecurityGroup Rule Sample
[**GetSecgroupRuleUsingGET1**](SecgroupRuleRestControllerV2Api.md#GetSecgroupRuleUsingGET1) | **Get** /v2/{projectId}/secgroups/{secgroupId}/secgroupRules/{secgroupRuleId} | Security Group By Security Group Id
[**UpdateUsingPUT1**](SecgroupRuleRestControllerV2Api.md#UpdateUsingPUT1) | **Put** /v2/{projectId}/secgroups/{secgroupId}/secgroupRules/{secgroupRuleId} | Update security group rule


# **CreateSecgroupRuleUsingPOST1**
> DataResponseSecgroupRuleEntity CreateSecgroupRuleUsingPOST1(ctx, createSecurityGroupRuleRequest, secgroupId, projectId)
Create security group rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSecurityGroupRuleRequest** | [**CreateSecurityGroupRuleRequest**](CreateSecurityGroupRuleRequest.md)| createSecurityGroupRuleRequest | 
  **secgroupId** | **string**| The secgroup id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseSecgroupRuleEntity**](DataResponse«SecgroupRuleEntity».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecgroupRuleUsingDELETE1**
> DeleteSecgroupRuleUsingDELETE1(ctx, projectId, secgroupRuleId, secgroupId)
Delete security group rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupRuleId** | **string**| The secgroup rule id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecGroupRuleSamplesUsingGET**
> DataResponseListSecurityGroupRulesSampleEntity GetSecGroupRuleSamplesUsingGET(ctx, projectId, secgroupId)
List SecurityGroup Rule Sample

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseListSecurityGroupRulesSampleEntity**](DataResponse«List«SecurityGroupRulesSampleEntity»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecgroupRuleUsingGET1**
> DataResponseListSecgroupruleDto GetSecgroupRuleUsingGET1(ctx, projectId, secgroupRuleId, secgroupId)
Security Group By Security Group Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupRuleId** | **string**| secgroupRuleId | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseListSecgroupruleDto**](DataResponse«List«SecgroupruleDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT1**
> DataResponseSecgroupruleDto UpdateUsingPUT1(ctx, projectId, secgroupRuleId, updateSecurityGroupRuleRequest, secgroupId)
Update security group rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupRuleId** | **string**| The secgroup rule id | 
  **updateSecurityGroupRuleRequest** | [**UpdateSecurityGroupRuleRequest**](UpdateSecurityGroupRuleRequest.md)| updateSecurityGroupRuleRequest | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseSecgroupruleDto**](DataResponse«SecgroupruleDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

