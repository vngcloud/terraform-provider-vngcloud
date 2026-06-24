# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

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
[**GetUserAuthenCredential**](KafkaClusterAPIApi.md#GetUserAuthenCredential) | **Get** /vdb-kafka/clusters/{clusterId}/users/{userId}/authen-creds | 
[**GetUserById**](KafkaClusterAPIApi.md#GetUserById) | **Get** /vdb-kafka/clusters/{clusterId}/users/{userId} | 
[**ListClusters**](KafkaClusterAPIApi.md#ListClusters) | **Get** /vdb-kafka/clusters | 
[**ListHistory**](KafkaClusterAPIApi.md#ListHistory) | **Get** /vdb-kafka/clusters/{clusterId}/history | 
[**ListTopic**](KafkaClusterAPIApi.md#ListTopic) | **Get** /vdb-kafka/clusters/{clusterId}/topics | 
[**ListUser**](KafkaClusterAPIApi.md#ListUser) | **Get** /vdb-kafka/clusters/{clusterId}/users | 
[**RegenerateUserAuthenCredential**](KafkaClusterAPIApi.md#RegenerateUserAuthenCredential) | **Put** /vdb-kafka/clusters/{clusterId}/users/{userId}/regenerate-creds | 
[**UpdateAuthentication**](KafkaClusterAPIApi.md#UpdateAuthentication) | **Put** /vdb-kafka/clusters/{clusterId}/authentication | 
[**UpdateBrokerCount**](KafkaClusterAPIApi.md#UpdateBrokerCount) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-broker-count | 
[**UpdateConfigGroup**](KafkaClusterAPIApi.md#UpdateConfigGroup) | **Put** /vdb-kafka/clusters/{clusterId}/config-group | 
[**UpdatePublicAccess**](KafkaClusterAPIApi.md#UpdatePublicAccess) | **Put** /vdb-kafka/clusters/{clusterId}/public-access | 
[**UpdateStorageSize**](KafkaClusterAPIApi.md#UpdateStorageSize) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-storage-size | 
[**UpdateStorageType**](KafkaClusterAPIApi.md#UpdateStorageType) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-storage-type | 
[**UpdateTopic**](KafkaClusterAPIApi.md#UpdateTopic) | **Put** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
[**UpdateUser**](KafkaClusterAPIApi.md#UpdateUser) | **Put** /vdb-kafka/clusters/{clusterId}/users/{userId} | 

# **CreateOrderCluster**
> WrapContentListOrderResponse CreateOrderCluster(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateKafkaClusterRequest**](CreateKafkaClusterRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***KafkaClusterAPIApiCreateOrderClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiCreateOrderClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userType** | **optional.**| ROOT_USER for Checkout flow or IAM_USER for Auto Payment flow. Available values: ROOT_USER, IAM_USER. Default value: ROOT_USER.  | 

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSecRule**
> SecurityGroupRuleDto CreateSecRule(ctx, body, clusterId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecurityGroupRuleCreateRequest**](SecurityGroupRuleCreateRequest.md)|  | 
  **clusterId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**SecurityGroupRuleDto**](SecurityGroupRuleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTopic**
> TopicDto CreateTopic(ctx, body, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicCreateRequest**](TopicCreateRequest.md)|  | 
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**TopicDto**](TopicDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> UserDto CreateUser(ctx, body, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserCreateRequest**](UserCreateRequest.md)|  | 
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCluster**
> string DeleteCluster(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> string DeleteSecRule(ctx, portalUserId, clusterId, secGroupRuleId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> string DeleteTopic(ctx, portalUserId, clusterId, topicId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> string DeleteUser(ctx, portalUserId, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> KafkaCluster GetClusterById(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**KafkaCluster**](KafkaCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopicById**
> TopicDto GetTopicById(ctx, portalUserId, clusterId, topicId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 
  **topicId** | **string**|  | 

### Return type

[**TopicDto**](TopicDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserAuthenCredential**
> []string GetUserAuthenCredential(ctx, portalUserId, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> UserDto GetUserById(ctx, portalUserId, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusters**
> []KafkaCluster ListClusters(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**[]KafkaCluster**](KafkaCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHistory**
> []HistoryDto ListHistory(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**[]HistoryDto**](HistoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTopic**
> []TopicDto ListTopic(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**[]TopicDto**](TopicDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUser**
> []UserDto ListUser(ctx, portalUserId, clusterId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **clusterId** | **string**|  | 

### Return type

[**[]UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateUserAuthenCredential**
> string RegenerateUserAuthenCredential(ctx, portalUserId, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> string UpdateAuthentication(ctx, clusterId, portalUserId, mtlsAuthen, saslAuthen)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **portalUserId** | **int32**|  | 
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
> WrapContentListOrderResponse UpdateBrokerCount(ctx, clusterId, portalUserId, count, rebalance, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **portalUserId** | **int32**|  | 
  **count** | **int32**|  | 
  **rebalance** | **bool**|  | 
 **optional** | ***KafkaClusterAPIApiUpdateBrokerCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiUpdateBrokerCountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userType** | **optional.String**| ROOT_USER for Checkout flow or IAM_USER for Auto Payment flow. Available values: ROOT_USER, IAM_USER. Default value: ROOT_USER.  | 

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfigGroup**
> string UpdateConfigGroup(ctx, clusterId, portalUserId, configGroupVersionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **portalUserId** | **int32**|  | 
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
> string UpdatePublicAccess(ctx, clusterId, portalUserId, enable)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **portalUserId** | **int32**|  | 
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
> WrapContentListOrderResponse UpdateStorageSize(ctx, clusterId, size, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **size** | **int32**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***KafkaClusterAPIApiUpdateStorageSizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiUpdateStorageSizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userType** | **optional.String**| ROOT_USER for Checkout flow or IAM_USER for Auto Payment flow. Available values: ROOT_USER, IAM_USER. Default value: ROOT_USER.  | 

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStorageType**
> WrapContentListOrderResponse UpdateStorageType(ctx, clusterId, storageType, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
  **storageType** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***KafkaClusterAPIApiUpdateStorageTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KafkaClusterAPIApiUpdateStorageTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userType** | **optional.String**| ROOT_USER for Checkout flow or IAM_USER for Auto Payment flow. Available values: ROOT_USER, IAM_USER. Default value: ROOT_USER.  | 

### Return type

[**WrapContentListOrderResponse**](WrapContentListOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTopic**
> string UpdateTopic(ctx, body, portalUserId, clusterId, topicId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicUpdateRequest**](TopicUpdateRequest.md)|  | 
  **portalUserId** | **int32**|  | 
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
> string UpdateUser(ctx, body, portalUserId, clusterId, userId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserUpdatePermissionsRequest**](UserUpdatePermissionsRequest.md)|  | 
  **portalUserId** | **int32**|  | 
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

