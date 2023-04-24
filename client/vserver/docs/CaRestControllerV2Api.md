# \CaRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCertificateAuthorityUsingDELETE**](CaRestControllerV2Api.md#DeleteCertificateAuthorityUsingDELETE) | **Delete** /v2/{projectId}/cas/{caId} | Delete Certificate Authority
[**GetAllListCertificateAuthorityUsingGET**](CaRestControllerV2Api.md#GetAllListCertificateAuthorityUsingGET) | **Get** /v2/{projectId}/cas | List All Certificate Authority
[**ImportCAUsingPOST**](CaRestControllerV2Api.md#ImportCAUsingPOST) | **Post** /v2/{projectId}/cas | Import Certificate Authority


# **DeleteCertificateAuthorityUsingDELETE**
> DeleteCertificateAuthorityUsingDELETE(ctx, caId, projectId)
Delete Certificate Authority

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caId** | **string**| The ca id | 
  **projectId** | **string**| The project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllListCertificateAuthorityUsingGET**
> PagingCaDto GetAllListCertificateAuthorityUsingGET(ctx, projectId, optional)
List All Certificate Authority

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***CaRestControllerV2ApiGetAllListCertificateAuthorityUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CaRestControllerV2ApiGetAllListCertificateAuthorityUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 

### Return type

[**PagingCaDto**](Paging«CADto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportCAUsingPOST**
> DataResponseCaDto ImportCAUsingPOST(ctx, importCARequest, projectId)
Import Certificate Authority

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importCARequest** | [**ImportCaRequest**](ImportCaRequest.md)| importCARequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseCaDto**](DataResponse«CADto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

