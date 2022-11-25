# \K8SClusterRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListClusterUsingGET**](K8SClusterRestControllerV2Api.md#ListClusterUsingGET) | **Get** /v2/{projectId}/clusters | List cluster


# **ListClusterUsingGET**
> PagingInterfaceK8SClusterDetail ListClusterUsingGET(ctx, projectId, optional)
List cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***K8SClusterRestControllerV2ApiListClusterUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8SClusterRestControllerV2ApiListClusterUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingInterfaceK8SClusterDetail**](Paging«InterfaceK8SClusterDetail».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

