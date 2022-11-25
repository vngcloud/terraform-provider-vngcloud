# \UserImageRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImageUsingPOST**](UserImageRestControllerV2Api.md#CreateImageUsingPOST) | **Post** /v2/{projectId}/user-images/servers/{serverId} | Start server
[**DeleteImageUsingDELETE**](UserImageRestControllerV2Api.md#DeleteImageUsingDELETE) | **Delete** /v2/{projectId}/user-images/{imageId} | Delete User Image
[**ListImageByIdUsingGET**](UserImageRestControllerV2Api.md#ListImageByIdUsingGET) | **Get** /v2/{projectId}/user-images/{imageId} | List User Image By ID
[**ListImageUsingGET**](UserImageRestControllerV2Api.md#ListImageUsingGET) | **Get** /v2/{projectId}/user-images | List User Image


# **CreateImageUsingPOST**
> DataResponse CreateImageUsingPOST(ctx, createUserImageRequest, projectId, serverId)
Start server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createUserImageRequest** | [**CreateUserImageRequest**](CreateUserImageRequest.md)| createUserImageRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImageUsingDELETE**
> DeleteImageUsingDELETE(ctx, imageId, projectId)
Delete User Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **imageId** | **string**| The image id | 
  **projectId** | **string**| The project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListImageByIdUsingGET**
> UserImage ListImageByIdUsingGET(ctx, imageId, projectId)
List User Image By ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **imageId** | **string**| The image id | 
  **projectId** | **string**| The project id | 

### Return type

[**UserImage**](UserImage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListImageUsingGET**
> PagingUserImage ListImageUsingGET(ctx, projectId, optional)
List User Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***UserImageRestControllerV2ApiListImageUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserImageRestControllerV2ApiListImageUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingUserImage**](Paging«UserImage».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

