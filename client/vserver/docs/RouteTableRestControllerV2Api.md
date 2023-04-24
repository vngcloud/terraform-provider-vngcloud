# \RouteTableRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/Ev4LiA/vserver/1.0.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteTableUsingPOST1**](RouteTableRestControllerV2Api.md#CreateRouteTableUsingPOST1) | **Post** /v2/{projectId}/route-table | Create Route-table
[**DeleteRouteTableUsingDELETE1**](RouteTableRestControllerV2Api.md#DeleteRouteTableUsingDELETE1) | **Delete** /v2/{projectId}/route-table/{routeId} | Delete Route Table
[**GetRouteTableUsingGET1**](RouteTableRestControllerV2Api.md#GetRouteTableUsingGET1) | **Get** /v2/{projectId}/route-table/{uuid} | Get Route table by uuid
[**ListRouteTablesWithPagingUsingGET**](RouteTableRestControllerV2Api.md#ListRouteTablesWithPagingUsingGET) | **Get** /v2/{projectId}/route-table/route/{routeTableId} | List Route from Route-Table
[**ListRouteTablesWithPagingUsingGET1**](RouteTableRestControllerV2Api.md#ListRouteTablesWithPagingUsingGET1) | **Get** /v2/{projectId}/route-table | List Route Table V2
[**UpdateRouteTableDetailUsingPUT**](RouteTableRestControllerV2Api.md#UpdateRouteTableDetailUsingPUT) | **Put** /v2/{projectId}/route-table/{uuid}/routes | Add or remove Route to Route-Table


# **CreateRouteTableUsingPOST1**
> DataResponseRouteTableDto CreateRouteTableUsingPOST1(ctx, createRouteTableRequest, projectId)
Create Route-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createRouteTableRequest** | [**CreateRouteTableRequest**](CreateRouteTableRequest.md)| createRouteTableRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseRouteTableDto**](DataResponse«RouteTableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteTableUsingDELETE1**
> DeleteRouteTableUsingDELETE1(ctx, projectId, routeId)
Delete Route Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **routeId** | **string**| routeId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteTableUsingGET1**
> DataResponseRouteTableDto GetRouteTableUsingGET1(ctx, projectId, uuid)
Get Route table by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **uuid** | **string**| The route-table uuid | 

### Return type

[**DataResponseRouteTableDto**](DataResponse«RouteTableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRouteTablesWithPagingUsingGET**
> DataResponseListRoutesDto ListRouteTablesWithPagingUsingGET(ctx, projectId, routeTableId)
List Route from Route-Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **routeTableId** | **string**| routeTableId | 

### Return type

[**DataResponseListRoutesDto**](DataResponse«List«RoutesDto»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRouteTablesWithPagingUsingGET1**
> PagingRouteTableDto ListRouteTablesWithPagingUsingGET1(ctx, name, page, projectId, size)
List Route Table V2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
  **projectId** | **string**| project id | 
  **size** | **string**| size | [default to 10]

### Return type

[**PagingRouteTableDto**](Paging«RouteTableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouteTableDetailUsingPUT**
> DataResponseRouteTableDto UpdateRouteTableDetailUsingPUT(ctx, projectId, updateRouteRequest, uuid)
Add or remove Route to Route-Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **updateRouteRequest** | [**RouteTableUpdateRouteRequest**](RouteTableUpdateRouteRequest.md)| updateRouteRequest | 
  **uuid** | **string**| The route-table uuid | 

### Return type

[**DataResponseRouteTableDto**](DataResponse«RouteTableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

