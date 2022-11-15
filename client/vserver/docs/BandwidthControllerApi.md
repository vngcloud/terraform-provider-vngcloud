# \BandwidthControllerApi

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTypeBandwidthUsingPUT**](BandwidthControllerApi.md#ChangeTypeBandwidthUsingPUT) | **Put** /v1/{project_id}/bandwidth/change-type | Change type bandwidth
[**CreateBandwidthUsingPOST**](BandwidthControllerApi.md#CreateBandwidthUsingPOST) | **Post** /v1/{project_id}/bandwidth | Create bandwidth
[**DeleteBandwidthUsingDELETE**](BandwidthControllerApi.md#DeleteBandwidthUsingDELETE) | **Delete** /v1/{project_id}/bandwidth | Delete bandwidth
[**GetBandwidthUsingGET**](BandwidthControllerApi.md#GetBandwidthUsingGET) | **Get** /v1/{project_id}/bandwidth/{uuid} | Get Bandwidth
[**ListAllTrafficUsingGET**](BandwidthControllerApi.md#ListAllTrafficUsingGET) | **Get** /v1/{project_id}/bandwidth/traffic | List all traffic
[**ListBandwidthsUsingGET**](BandwidthControllerApi.md#ListBandwidthsUsingGET) | **Get** /v1/{project_id}/bandwidth | List Bandwidths
[**ListPackagesUsingGET**](BandwidthControllerApi.md#ListPackagesUsingGET) | **Get** /v1/{project_id}/bandwidth/packages | List Bandwidth Packages
[**ListResourceIpsUsingGET**](BandwidthControllerApi.md#ListResourceIpsUsingGET) | **Get** /v1/{project_id}/bandwidth/IP | List resource IPs
[**RecoverBandwidthUsingPUT**](BandwidthControllerApi.md#RecoverBandwidthUsingPUT) | **Put** /v1/{project_id}/bandwidth/recover | Recover bandwidth
[**RenewBandwidthUsingPUT**](BandwidthControllerApi.md#RenewBandwidthUsingPUT) | **Put** /v1/{project_id}/bandwidth/renew | Renew bandwidth


# **ChangeTypeBandwidthUsingPUT**
> BandwidthResponse ChangeTypeBandwidthUsingPUT(ctx, changeTypeBandwidthRequest, projectId)
Change type bandwidth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeTypeBandwidthRequest** | [**ChangeTypeBandwidthRequest**](ChangeTypeBandwidthRequest.md)| changeTypeBandwidthRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BandwidthResponse**](BandwidthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBandwidthUsingPOST**
> BandwidthResponse CreateBandwidthUsingPOST(ctx, createBandwidthRequest, projectId)
Create bandwidth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createBandwidthRequest** | [**CreateBandwidthRequest**](CreateBandwidthRequest.md)| createBandwidthRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BandwidthResponse**](BandwidthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBandwidthUsingDELETE**
> BaseResponse DeleteBandwidthUsingDELETE(ctx, deleteBandwidthRequest, projectId)
Delete bandwidth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteBandwidthRequest** | [**DeleteBandwidthRequest**](DeleteBandwidthRequest.md)| deleteBandwidthRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBandwidthUsingGET**
> BandwidthResponse GetBandwidthUsingGET(ctx, projectId, uuid)
Get Bandwidth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The bandwidth id | 

### Return type

[**BandwidthResponse**](BandwidthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllTrafficUsingGET**
> BandwidthTrafficResponse ListAllTrafficUsingGET(ctx, projectId)
List all traffic

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**BandwidthTrafficResponse**](BandwidthTrafficResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBandwidthsUsingGET**
> BandwidthResponse ListBandwidthsUsingGET(ctx, projectId)
List Bandwidths

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**BandwidthResponse**](BandwidthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPackagesUsingGET**
> BandwidthPackageResponse ListPackagesUsingGET(ctx, projectId)
List Bandwidth Packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**BandwidthPackageResponse**](BandwidthPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListResourceIpsUsingGET**
> BandwidthResourceIpResponse ListResourceIpsUsingGET(ctx, listResourceIpRequest, projectId)
List resource IPs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listResourceIpRequest** | [**ListResourceIpRequest**](ListResourceIpRequest.md)| listResourceIpRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BandwidthResourceIpResponse**](BandwidthResourceIpResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecoverBandwidthUsingPUT**
> BandwidthResponse RecoverBandwidthUsingPUT(ctx, projectId, recoverBandwidthRequest)
Recover bandwidth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **recoverBandwidthRequest** | [**RecoverBandwidthRequest**](RecoverBandwidthRequest.md)| recoverBandwidthRequest | 

### Return type

[**BandwidthResponse**](BandwidthResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewBandwidthUsingPUT**
> BaseResponse RenewBandwidthUsingPUT(ctx, projectId, renewBandwidthRequest)
Renew bandwidth

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **renewBandwidthRequest** | [**RenewBandwidthRequest**](RenewBandwidthRequest.md)| renewBandwidthRequest | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

