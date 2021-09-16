# \QuotaRestControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitUsedUsingPUT**](QuotaRestControllerApi.md#InitUsedUsingPUT) | **Put** /v1/{project_id}/quota/init-used | initUsed
[**ListByUserUsingGET**](QuotaRestControllerApi.md#ListByUserUsingGET) | **Get** /v1/{project_id}/quota/user/{user_id} | listByUser
[**ListQuotaUsedUsingGET**](QuotaRestControllerApi.md#ListQuotaUsedUsingGET) | **Get** /v1/{project_id}/quote-used | listQuotaUsed
[**ListUsingGET**](QuotaRestControllerApi.md#ListUsingGET) | **Get** /v1/{project_id}/quota | list
[**UpdateUsingPUT**](QuotaRestControllerApi.md#UpdateUsingPUT) | **Put** /v1/{project_id}/quota | update


# **InitUsedUsingPUT**
> InterfaceProjectQuotaResponse InitUsedUsingPUT(ctx, projectId)
initUsed

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

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **UpdateUsingPUT**
> InterfaceProjectQuotaResponse UpdateUsingPUT(ctx, projectId, updateQuotaVsgRequest)
update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project_id | 
  **updateQuotaVsgRequest** | [**UpdateQuotaVsgRequest**](UpdateQuotaVsgRequest.md)| updateQuotaVsgRequest | 

### Return type

[**InterfaceProjectQuotaResponse**](InterfaceProjectQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

