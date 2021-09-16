# \RouteTableControllerApi

All URIs are relative to *https://api.vngcloud.tech*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteTableUsingPOST**](RouteTableControllerApi.md#CreateRouteTableUsingPOST) | **Post** /v1/{project_id}/route-table | Create route-table
[**DeleteRouteTableUsingDELETE**](RouteTableControllerApi.md#DeleteRouteTableUsingDELETE) | **Delete** /v1/{project_id}/route-table/{uuid} | Delete Route Table
[**DeleteRouteTablesUsingDELETE**](RouteTableControllerApi.md#DeleteRouteTablesUsingDELETE) | **Delete** /v1/{project_id}/route-table | Delete all route-tables
[**GetRouteTableUsingGET**](RouteTableControllerApi.md#GetRouteTableUsingGET) | **Get** /v1/{project_id}/route-table/{uuid} | Get route-table by uuid
[**ListRouteTablesUsingGET**](RouteTableControllerApi.md#ListRouteTablesUsingGET) | **Get** /v1/{project_id}/route-table | List route-tables
[**UpdateRouteTableRoutesUsingPUT**](RouteTableControllerApi.md#UpdateRouteTableRoutesUsingPUT) | **Put** /v1/{project_id}/route-table/{uuid}/routes | Add/remove routes of route-table
[**UpdateRouteTableSubnetsUsingPUT**](RouteTableControllerApi.md#UpdateRouteTableSubnetsUsingPUT) | **Put** /v1/{project_id}/route-table/{uuid}/subnets | Add/remove subnets of route-table


# **CreateRouteTableUsingPOST**
> BaseResponse CreateRouteTableUsingPOST(ctx, createRouteTableRequest, projectId)
Create route-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createRouteTableRequest** | [**CreateRouteTableRequest**](CreateRouteTableRequest.md)| createRouteTableRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteTableUsingDELETE**
> BaseResponse DeleteRouteTableUsingDELETE(ctx, projectId, uuid)
Delete Route Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The route-table uuid | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteTablesUsingDELETE**
> BaseResponse DeleteRouteTablesUsingDELETE(ctx, projectId)
Delete all route-tables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteTableUsingGET**
> RouteTableResponse GetRouteTableUsingGET(ctx, projectId, uuid)
Get route-table by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The route-table uuid | 

### Return type

[**RouteTableResponse**](RouteTableResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRouteTablesUsingGET**
> ListRouteTablesResponse ListRouteTablesUsingGET(ctx, projectId)
List route-tables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 

### Return type

[**ListRouteTablesResponse**](ListRouteTablesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouteTableRoutesUsingPUT**
> RouteTableResponse UpdateRouteTableRoutesUsingPUT(ctx, projectId, updateRouteRequest, uuid)
Add/remove routes of route-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateRouteRequest** | [**RouteTableUpdateRouteRequest**](RouteTableUpdateRouteRequest.md)| updateRouteRequest | 
  **uuid** | **string**| The route-table uuid | 

### Return type

[**RouteTableResponse**](RouteTableResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouteTableSubnetsUsingPUT**
> RouteTableResponse UpdateRouteTableSubnetsUsingPUT(ctx, projectId, updateSubnetRequest, uuid)
Add/remove subnets of route-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateSubnetRequest** | [**RouteTableUpdateSubnetRequest**](RouteTableUpdateSubnetRequest.md)| updateSubnetRequest | 
  **uuid** | **string**| The route-table uuid | 

### Return type

[**RouteTableResponse**](RouteTableResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

