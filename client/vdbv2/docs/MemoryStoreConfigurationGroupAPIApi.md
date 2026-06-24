# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig1**](MemoryStoreConfigurationGroupAPIApi.md#CreateConfig1) | **Post** /vdb-memory/v1/configurations/create | 
[**DeleteConfigs**](MemoryStoreConfigurationGroupAPIApi.md#DeleteConfigs) | **Post** /vdb-memory/v1/configurations/delete | 
[**GetConfigParams1**](MemoryStoreConfigurationGroupAPIApi.md#GetConfigParams1) | **Get** /vdb-memory/v1/configurations/params | 
[**GetConfigsById1**](MemoryStoreConfigurationGroupAPIApi.md#GetConfigsById1) | **Get** /vdb-memory/v1/configurations/{configId}/detail | 
[**GetListConfigs1**](MemoryStoreConfigurationGroupAPIApi.md#GetListConfigs1) | **Get** /vdb-memory/v1/configurations | 
[**UpdateConfig1**](MemoryStoreConfigurationGroupAPIApi.md#UpdateConfig1) | **Put** /vdb-memory/v1/configurations/update | 

# **CreateConfig1**
> WrapContentItemConfigInfo CreateConfig1(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMemConfigGroupRequest**](CreateMemConfigGroupRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigs**
> WrapContentListDeleteConfigResponse DeleteConfigs(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DeleteConfigGroupRequest**](DeleteConfigGroupRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListDeleteConfigResponse**](WrapContentListDeleteConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigParams1**
> WrapContentListConfigurationParamInfo GetConfigParams1(ctx, portalUserId, datastoreType, datastoreVersion, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **datastoreType** | **string**|  | 
  **datastoreVersion** | **string**|  | 
 **optional** | ***MemoryStoreConfigurationGroupAPIApiGetConfigParams1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreConfigurationGroupAPIApiGetConfigParams1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deployType** | **optional.String**|  | 

### Return type

[**WrapContentListConfigurationParamInfo**](WrapContentListConfigurationParamInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigsById1**
> WrapContentItemConfigInfo GetConfigsById1(ctx, configId)


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

# **GetListConfigs1**
> WrapContentConfigInfoGatewayResponse GetListConfigs1(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreConfigurationGroupAPIApiGetListConfigs1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreConfigurationGroupAPIApiGetListConfigs1Opts struct
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
> WrapContentItemConfigInfo UpdateConfig1(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateConfigGroupRequest**](UpdateConfigGroupRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

