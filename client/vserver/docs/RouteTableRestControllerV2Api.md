# \RouteTableRestControllerV2Api

All URIs are relative to *https://localhost:8089*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteTableUsingPOST1**](RouteTableRestControllerV2Api.md#CreateRouteTableUsingPOST1) | **Post** /v2/{projectId}/route-table | Create Route-table
[**DeleteRouteTableUsingDELETE1**](RouteTableRestControllerV2Api.md#DeleteRouteTableUsingDELETE1) | **Delete** /v2/{projectId}/route-table/{routeId} | Delete Route Table
[**GetRouteTableUsingGET1**](RouteTableRestControllerV2Api.md#GetRouteTableUsingGET1) | **Get** /v2/{projectId}/route-table/{uuid} | Get Route table by uuid
[**ListRouteFromRouteTableIdUsingGET**](RouteTableRestControllerV2Api.md#ListRouteFromRouteTableIdUsingGET) | **Get** /v2/{projectId}/route-table/route/{routeTableId} | List Route from Route-Table
[**ListRouteTableUsingGET**](RouteTableRestControllerV2Api.md#ListRouteTableUsingGET) | **Get** /v2/{projectId}/route-table | List Route Table V2
[**UpdateRouteTableDetailUsingPUT**](RouteTableRestControllerV2Api.md#UpdateRouteTableDetailUsingPUT) | **Put** /v2/{projectId}/route-table/{uuid}/routes | Add or remove Route to Route-Table


# **CreateRouteTableUsingPOST1**
> map[string]string CreateRouteTableUsingPOST1(ctx, authorization, createReq, portalUserId, projectId)
Create Route-table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **createReq** | [**CreateRouteTableRequest**](CreateRouteTableRequest.md)| createReq | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteTableUsingDELETE1**
> DeleteRouteTableUsingDELETE1(ctx, authorization, portalUserId, projectId, routeId)
Delete Route Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **portalUserId** | **int32**| portal-user-id | 
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
> DataResponseRouteTableDto GetRouteTableUsingGET1(ctx, authorization, portalUserId, projectId, uuid)
Get Route table by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **portalUserId** | **int32**| portal-user-id | 
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

# **ListRouteFromRouteTableIdUsingGET**
> DataResponseListRoutesDto ListRouteFromRouteTableIdUsingGET(ctx, authorization, portalUserId, projectId, routeTableId)
List Route from Route-Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **portalUserId** | **int32**| portal-user-id | 
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

# **ListRouteTableUsingGET**
> PagingRouteTableDto ListRouteTableUsingGET(ctx, authorization, name, page, portalUserId, projectId, size)
List Route Table V2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **name** | **string**| name | 
  **page** | **string**| page | [default to 1]
  **portalUserId** | **int32**| portal-user-id | 
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
> UpdateRouteTableDetailUsingPUT(ctx, authorization, changeRoutesReq, portalUserId, projectId, uuid)
Add or remove Route to Route-Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **changeRoutesReq** | [**ChangeRoutesRequest**](ChangeRoutesRequest.md)| changeRoutesReq | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 
  **uuid** | **string**| The route-table uuid | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

