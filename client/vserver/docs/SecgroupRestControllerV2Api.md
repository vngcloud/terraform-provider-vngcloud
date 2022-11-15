# \SecgroupRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecgroupUsingPOST1**](SecgroupRestControllerV2Api.md#CreateSecgroupUsingPOST1) | **Post** /v2/{projectId}/secgroups | Create security group
[**DeleteSecgroupUsingDELETE1**](SecgroupRestControllerV2Api.md#DeleteSecgroupUsingDELETE1) | **Delete** /v2/{projectId}/secgroups/{secgroupId} | Delete security group
[**GetAlllistSecgroupUsingGET**](SecgroupRestControllerV2Api.md#GetAlllistSecgroupUsingGET) | **Get** /v2/{projectId}/secgroups | List All SecurityGroup
[**GetListSecgroupRuleBySecgroupUsingGET**](SecgroupRestControllerV2Api.md#GetListSecgroupRuleBySecgroupUsingGET) | **Get** /v2/{projectId}/secgroups/{secgroupId}/secGroupRules | Security Group Rule By Security Group Id
[**GetListServerBySecgroupUsingGET**](SecgroupRestControllerV2Api.md#GetListServerBySecgroupUsingGET) | **Get** /v2/{projectId}/secgroups/{secgroupId}/servers | Infomation&#39;s Server of Security Group 
[**GetSecgroupUsingGET1**](SecgroupRestControllerV2Api.md#GetSecgroupUsingGET1) | **Get** /v2/{projectId}/secgroups/{secgroupId} | Security Group By Security Group Id
[**UpdateSecgroupUsingPUT1**](SecgroupRestControllerV2Api.md#UpdateSecgroupUsingPUT1) | **Put** /v2/{projectId}/secgroups/{secgroupId} | Update security group


# **CreateSecgroupUsingPOST1**
> DataResponseSecgroupEntity CreateSecgroupUsingPOST1(ctx, createSecurityGroupRequest, projectId)
Create security group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSecurityGroupRequest** | [**CreateSecurityGroupRequest**](CreateSecurityGroupRequest.md)| createSecurityGroupRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseSecgroupEntity**](DataResponse«SecgroupEntity».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecgroupUsingDELETE1**
> DeleteSecgroupUsingDELETE1(ctx, projectId, secgroupId)
Delete security group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlllistSecgroupUsingGET**
> PagingSecgroupDto GetAlllistSecgroupUsingGET(ctx, projectId, optional)
List All SecurityGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***SecgroupRestControllerV2ApiGetAlllistSecgroupUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecgroupRestControllerV2ApiGetAlllistSecgroupUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 

### Return type

[**PagingSecgroupDto**](Paging«SecgroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListSecgroupRuleBySecgroupUsingGET**
> DataResponseListSecgroupruleDto GetListSecgroupRuleBySecgroupUsingGET(ctx, projectId, secgroupId)
Security Group Rule By Security Group Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseListSecgroupruleDto**](DataResponse«List«SecgroupruleDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListServerBySecgroupUsingGET**
> DataResponseListServerOfSecgroupDto GetListServerBySecgroupUsingGET(ctx, projectId, secgroupId)
Infomation's Server of Security Group 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseListServerOfSecgroupDto**](DataResponse«List«ServerOfSecgroupDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecgroupUsingGET1**
> DataResponseSecgroupDto GetSecgroupUsingGET1(ctx, projectId, secgroupId)
Security Group By Security Group Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseSecgroupDto**](DataResponse«SecgroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecgroupUsingPUT1**
> DataResponseSecgroupDto UpdateSecgroupUsingPUT1(ctx, editSecurityGroupRequest, projectId, secgroupId)
Update security group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **editSecurityGroupRequest** | [**EditSecurityGroupRequest**](EditSecurityGroupRequest.md)| editSecurityGroupRequest | 
  **projectId** | **string**| The project id | 
  **secgroupId** | **string**| The secgroup id | 

### Return type

[**DataResponseSecgroupDto**](DataResponse«SecgroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

