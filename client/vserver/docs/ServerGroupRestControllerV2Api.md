# \ServerGroupRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerGroupUsingPOST1**](ServerGroupRestControllerV2Api.md#CreateServerGroupUsingPOST1) | **Post** /v2/{projectId}/serverGroups | Create server group
[**DeleteServerGroupUsingDELETE1**](ServerGroupRestControllerV2Api.md#DeleteServerGroupUsingDELETE1) | **Delete** /v2/{projectId}/serverGroups/{serverGroupId} | Delete server group
[**GetListServerGroupWithPagingUsingGET**](ServerGroupRestControllerV2Api.md#GetListServerGroupWithPagingUsingGET) | **Get** /v2/{projectId}/serverGroups | List server group with Paging
[**GetServerGroupUsingGET1**](ServerGroupRestControllerV2Api.md#GetServerGroupUsingGET1) | **Get** /v2/{projectId}/serverGroups/{serverGroupId} | Get server group
[**ListServerGroupPolicyUsingGET1**](ServerGroupRestControllerV2Api.md#ListServerGroupPolicyUsingGET1) | **Get** /v2/{projectId}/serverGroups/policies | List server group policy
[**UpdateServerGroupUsingPUT1**](ServerGroupRestControllerV2Api.md#UpdateServerGroupUsingPUT1) | **Put** /v2/{projectId}/serverGroups/{serverGroupId} | Update server group


# **CreateServerGroupUsingPOST1**
> DataResponseServerGroupDto CreateServerGroupUsingPOST1(ctx, createServerGroupRequest, projectId)
Create server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createServerGroupRequest** | [**CreateServerGroupRequest**](CreateServerGroupRequest.md)| createServerGroupRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseServerGroupDto**](DataResponse«ServerGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerGroupUsingDELETE1**
> DeleteServerGroupUsingDELETE1(ctx, projectId, serverGroupId)
Delete server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverGroupId** | **string**| The server group id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListServerGroupWithPagingUsingGET**
> PagingServerGroupDetail GetListServerGroupWithPagingUsingGET(ctx, name, page, projectId, size)
List server group with Paging

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
  **projectId** | **string**| project id | 
  **size** | **string**| size | [default to 10]

### Return type

[**PagingServerGroupDetail**](Paging«ServerGroupDetail».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupUsingGET1**
> DataResponseServerGroupDto GetServerGroupUsingGET1(ctx, projectId, serverGroupId)
Get server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverGroupId** | **string**| The server group id | 

### Return type

[**DataResponseServerGroupDto**](DataResponse«ServerGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerGroupPolicyUsingGET1**
> DataResponseListServerGroupPolicy ListServerGroupPolicyUsingGET1(ctx, projectId)
List server group policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListServerGroupPolicy**](DataResponse«List«ServerGroupPolicy»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServerGroupUsingPUT1**
> DataResponseServerGroupDto UpdateServerGroupUsingPUT1(ctx, projectId, serverGroupId, updateServerGroupRequest)
Update server group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverGroupId** | **string**| The server group id | 
  **updateServerGroupRequest** | [**UpdateServerGroupRequestV2**](UpdateServerGroupRequestV2.md)| updateServerGroupRequest | 

### Return type

[**DataResponseServerGroupDto**](DataResponse«ServerGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

