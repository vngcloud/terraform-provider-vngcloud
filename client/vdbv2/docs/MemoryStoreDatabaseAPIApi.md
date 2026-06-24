# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseInstanceReplica**](MemoryStoreDatabaseAPIApi.md#CreateDatabaseInstanceReplica) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/create-replicas | 
[**CreateMemoryStoreDatabaseInstance**](MemoryStoreDatabaseAPIApi.md#CreateMemoryStoreDatabaseInstance) | **Post** /vdb-memory/v1/payment/database-instances | 
[**DeleteDatabaseInstances1**](MemoryStoreDatabaseAPIApi.md#DeleteDatabaseInstances1) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/delete | 
[**DetachReplica1**](MemoryStoreDatabaseAPIApi.md#DetachReplica1) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/detach-replica | 
[**GetDatabaseInstancesById1**](MemoryStoreDatabaseAPIApi.md#GetDatabaseInstancesById1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId} | 
[**GetDatabaseInstancesByUser1**](MemoryStoreDatabaseAPIApi.md#GetDatabaseInstancesByUser1) | **Get** /vdb-memory/v1/database-instances | 
[**GetHistoryDB1**](MemoryStoreDatabaseAPIApi.md#GetHistoryDB1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/histories | 
[**GetListBackupsByInstanceId1**](MemoryStoreDatabaseAPIApi.md#GetListBackupsByInstanceId1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/backups | 
[**GetListReplicas1**](MemoryStoreDatabaseAPIApi.md#GetListReplicas1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/replicas | 
[**GetSecurityRules1**](MemoryStoreDatabaseAPIApi.md#GetSecurityRules1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/secrules | 
[**ResizeInstance1**](MemoryStoreDatabaseAPIApi.md#ResizeInstance1) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/resize-instance | 
[**RestartDatabaseInstances1**](MemoryStoreDatabaseAPIApi.md#RestartDatabaseInstances1) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/reboot | 
[**StartDatabaseInstances1**](MemoryStoreDatabaseAPIApi.md#StartDatabaseInstances1) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/start | 
[**StopDatabaseInstances1**](MemoryStoreDatabaseAPIApi.md#StopDatabaseInstances1) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/shutdown | 
[**UpdateDatabaseConfigGroup1**](MemoryStoreDatabaseAPIApi.md#UpdateDatabaseConfigGroup1) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/update-config-group | 
[**UpdateDatabaseSetting1**](MemoryStoreDatabaseAPIApi.md#UpdateDatabaseSetting1) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/update-setting | 
[**UpdateSecurityRules**](MemoryStoreDatabaseAPIApi.md#UpdateSecurityRules) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/secrules | 

# **CreateDatabaseInstanceReplica**
> WrapContentListOrderResponse CreateDatabaseInstanceReplica(ctx, body, dbInstanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMemDbInstanceReplicaRequest**](CreateMemDbInstanceReplicaRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiCreateDatabaseInstanceReplicaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiCreateDatabaseInstanceReplicaOpts struct
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

# **CreateMemoryStoreDatabaseInstance**
> WrapContentListOrderResponse CreateMemoryStoreDatabaseInstance(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMemDbInstanceRequest**](CreateMemDbInstanceRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiCreateMemoryStoreDatabaseInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiCreateMemoryStoreDatabaseInstanceOpts struct
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

# **DeleteDatabaseInstances1**
> WrapContentListActionDbInstancesResponse DeleteDatabaseInstances1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteDbInstanceRequest**](DeleteDbInstanceRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachReplica1**
> WrapContentListActionDbInstancesResponse DetachReplica1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseInstancesById1**
> WrapContentDbInstanceInfo GetDatabaseInstancesById1(ctx, portalUserId, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentDbInstanceInfo**](WrapContentDbInstanceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseInstancesByUser1**
> WrapContentDatabaseInstancesGatewayResponse GetDatabaseInstancesByUser1(ctx, filterRequest, pageNumber, pageSize)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filterRequest** | [**FilterRequest1**](.md)|  | 
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

# **GetHistoryDB1**
> WrapContentDbInstancesHistoryGatewayResponse GetHistoryDB1(ctx, dbInstanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiGetHistoryDB1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiGetHistoryDB1Opts struct
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

# **GetListBackupsByInstanceId1**
> WrapContentListBackupInfo GetListBackupsByInstanceId1(ctx, portalUserId, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListBackupInfo**](WrapContentListBackupInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListReplicas1**
> WrapContentObject GetListReplicas1(ctx, portalUserId, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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
> WrapContentListSecurityGroupRuleEntity GetSecurityRules1(ctx, portalUserId, dbInstanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **dbInstanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeInstance1**
> WrapContentListOrderResponse ResizeInstance1(ctx, body, dbInstanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResizeFlavorRequest**](ResizeFlavorRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreDatabaseAPIApiResizeInstance1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreDatabaseAPIApiResizeInstance1Opts struct
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

# **RestartDatabaseInstances1**
> WrapContentListActionDbInstancesResponse RestartDatabaseInstances1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDatabaseInstances1**
> WrapContentListActionDbInstancesResponse StartDatabaseInstances1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopDatabaseInstances1**
> WrapContentListActionDbInstancesResponse StopDatabaseInstances1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseConfigGroup1**
> WrapContentObject UpdateDatabaseConfigGroup1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDbConfigGroupRequest**](UpdateDbConfigGroupRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseSetting1**
> WrapContentObject UpdateDatabaseSetting1(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateMemDbSettingRequest**](UpdateMemDbSettingRequest.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityRules**
> WrapContentListSecurityGroupRuleEntity UpdateSecurityRules(ctx, body, dbInstanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateSecurityGroupRuleDetail**](UpdateSecurityGroupRuleDetail.md)|  | 
  **dbInstanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

