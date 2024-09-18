# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrderCluster**](KafkaClusterAPIApi.md#CreateOrderCluster) | **Post** /vdb-kafka/clusters | 
[**CreateSecRule**](KafkaClusterAPIApi.md#CreateSecRule) | **Post** /vdb-kafka/clusters/{clusterId}/security-group-rules | 
[**CreateTopic**](KafkaClusterAPIApi.md#CreateTopic) | **Post** /vdb-kafka/clusters/{clusterId}/topics | 
[**CreateUser**](KafkaClusterAPIApi.md#CreateUser) | **Post** /vdb-kafka/clusters/{clusterId}/users | 
[**DeleteCluster**](KafkaClusterAPIApi.md#DeleteCluster) | **Delete** /vdb-kafka/clusters/{clusterId} | 
[**DeleteSecRule**](KafkaClusterAPIApi.md#DeleteSecRule) | **Delete** /vdb-kafka/clusters/{clusterId}/security-group-rules/{secGroupRuleId} | 
[**DeleteTopic**](KafkaClusterAPIApi.md#DeleteTopic) | **Delete** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
[**DeleteUser**](KafkaClusterAPIApi.md#DeleteUser) | **Delete** /vdb-kafka/clusters/{clusterId}/users/{userId} | 
[**GetClusterById**](KafkaClusterAPIApi.md#GetClusterById) | **Get** /vdb-kafka/clusters/{clusterId} | 
[**GetTopicById**](KafkaClusterAPIApi.md#GetTopicById) | **Get** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
[**GetUserAuthenCreds**](KafkaClusterAPIApi.md#GetUserAuthenCreds) | **Get** /vdb-kafka/clusters/{clusterId}/users/{userId}/authen-creds | 
[**GetUserById**](KafkaClusterAPIApi.md#GetUserById) | **Get** /vdb-kafka/clusters/{clusterId}/users/{userId} | 
[**ListClusters**](KafkaClusterAPIApi.md#ListClusters) | **Get** /vdb-kafka/clusters | 
[**ListHistory**](KafkaClusterAPIApi.md#ListHistory) | **Get** /vdb-kafka/clusters/{clusterId}/history | 
[**ListTopic**](KafkaClusterAPIApi.md#ListTopic) | **Get** /vdb-kafka/clusters/{clusterId}/topics | 
[**ListUser**](KafkaClusterAPIApi.md#ListUser) | **Get** /vdb-kafka/clusters/{clusterId}/users | 
[**RegenerateUserAuthenCreds**](KafkaClusterAPIApi.md#RegenerateUserAuthenCreds) | **Put** /vdb-kafka/clusters/{clusterId}/users/{userId}/regenerate-creds | 
[**UpdateAuthentication**](KafkaClusterAPIApi.md#UpdateAuthentication) | **Put** /vdb-kafka/clusters/{clusterId}/authentication | 
[**UpdateBrokerCount**](KafkaClusterAPIApi.md#UpdateBrokerCount) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-broker-count | 
[**UpdateConfigGroup**](KafkaClusterAPIApi.md#UpdateConfigGroup) | **Put** /vdb-kafka/clusters/{clusterId}/config-group | 
[**UpdatePublicAccess**](KafkaClusterAPIApi.md#UpdatePublicAccess) | **Put** /vdb-kafka/clusters/{clusterId}/public-access | 
[**UpdateStorageSize**](KafkaClusterAPIApi.md#UpdateStorageSize) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-storage-size | 
[**UpdateStorageType**](KafkaClusterAPIApi.md#UpdateStorageType) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-storage-type | 
[**UpdateTopic**](KafkaClusterAPIApi.md#UpdateTopic) | **Put** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
[**UpdateUser**](KafkaClusterAPIApi.md#UpdateUser) | **Put** /vdb-kafka/clusters/{clusterId}/users/{userId} | 

# **CreateOrderCluster**
> WrapContentListOrderResponse CreateOrderCluster(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateKafkaClusterRequest**](CreateKafkaClusterRequest.md)|  | 
 **optional** | ***KafkaClusterAPIApiCreateOrderClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiCreateOrderClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userType** | **optional.**|  | [default to ROOT_USER]

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSecRule**
> string CreateSecRule(ctx, body, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTopic**
> string CreateTopic(ctx, body, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> string CreateUser(ctx, body, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCluster**
> string DeleteCluster(ctx, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecRule**
> string DeleteSecRule(ctx, clusterId, secGroupRuleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **secGroupRuleId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTopic**
> string DeleteTopic(ctx, clusterId, topicId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **topicId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> string DeleteUser(ctx, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterById**
> string GetClusterById(ctx, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopicById**
> string GetTopicById(ctx, clusterId, topicId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **topicId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserAuthenCreds**
> []string GetUserAuthenCreds(ctx, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserById**
> string GetUserById(ctx, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusters**
> string ListClusters(ctx, )


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

# **ListHistory**
> string ListHistory(ctx, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTopic**
> string ListTopic(ctx, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUser**
> string ListUser(ctx, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateUserAuthenCreds**
> string RegenerateUserAuthenCreds(ctx, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthentication**
> string UpdateAuthentication(ctx, clusterId, mtlsAuthen, saslAuthen)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **mtlsAuthen** | **bool**|  | 
  **saslAuthen** | **bool**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBrokerCount**
> WrapContentListOrderResponse UpdateBrokerCount(ctx, clusterId, count, rebalance, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **count** | **int32**|  | 
  **rebalance** | **bool**|  | 
 **optional** | ***KafkaClusterAPIApiUpdateBrokerCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiUpdateBrokerCountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userType** | **optional.String**|  | [default to ROOT_USER]

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigGroup**
> string UpdateConfigGroup(ctx, clusterId, configGroupVersionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **configGroupVersionId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePublicAccess**
> string UpdatePublicAccess(ctx, clusterId, enable)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **enable** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStorageSize**
> WrapContentListOrderResponse UpdateStorageSize(ctx, clusterId, size, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **size** | **int32**|  | 
 **optional** | ***KafkaClusterAPIApiUpdateStorageSizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiUpdateStorageSizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userType** | **optional.String**|  | [default to ROOT_USER]

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStorageType**
> WrapContentListOrderResponse UpdateStorageType(ctx, clusterId, storageType, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **storageType** | **string**|  | 
 **optional** | ***KafkaClusterAPIApiUpdateStorageTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiUpdateStorageTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userType** | **optional.String**|  | [default to ROOT_USER]

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTopic**
> string UpdateTopic(ctx, body, clusterId, topicId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **clusterId** | **string**|  | 
  **topicId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> string UpdateUser(ctx, body, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **clusterId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

