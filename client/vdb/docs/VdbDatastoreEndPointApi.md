# \VdbDatastoreEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListDatastoreByEngineGroupUsingGET**](VdbDatastoreEndPointApi.md#GetListDatastoreByEngineGroupUsingGET) | **Get** /v1/trv/{projectId}/datastore | getListDatastoreByEngineGroup


# **GetListDatastoreByEngineGroupUsingGET**
> BaseResponse GetListDatastoreByEngineGroupUsingGET(ctx, engineGroup, projectId)
getListDatastoreByEngineGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineGroup** | **int32**| engineGroup | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

