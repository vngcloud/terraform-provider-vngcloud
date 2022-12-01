# \PersistentVolumeRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePersistentVolumeUsingDELETE**](PersistentVolumeRestControllerV2Api.md#DeletePersistentVolumeUsingDELETE) | **Delete** /v2/{projectId}/persistent-volumes/{pvId} | Delete Persistent Volume
[**ListPersistentVolumeUsingGET**](PersistentVolumeRestControllerV2Api.md#ListPersistentVolumeUsingGET) | **Get** /v2/{projectId}/persistent-volumes | List Persistent Volume


# **DeletePersistentVolumeUsingDELETE**
> DataResponse DeletePersistentVolumeUsingDELETE(ctx, projectId, pvId)
Delete Persistent Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **pvId** | **string**| The persistent volume id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPersistentVolumeUsingGET**
> PagingPersistentVolumeDto ListPersistentVolumeUsingGET(ctx, projectId, optional)
List Persistent Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***PersistentVolumeRestControllerV2ApiListPersistentVolumeUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersistentVolumeRestControllerV2ApiListPersistentVolumeUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingPersistentVolumeDto**](Paging«PersistentVolumeDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

