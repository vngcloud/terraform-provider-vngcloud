# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseInstanceReplica**](MemoryStoreDatabaseAPIApi.md#CreateDatabaseInstanceReplica) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/create-replicas | 
[**CreateMemoryStoreDatabaseInstance**](MemoryStoreDatabaseAPIApi.md#CreateMemoryStoreDatabaseInstance) | **Post** /vdb-memory/v1/payment/database-instances | 
[**DeleteDatabaseInstances**](MemoryStoreDatabaseAPIApi.md#DeleteDatabaseInstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/delete | 
[**DetachReplica**](MemoryStoreDatabaseAPIApi.md#DetachReplica) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/detach-replica | 
[**ExchangePocResource**](MemoryStoreDatabaseAPIApi.md#ExchangePocResource) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/exchange-poc | 
[**GetDatabaseInstancesById**](MemoryStoreDatabaseAPIApi.md#GetDatabaseInstancesById) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId} | 
[**GetDatabaseInstancesByUser**](MemoryStoreDatabaseAPIApi.md#GetDatabaseInstancesByUser) | **Get** /vdb-memory/v1/database-instances | 
[**GetHistoryDB**](MemoryStoreDatabaseAPIApi.md#GetHistoryDB) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/histories | 
[**GetListBackupsByInstanceId**](MemoryStoreDatabaseAPIApi.md#GetListBackupsByInstanceId) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/backups | 
[**GetListReplicas**](MemoryStoreDatabaseAPIApi.md#GetListReplicas) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/replicas | 
[**GetSecurityRules1**](MemoryStoreDatabaseAPIApi.md#GetSecurityRules1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/secrules | 
[**RenewResource**](MemoryStoreDatabaseAPIApi.md#RenewResource) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/renew-resource | 
[**ResizeInstance**](MemoryStoreDatabaseAPIApi.md#ResizeInstance) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/resize-instance | 
[**RestartDatabaseInstances**](MemoryStoreDatabaseAPIApi.md#RestartDatabaseInstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/reboot | 
[**StartDatabaseInstances**](MemoryStoreDatabaseAPIApi.md#StartDatabaseInstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/start | 
[**StopDatabaseInstances**](MemoryStoreDatabaseAPIApi.md#StopDatabaseInstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/shutdown | 
[**UpdateDatabaseConfigGroup**](MemoryStoreDatabaseAPIApi.md#UpdateDatabaseConfigGroup) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/update-config-group | 
[**UpdateDatabaseSetting**](MemoryStoreDatabaseAPIApi.md#UpdateDatabaseSetting) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/update-setting | 
[**UpdateSecurityRules**](MemoryStoreDatabaseAPIApi.md#UpdateSecurityRules) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/secrules | 

# **CreateDatabaseInstanceReplica**
> WrapContentListOrderResponse CreateDatabaseInstanceReplica(ctx, body, dbInstanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiCreateDatabaseInstanceReplicaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiCreateDatabaseInstanceReplicaOpts struct
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

# **CreateMemoryStoreDatabaseInstance**
> WrapContentListOrderResponse CreateMemoryStoreDatabaseInstance(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiCreateMemoryStoreDatabaseInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiCreateMemoryStoreDatabaseInstanceOpts struct
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

# **DeleteDatabaseInstances**
> WrapContentListActionDbInstancesResponse DeleteDatabaseInstances(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachReplica**
> WrapContentListActionDbInstancesResponse DetachReplica(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExchangePocResource**
> WrapContentListOrderResponse ExchangePocResource(ctx, body, dbInstanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiExchangePocResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiExchangePocResourceOpts struct
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

# **GetDatabaseInstancesById**
> WrapContentDbInstanceInfo GetDatabaseInstancesById(ctx, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentDbInstanceInfo**](WrapContentDbInstanceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseInstancesByUser**
> WrapContentDatabaseInstancesGatewayResponse GetDatabaseInstancesByUser(ctx, filterRequest, pageNumber, pageSize)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterRequest** | [**FilterRequest**](.md)|  | 
  **pageNumber** | **int32**|  | 
  **pageSize** | **int32**|  | 

### Return type

[**WrapContentDatabaseInstancesGatewayResponse**](WrapContentDatabaseInstancesGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryDB**
> WrapContentDbInstancesHistoryGatewayResponse GetHistoryDB(ctx, dbInstanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiGetHistoryDBOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiGetHistoryDBOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 10]

### Return type

[**WrapContentDbInstancesHistoryGatewayResponse**](WrapContentDBInstancesHistoryGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListBackupsByInstanceId**
> WrapContentListBackupInfo GetListBackupsByInstanceId(ctx, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListBackupInfo**](WrapContentListBackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListReplicas**
> WrapContentObject GetListReplicas(ctx, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityRules1**
> WrapContentListSecurityGroupRuleEntity GetSecurityRules1(ctx, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewResource**
> WrapContentListOrderResponse RenewResource(ctx, body, dbInstanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiRenewResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiRenewResourceOpts struct
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

# **ResizeInstance**
> WrapContentListOrderResponse ResizeInstance(ctx, body, dbInstanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiResizeInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiResizeInstanceOpts struct
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

# **RestartDatabaseInstances**
> WrapContentListActionDbInstancesResponse RestartDatabaseInstances(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDatabaseInstances**
> WrapContentListActionDbInstancesResponse StartDatabaseInstances(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopDatabaseInstances**
> WrapContentListActionDbInstancesResponse StopDatabaseInstances(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseConfigGroup**
> WrapContentObject UpdateDatabaseConfigGroup(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseSetting**
> WrapContentObject UpdateDatabaseSetting(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityRules**
> WrapContentListSecurityGroupRuleEntity UpdateSecurityRules(ctx, body, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

