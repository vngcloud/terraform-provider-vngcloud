# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVolumeTypes1**](KafkaMiscAPIApi.md#GetVolumeTypes1) | **Get** /vdb-kafka/database/volume-types | 
[**ListAppConfigs**](KafkaMiscAPIApi.md#ListAppConfigs) | **Get** /vdb-kafka/database/configs | 
[**ListCode**](KafkaMiscAPIApi.md#ListCode) | **Get** /vdb-kafka/database/codes | 
[**ListFamily**](KafkaMiscAPIApi.md#ListFamily) | **Get** /vdb-kafka/database/families | 
[**ListFlavor**](KafkaMiscAPIApi.md#ListFlavor) | **Get** /vdb-kafka/database/flavors | 

# **GetVolumeTypes1**
> WrapContentVolumeTypeGatewayResponse GetVolumeTypes1(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentVolumeTypeGatewayResponse**](WrapContentVolumeTypeGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppConfigs**
> WrapContentMapStringString ListAppConfigs(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentMapStringString**](WrapContentMapStringString.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCode**
> WrapContentSetFlavorCode ListCode(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentSetFlavorCode**](WrapContentSetFlavorCode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFamily**
> WrapContentListInstanceFamily ListFamily(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListInstanceFamily**](WrapContentListInstanceFamily.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFlavor**
> WrapContentListFlavorInfo ListFlavor(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***KafkaMiscAPIApiListFlavorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaMiscAPIApiListFlavorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | [default to kafka]
 **version** | **optional.String**|  | [default to 1.0]

### Return type

[**WrapContentListFlavorInfo**](WrapContentListFlavorInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

