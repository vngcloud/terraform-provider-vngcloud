# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigGroup**](KafkaConfigurationGroupAPIApi.md#CreateConfigGroup) | **Post** /vdb-kafka/config-groups | 
[**CreateConfigGroupVersion**](KafkaConfigurationGroupAPIApi.md#CreateConfigGroupVersion) | **Post** /vdb-kafka/config-groups/{configGroupId}/versions | 
[**DeleteConfigGroup**](KafkaConfigurationGroupAPIApi.md#DeleteConfigGroup) | **Delete** /vdb-kafka/config-groups/{configGroupId} | 
[**GetConfigGroupById**](KafkaConfigurationGroupAPIApi.md#GetConfigGroupById) | **Get** /vdb-kafka/config-groups/{configGroupId} | 
[**GetConfigGroupVersionById**](KafkaConfigurationGroupAPIApi.md#GetConfigGroupVersionById) | **Get** /vdb-kafka/config-groups/{configGroupId}/versions/{configGroupVersionId} | 
[**ListConfigGroup**](KafkaConfigurationGroupAPIApi.md#ListConfigGroup) | **Get** /vdb-kafka/config-groups | 

# **CreateConfigGroup**
> string CreateConfigGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConfigGroupVersion**
> string CreateConfigGroupVersion(ctx, body, configGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **configGroupId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigGroup**
> string DeleteConfigGroup(ctx, configGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configGroupId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigGroupById**
> string GetConfigGroupById(ctx, configGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configGroupId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigGroupVersionById**
> string GetConfigGroupVersionById(ctx, configGroupId, configGroupVersionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configGroupId** | **string**|  | 
  **configGroupVersionId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConfigGroup**
> string ListConfigGroup(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

