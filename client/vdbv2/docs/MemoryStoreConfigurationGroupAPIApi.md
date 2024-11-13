# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig**](MemoryStoreConfigurationGroupAPIApi.md#CreateConfig) | **Post** /vdb-memory/v1/configurations/create | 
[**DeleteConfigs**](MemoryStoreConfigurationGroupAPIApi.md#DeleteConfigs) | **Post** /vdb-memory/v1/configurations/delete | 
[**GetConfigParams**](MemoryStoreConfigurationGroupAPIApi.md#GetConfigParams) | **Get** /vdb-memory/v1/configurations/params | 
[**GetConfigsById**](MemoryStoreConfigurationGroupAPIApi.md#GetConfigsById) | **Get** /vdb-memory/v1/configurations/{configId}/detail | 
[**GetListConfigs**](MemoryStoreConfigurationGroupAPIApi.md#GetListConfigs) | **Get** /vdb-memory/v1/configurations | 
[**UpdateConfig**](MemoryStoreConfigurationGroupAPIApi.md#UpdateConfig) | **Put** /vdb-memory/v1/configurations/update | 

# **CreateConfig**
> WrapContentItemConfigInfo CreateConfig(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigs**
> WrapContentListDeleteConfigResponse DeleteConfigs(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

[**WrapContentListDeleteConfigResponse**](WrapContentListDeleteConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigParams**
> WrapContentListConfigurationParamInfo GetConfigParams(ctx, datastoreType, datastoreVersion)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datastoreType** | **string**|  | 
  **datastoreVersion** | **string**|  | 

### Return type

[**WrapContentListConfigurationParamInfo**](WrapContentListConfigurationParamInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigsById**
> WrapContentItemConfigInfo GetConfigsById(ctx, configId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListConfigs**
> WrapContentConfigInfoGatewayResponse GetListConfigs(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemoryStoreConfigurationGroupAPIApiGetListConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreConfigurationGroupAPIApiGetListConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 10]

### Return type

[**WrapContentConfigInfoGatewayResponse**](WrapContentConfigInfoGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfig**
> WrapContentItemConfigInfo UpdateConfig(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

