# \QuotaRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQuotaUsedUsingGET1**](QuotaRestControllerV2Api.md#ListQuotaUsedUsingGET1) | **Get** /v2/{projectId}/quotas/quotaUsed | List Quota-used


# **ListQuotaUsedUsingGET1**
> DataResponseListQuotaUsedDto ListQuotaUsedUsingGET1(ctx, projectId)
List Quota-used

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 

### Return type

[**DataResponseListQuotaUsedDto**](DataResponse«List«QuotaUsedDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

