# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig1**](RelationalConfigurationGroupAPIApi.md#CreateConfig1) | **Post** /v1/configurations/create | 
[**DeleteConfigs1**](RelationalConfigurationGroupAPIApi.md#DeleteConfigs1) | **Delete** /v1/configurations/delete | 
[**GetConfigParams1**](RelationalConfigurationGroupAPIApi.md#GetConfigParams1) | **Get** /v1/configurations/params | 
[**GetConfigsById1**](RelationalConfigurationGroupAPIApi.md#GetConfigsById1) | **Get** /v1/configurations/id | 
[**GetListConfigs1**](RelationalConfigurationGroupAPIApi.md#GetListConfigs1) | **Get** /v1/configurations | 
[**UpdateConfig1**](RelationalConfigurationGroupAPIApi.md#UpdateConfig1) | **Put** /v1/configurations/update | 

# **CreateConfig1**
> WrapContentItemConfigInfo CreateConfig1(ctx, body)


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

# **DeleteConfigs1**
> WrapContentListDeleteConfigResponse DeleteConfigs1(ctx, body)


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

# **GetConfigParams1**
> WrapContentListConfigurationParamInfo GetConfigParams1(ctx, datastoreType, datastoreVersion)


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

# **GetConfigsById1**
> WrapContentItemConfigInfo GetConfigsById1(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListConfigs1**
> WrapContentConfigInfoGatewayResponse GetListConfigs1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RelationalConfigurationGroupAPIApiGetListConfigs1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalConfigurationGroupAPIApiGetListConfigs1Opts struct
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

# **UpdateConfig1**
> WrapContentItemConfigInfo UpdateConfig1(ctx, body)


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

