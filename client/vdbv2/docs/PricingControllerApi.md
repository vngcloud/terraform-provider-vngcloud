# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculatePrice**](PricingControllerApi.md#CalculatePrice) | **Post** /v1/price | 

# **CalculatePrice**
> CalculatePriceResponse CalculatePrice(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CalculatePriceRequest**](CalculatePriceRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**CalculatePriceResponse**](CalculatePriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

