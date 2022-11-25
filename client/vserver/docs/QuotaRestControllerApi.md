# \QuotaRestControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListByUserUsingGET**](QuotaRestControllerApi.md#ListByUserUsingGET) | **Get** /v1/{project_id}/quota/user/{user_id} | listByUser
[**ListQuotaUsedUsingGET**](QuotaRestControllerApi.md#ListQuotaUsedUsingGET) | **Get** /v1/{project_id}/quota-used | listQuotaUsed
[**ListUsingGET**](QuotaRestControllerApi.md#ListUsingGET) | **Get** /v1/{project_id}/quota | list


# **ListByUserUsingGET**
> InterfaceProjectQuotaResponse ListByUserUsingGET(ctx, projectId, userId)
listByUser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **userId** | **int32**| user_id | 

### Return type

[**InterfaceProjectQuotaResponse**](InterfaceProjectQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotaUsedUsingGET**
> QuotaResponse ListQuotaUsedUsingGET(ctx, projectId)
listQuotaUsed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**QuotaResponse**](QuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET**
> InterfaceProjectQuotaResponse ListUsingGET(ctx, projectId)
list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 

### Return type

[**InterfaceProjectQuotaResponse**](InterfaceProjectQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

