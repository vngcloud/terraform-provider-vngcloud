# \ElasticIpRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListElasticIpUsingGET**](ElasticIpRestControllerV2Api.md#ListElasticIpUsingGET) | **Get** /v2/{projectId}/elastic-ips | List elastic-ip


# **ListElasticIpUsingGET**
> PagingElasticIpDetail ListElasticIpUsingGET(ctx, projectId, optional)
List elastic-ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***ElasticIpRestControllerV2ApiListElasticIpUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElasticIpRestControllerV2ApiListElasticIpUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingElasticIpDetail**](Paging«ElasticIpDetail».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

