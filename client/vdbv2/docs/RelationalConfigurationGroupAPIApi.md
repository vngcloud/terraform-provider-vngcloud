# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig**](RelationalConfigurationGroupAPIApi.md#CreateConfig) | **Post** /vdb-relational/v1/configurations/create | 
[**DeleteConfigs1**](RelationalConfigurationGroupAPIApi.md#DeleteConfigs1) | **Delete** /vdb-relational/v1/configurations/delete | 
[**GetConfigParams**](RelationalConfigurationGroupAPIApi.md#GetConfigParams) | **Get** /vdb-relational/v1/configurations/params | 
[**GetConfigsById**](RelationalConfigurationGroupAPIApi.md#GetConfigsById) | **Get** /vdb-relational/v1/configurations/id | 
[**GetListConfigs**](RelationalConfigurationGroupAPIApi.md#GetListConfigs) | **Get** /vdb-relational/v1/configurations | 
[**UpdateConfig**](RelationalConfigurationGroupAPIApi.md#UpdateConfig) | **Put** /vdb-relational/v1/configurations/update | 

# **CreateConfig**
> WrapContentItemConfigInfo CreateConfig(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateConfigGroupRequest**](CreateConfigGroupRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigs1**
> WrapContentListDeleteConfigResponse DeleteConfigs1(ctx, body, portalUserId)


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

# **GetConfigParams**
> WrapContentListConfigurationParamInfo GetConfigParams(ctx, portalUserId, datastoreType, datastoreVersion, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **datastoreType** | **string**|  | 
  **datastoreVersion** | **string**|  | 
 **optional** | ***RelationalConfigurationGroupAPIApiGetConfigParamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalConfigurationGroupAPIApiGetConfigParamsOpts struct
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

# **GetConfigsById**
> WrapContentItemConfigInfo GetConfigsById(ctx, portalUserId, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **id** | **string**|  | 

### Return type

[**WrapContentItemConfigInfo**](WrapContentItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListConfigs**
> WrapContentConfigInfoGatewayResponse GetListConfigs(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalConfigurationGroupAPIApiGetListConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalConfigurationGroupAPIApiGetListConfigsOpts struct
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
> WrapContentItemConfigInfo UpdateConfig(ctx, body, portalUserId)


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

