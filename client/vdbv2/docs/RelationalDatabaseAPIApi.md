# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelationalDatabaseInstance**](RelationalDatabaseAPIApi.md#CreateRelationalDatabaseInstance) | **Post** /vdb-relational/v1/payment/database-instances | 
[**CreateRelationalDatabaseInstanceReplica**](RelationalDatabaseAPIApi.md#CreateRelationalDatabaseInstanceReplica) | **Post** /vdb-relational/v1/database-instances/{instanceId}/create-replicas | 
[**DeleteDatabaseInstances**](RelationalDatabaseAPIApi.md#DeleteDatabaseInstances) | **Post** /vdb-relational/v1/database-instances/{instanceId}/delete | 
[**DetachReplica**](RelationalDatabaseAPIApi.md#DetachReplica) | **Post** /vdb-relational/v1/database-instances/{instanceId}/detach-replica | 
[**GetAllDatastore**](RelationalDatabaseAPIApi.md#GetAllDatastore) | **Get** /vdb-relational/v1/database-instances/datastore | 
[**GetAllInstanceFamily**](RelationalDatabaseAPIApi.md#GetAllInstanceFamily) | **Get** /vdb-relational/v1/database-instances/families | 
[**GetDBInstanceConfig**](RelationalDatabaseAPIApi.md#GetDBInstanceConfig) | **Get** /vdb-relational/v1/database-instances/configuration | 
[**GetDatabaseInstancesById**](RelationalDatabaseAPIApi.md#GetDatabaseInstancesById) | **Get** /vdb-relational/v1/database-instances/id/{dbInstanceId} | 
[**GetDatabaseInstancesByUser**](RelationalDatabaseAPIApi.md#GetDatabaseInstancesByUser) | **Get** /vdb-relational/v1/database-instances | 
[**GetEngine**](RelationalDatabaseAPIApi.md#GetEngine) | **Get** /vdb-relational/v1/database-instances/engine | 
[**GetFlavorCodes**](RelationalDatabaseAPIApi.md#GetFlavorCodes) | **Get** /vdb-relational/v1/database-instances/flavor_zones/codes | 
[**GetFlavors**](RelationalDatabaseAPIApi.md#GetFlavors) | **Get** /vdb-relational/v1/database-instances/flavors | 
[**GetHistoryDB**](RelationalDatabaseAPIApi.md#GetHistoryDB) | **Get** /vdb-relational/v1/database-instances/{instanceId}/histories | 
[**GetListNetwork**](RelationalDatabaseAPIApi.md#GetListNetwork) | **Get** /vdb-relational/v1/database-instances/networks | 
[**GetListReplicas**](RelationalDatabaseAPIApi.md#GetListReplicas) | **Get** /vdb-relational/v1/database-instances/{replicaSourceId}/replicas | 
[**GetSecurityRules**](RelationalDatabaseAPIApi.md#GetSecurityRules) | **Get** /vdb-relational/v1/database-instances/{instanceId}/secrules | 
[**GetVolumeTypes**](RelationalDatabaseAPIApi.md#GetVolumeTypes) | **Get** /vdb-relational/v1/database-instances/volume/types | 
[**ListSubnet**](RelationalDatabaseAPIApi.md#ListSubnet) | **Get** /vdb-relational/v1/database-instances/networks/subnets | 
[**ListZone**](RelationalDatabaseAPIApi.md#ListZone) | **Get** /vdb-relational/v1/database-instances/zones | 
[**ResizeInstance**](RelationalDatabaseAPIApi.md#ResizeInstance) | **Post** /vdb-relational/v1/database-instances/{instanceId}/resize-instance | 
[**ResizeStorage**](RelationalDatabaseAPIApi.md#ResizeStorage) | **Post** /vdb-relational/v1/database-instances/{instanceId}/resize-storage | 
[**RestartDatabaseInstances**](RelationalDatabaseAPIApi.md#RestartDatabaseInstances) | **Post** /vdb-relational/v1/database-instances/{instanceId}/reboot | 
[**StartDatabaseInstances**](RelationalDatabaseAPIApi.md#StartDatabaseInstances) | **Post** /vdb-relational/v1/database-instances/{instanceId}/start | 
[**StopDatabaseInstances**](RelationalDatabaseAPIApi.md#StopDatabaseInstances) | **Post** /vdb-relational/v1/database-instances/{instanceId}/shutdown | 
[**UpdateDatabaseConfigGroup**](RelationalDatabaseAPIApi.md#UpdateDatabaseConfigGroup) | **Put** /vdb-relational/v1/database-instances/{instanceId}/update/config-group | 
[**UpdateDatabaseSetting**](RelationalDatabaseAPIApi.md#UpdateDatabaseSetting) | **Put** /vdb-relational/v1/database-instances/{instanceId}/update/setting | 
[**UpdateSecurityGroupRules**](RelationalDatabaseAPIApi.md#UpdateSecurityGroupRules) | **Put** /vdb-relational/v1/database-instances/{instanceId}/secrules | 

# **CreateRelationalDatabaseInstance**
> WrapContentListOrderResponse CreateRelationalDatabaseInstance(ctx, body, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDbInstanceRequest**](CreateDbInstanceRequest.md)|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceOpts struct
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

# **CreateRelationalDatabaseInstanceReplica**
> WrapContentListOrderResponse CreateRelationalDatabaseInstanceReplica(ctx, body, instanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDbInstanceReplicaRequest**](CreateDbInstanceReplicaRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceReplicaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceReplicaOpts struct
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

# **DeleteDatabaseInstances**
> WrapContentListActionDbInstancesResponse DeleteDatabaseInstances(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteDbInstanceRequest**](DeleteDbInstanceRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachReplica**
> WrapContentListActionDbInstancesResponse DetachReplica(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllDatastore**
> WrapContentListEngineVersion GetAllDatastore(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListEngineVersion**](WrapContentListEngineVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInstanceFamily**
> WrapContentListInstanceFamily GetAllInstanceFamily(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListInstanceFamily**](WrapContentListInstanceFamily.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDBInstanceConfig**
> WrapContentListItemConfigInfo GetDBInstanceConfig(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListItemConfigInfo**](WrapContentListItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseInstancesById**
> WrapContentDbInstanceInfo GetDatabaseInstancesById(ctx, portalUserId, dbInstanceId)


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

# **GetDatabaseInstancesByUser**
> WrapContentDatabaseInstancesGatewayResponse GetDatabaseInstancesByUser(ctx, portalUserId, filterRequest, pageNumber, pageSize)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
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

# **GetEngine**
> WrapContentSetEngine GetEngine(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentSetEngine**](WrapContentSetEngine.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlavorCodes**
> WrapContentSetFlavorCode GetFlavorCodes(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentSetFlavorCode**](WrapContentSetFlavorCode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlavors**
> WrapContentListFlavorInfo GetFlavors(ctx, type_, version, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **version** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiGetFlavorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiGetFlavorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **zoneId** | **optional.String**|  | [default to HCM03-1A]

### Return type

[**WrapContentListFlavorInfo**](WrapContentListFlavorInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryDB**
> WrapContentDbInstancesHistoryGatewayResponse GetHistoryDB(ctx, instanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiGetHistoryDBOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiGetHistoryDBOpts struct
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

# **GetListNetwork**
> WrapContentListNetworkResponse GetListNetwork(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListNetworkResponse**](WrapContentListNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListReplicas**
> WrapContentObject GetListReplicas(ctx, portalUserId, replicaSourceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **replicaSourceId** | **string**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityRules**
> WrapContentListSecurityGroupRuleEntity GetSecurityRules(ctx, portalUserId, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeTypes**
> WrapContentVolumeTypeGatewayResponse GetVolumeTypes(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiGetVolumeTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiGetVolumeTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneId** | **optional.String**|  | [default to HCM03-1A]

### Return type

[**WrapContentVolumeTypeGatewayResponse**](WrapContentVolumeTypeGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubnet**
> WrapContentListNetworkResponseV2 ListSubnet(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiListSubnetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiListSubnetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneId** | **optional.String**|  | 

### Return type

[**WrapContentListNetworkResponseV2**](WrapContentListNetworkResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListZone**
> WrapContentListZoneInfo ListZone(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListZoneInfo**](WrapContentListZoneInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeInstance**
> WrapContentListOrderResponse ResizeInstance(ctx, body, instanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResizeFlavorRequest**](ResizeFlavorRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiResizeInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiResizeInstanceOpts struct
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

# **ResizeStorage**
> WrapContentListOrderResponse ResizeStorage(ctx, body, instanceId, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResizeVolumeRequest**](ResizeVolumeRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***RelationalDatabaseAPIApiResizeStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiResizeStorageOpts struct
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

# **RestartDatabaseInstances**
> WrapContentListActionDbInstancesResponse RestartDatabaseInstances(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDatabaseInstances**
> WrapContentListActionDbInstancesResponse StartDatabaseInstances(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopDatabaseInstances**
> WrapContentListActionDbInstancesResponse StopDatabaseInstances(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActionDbInstanceRequest**](ActionDbInstanceRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseConfigGroup**
> WrapContentObject UpdateDatabaseConfigGroup(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDbConfigGroupRequest**](UpdateDbConfigGroupRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseSetting**
> WrapContentObject UpdateDatabaseSetting(ctx, body, instanceId, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDbSettingRequest**](UpdateDbSettingRequest.md)|  | 
  **instanceId** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecurityGroupRules**
> WrapContentListSecurityGroupRuleEntity UpdateSecurityGroupRules(ctx, body, portalUserId, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]UpdateSecurityGroupRuleDetail**](UpdateSecurityGroupRuleDetail.md)|  | 
  **portalUserId** | **int32**|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

