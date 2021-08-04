# \VdbConfigurationGroupEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigUsingPOST**](VdbConfigurationGroupEndPointApi.md#CreateConfigUsingPOST) | **Post** /v1/trv/{projectId}/configuration | createConfig
[**DeleteConfigUsingDELETE**](VdbConfigurationGroupEndPointApi.md#DeleteConfigUsingDELETE) | **Delete** /v1/trv/{projectId}/configuration/{configId} | deleteConfig
[**GetAllConfigsUsingGET**](VdbConfigurationGroupEndPointApi.md#GetAllConfigsUsingGET) | **Get** /v1/trv/{projectId}/configuration | getAllConfigs
[**GetConfigByIdUsingGET**](VdbConfigurationGroupEndPointApi.md#GetConfigByIdUsingGET) | **Get** /v1/trv/{projectId}/configuration/{configId} | getConfigById
[**GetConfigParamsUsingGET**](VdbConfigurationGroupEndPointApi.md#GetConfigParamsUsingGET) | **Get** /v1/trv/{projectId}/configuration/params/{datastoreType}/{datastoreVersion} | getConfigParams
[**UpdateConfigValuesUsingPUT**](VdbConfigurationGroupEndPointApi.md#UpdateConfigValuesUsingPUT) | **Put** /v1/trv/{projectId}/configuration/{configId} | updateConfigValues


# **CreateConfigUsingPOST**
> ConfigGroupResponse CreateConfigUsingPOST(ctx, projectId, request)
createConfig

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**ConfigurationRequest**](ConfigurationRequest.md)| request | 

### Return type

[**ConfigGroupResponse**](ConfigGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigUsingDELETE**
> DeleteConfigGroupResponse DeleteConfigUsingDELETE(ctx, configId, projectId, request)
deleteConfig

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**| configId | 
  **projectId** | **string**| projectId | 
  **request** | [**ConfigurationRequest**](ConfigurationRequest.md)| request | 

### Return type

[**DeleteConfigGroupResponse**](DeleteConfigGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllConfigsUsingGET**
> ListConfigGroupResponse GetAllConfigsUsingGET(ctx, projectId)
getAllConfigs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**ListConfigGroupResponse**](ListConfigGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigByIdUsingGET**
> ConfigGroupResponse GetConfigByIdUsingGET(ctx, configId, projectId)
getConfigById

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**| configId | 
  **projectId** | **string**| projectId | 

### Return type

[**ConfigGroupResponse**](ConfigGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigParamsUsingGET**
> ListConfigParamResult GetConfigParamsUsingGET(ctx, datastoreType, datastoreVersion, projectId)
getConfigParams

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datastoreType** | **string**| datastoreType | 
  **datastoreVersion** | **string**| datastoreVersion | 
  **projectId** | **string**| projectId | 

### Return type

[**ListConfigParamResult**](ListConfigParamResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigValuesUsingPUT**
> ConfigGroupResponse UpdateConfigValuesUsingPUT(ctx, configId, projectId, request)
updateConfigValues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**| configId | 
  **projectId** | **string**| projectId | 
  **request** | [**ConfigurationRequest**](ConfigurationRequest.md)| request | 

### Return type

[**ConfigGroupResponse**](ConfigGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

