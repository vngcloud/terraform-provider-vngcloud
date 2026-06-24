# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigGroup**](KafkaConfigurationGroupAPIApi.md#CreateConfigGroup) | **Post** /vdb-kafka/config-groups | 
[**CreateConfigGroupVersion**](KafkaConfigurationGroupAPIApi.md#CreateConfigGroupVersion) | **Post** /vdb-kafka/config-groups/{configGroupId}/versions | 
[**DeleteConfigGroup**](KafkaConfigurationGroupAPIApi.md#DeleteConfigGroup) | **Delete** /vdb-kafka/config-groups/{configGroupId} | 
[**GetConfigGroupById**](KafkaConfigurationGroupAPIApi.md#GetConfigGroupById) | **Get** /vdb-kafka/config-groups/{configGroupId} | 
[**GetConfigGroupVersionById**](KafkaConfigurationGroupAPIApi.md#GetConfigGroupVersionById) | **Get** /vdb-kafka/config-groups/{configGroupId}/versions/{configGroupVersionId} | 
[**ListConfigGroup**](KafkaConfigurationGroupAPIApi.md#ListConfigGroup) | **Get** /vdb-kafka/config-groups | 

# **CreateConfigGroup**
> ConfigGroupDto CreateConfigGroup(ctx, body, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigGroupCreateRequest**](ConfigGroupCreateRequest.md)|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**ConfigGroupDto**](ConfigGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConfigGroupVersion**
> ConfigGroupVersionDto CreateConfigGroupVersion(ctx, body, portalUserId, configGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigGroupVersionCreateRequest**](ConfigGroupVersionCreateRequest.md)|  | 
  **portalUserId** | **int32**|  | 
  **configGroupId** | **string**|  | 

### Return type

[**ConfigGroupVersionDto**](ConfigGroupVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigGroup**
> string DeleteConfigGroup(ctx, portalUserId, configGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> ConfigGroupDto GetConfigGroupById(ctx, portalUserId, configGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **configGroupId** | **string**|  | 

### Return type

[**ConfigGroupDto**](ConfigGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigGroupVersionById**
> ConfigGroupVersionDto GetConfigGroupVersionById(ctx, portalUserId, configGroupId, configGroupVersionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **configGroupId** | **string**|  | 
  **configGroupVersionId** | **string**|  | 

### Return type

[**ConfigGroupVersionDto**](ConfigGroupVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConfigGroup**
> []ConfigGroupDto ListConfigGroup(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**[]ConfigGroupDto**](ConfigGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

