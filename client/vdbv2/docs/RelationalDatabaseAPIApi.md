# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelationalDatabaseInstance**](RelationalDatabaseAPIApi.md#CreateRelationalDatabaseInstance) | **Post** /v1/payment/database-instances | 
[**CreateRelationalDatabaseInstanceReplica**](RelationalDatabaseAPIApi.md#CreateRelationalDatabaseInstanceReplica) | **Post** /v1/database-instances/{id}/create-replicas | 
[**DeleteDatabaseInstances1**](RelationalDatabaseAPIApi.md#DeleteDatabaseInstances1) | **Post** /v1/database-instances/{instanceId}/delete | 
[**DetachReplica1**](RelationalDatabaseAPIApi.md#DetachReplica1) | **Post** /v1/database-instances/{instanceId}/detach-replica | 
[**ExchangePocResource1**](RelationalDatabaseAPIApi.md#ExchangePocResource1) | **Post** /v1/database-instances/exchange-poc | 
[**GetAllDatastore1**](RelationalDatabaseAPIApi.md#GetAllDatastore1) | **Get** /v1/database-instances/datastore | 
[**GetAllInstanceFamily1**](RelationalDatabaseAPIApi.md#GetAllInstanceFamily1) | **Get** /v1/database-instances/families | 
[**GetDBInstanceConfig1**](RelationalDatabaseAPIApi.md#GetDBInstanceConfig1) | **Get** /v1/database-instances/configuration | 
[**GetDatabaseInstancesById1**](RelationalDatabaseAPIApi.md#GetDatabaseInstancesById1) | **Get** /v1/database-instances/id/{dbInstanceId} | 
[**GetDatabaseInstancesByUser1**](RelationalDatabaseAPIApi.md#GetDatabaseInstancesByUser1) | **Get** /v1/database-instances | 
[**GetEngine1**](RelationalDatabaseAPIApi.md#GetEngine1) | **Get** /v1/database-instances/engine | 
[**GetFlavorCodes1**](RelationalDatabaseAPIApi.md#GetFlavorCodes1) | **Get** /v1/database-instances/flavor_zones/codes | 
[**GetFlavors1**](RelationalDatabaseAPIApi.md#GetFlavors1) | **Get** /v1/database-instances/flavors | 
[**GetHistoryDB1**](RelationalDatabaseAPIApi.md#GetHistoryDB1) | **Get** /v1/database-instances/{instanceId}/histories | 
[**GetListNetwork1**](RelationalDatabaseAPIApi.md#GetListNetwork1) | **Get** /v1/database-instances/networks | 
[**GetListReplicas1**](RelationalDatabaseAPIApi.md#GetListReplicas1) | **Get** /v1/database-instances/{replicaSourceId}/replicas | 
[**GetSecurityRules**](RelationalDatabaseAPIApi.md#GetSecurityRules) | **Put** /v1/database-instances/{instanceId}/secrules | 
[**GetSecurityRules2**](RelationalDatabaseAPIApi.md#GetSecurityRules2) | **Get** /v1/database-instances/{instanceId}/secrules | 
[**GetVolumeTypes2**](RelationalDatabaseAPIApi.md#GetVolumeTypes2) | **Get** /v1/database-instances/volume/types | 
[**ListNetwork1**](RelationalDatabaseAPIApi.md#ListNetwork1) | **Get** /v1/database-instances/networks/subnets | 
[**RenewResource1**](RelationalDatabaseAPIApi.md#RenewResource1) | **Post** /v1/database-instances/{resourceId}/renew-resource | 
[**ResizeInstance1**](RelationalDatabaseAPIApi.md#ResizeInstance1) | **Post** /v1/database-instances/{instanceId}/resize-instance | 
[**ResizeStorage**](RelationalDatabaseAPIApi.md#ResizeStorage) | **Post** /v1/database-instances/{instanceId}/resize-storage | 
[**RestartDatabaseInstances1**](RelationalDatabaseAPIApi.md#RestartDatabaseInstances1) | **Post** /v1/database-instances/{instanceId}/reboot | 
[**StartDatabaseInstances1**](RelationalDatabaseAPIApi.md#StartDatabaseInstances1) | **Post** /v1/database-instances/{instanceId}/start | 
[**StopDatabaseInstances1**](RelationalDatabaseAPIApi.md#StopDatabaseInstances1) | **Post** /v1/database-instances/{instanceId}/shutdown | 
[**UpdateDatabaseConfigGroup1**](RelationalDatabaseAPIApi.md#UpdateDatabaseConfigGroup1) | **Put** /v1/database-instances/{instanceId}/update/config-group | 
[**UpdateDatabaseSetting1**](RelationalDatabaseAPIApi.md#UpdateDatabaseSetting1) | **Put** /v1/database-instances/{instanceId}/update/setting | 

# **CreateRelationalDatabaseInstance**
> WrapContentListOrderResponse CreateRelationalDatabaseInstance(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceOpts struct
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

# **CreateRelationalDatabaseInstanceReplica**
> WrapContentListOrderResponse CreateRelationalDatabaseInstanceReplica(ctx, body, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **id** | **string**|  | 
 **optional** | ***RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceReplicaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiCreateRelationalDatabaseInstanceReplicaOpts struct
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

# **DeleteDatabaseInstances1**
> WrapContentListActionDbInstancesResponse DeleteDatabaseInstances1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachReplica1**
> WrapContentListActionDbInstancesResponse DetachReplica1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExchangePocResource1**
> WrapContentListOrderResponse ExchangePocResource1(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
 **optional** | ***RelationalDatabaseAPIApiExchangePocResource1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiExchangePocResource1Opts struct
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

# **GetAllDatastore1**
> WrapContentListEngineVersion GetAllDatastore1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListEngineVersion**](WrapContentListEngineVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInstanceFamily1**
> WrapContentListInstanceFamily GetAllInstanceFamily1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListInstanceFamily**](WrapContentListInstanceFamily.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDBInstanceConfig1**
> WrapContentListItemConfigInfo GetDBInstanceConfig1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListItemConfigInfo**](WrapContentListItemConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseInstancesById1**
> WrapContentDbInstanceInfo GetDatabaseInstancesById1(ctx, dbInstanceId)


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

# **GetDatabaseInstancesByUser1**
> WrapContentDatabaseInstancesGatewayResponse GetDatabaseInstancesByUser1(ctx, filterRequest, pageNumber, pageSize)


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

# **GetEngine1**
> WrapContentSetEngine GetEngine1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentSetEngine**](WrapContentSetEngine.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlavorCodes1**
> WrapContentSetFlavorCode GetFlavorCodes1(ctx, portalUserId)


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

# **GetFlavors1**
> WrapContentListFlavorInfo GetFlavors1(ctx, type_, version, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **version** | **string**|  | 
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentListFlavorInfo**](WrapContentListFlavorInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryDB1**
> WrapContentDbInstancesHistoryGatewayResponse GetHistoryDB1(ctx, instanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **string**|  | 
 **optional** | ***RelationalDatabaseAPIApiGetHistoryDB1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiGetHistoryDB1Opts struct
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

# **GetListNetwork1**
> WrapContentListNetworkResponse GetListNetwork1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListNetworkResponse**](WrapContentListNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListReplicas1**
> WrapContentObject GetListReplicas1(ctx, replicaSourceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
> WrapContentListSecurityGroupRuleEntity GetSecurityRules(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]SecurityGroupRuleEntity**](SecurityGroupRuleEntity.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurityRules2**
> WrapContentListSecurityGroupRuleEntity GetSecurityRules2(ctx, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListSecurityGroupRuleEntity**](WrapContentListSecurityGroupRuleEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVolumeTypes2**
> WrapContentVolumeTypeGatewayResponse GetVolumeTypes2(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentVolumeTypeGatewayResponse**](WrapContentVolumeTypeGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetwork1**
> WrapContentListNetworkResponseV2 ListNetwork1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentListNetworkResponseV2**](WrapContentListNetworkResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewResource1**
> WrapContentListOrderResponse RenewResource1(ctx, body, resourceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **resourceId** | **string**|  | 
 **optional** | ***RelationalDatabaseAPIApiRenewResource1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiRenewResource1Opts struct
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

# **ResizeInstance1**
> WrapContentListOrderResponse ResizeInstance1(ctx, body, instanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 
 **optional** | ***RelationalDatabaseAPIApiResizeInstance1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiResizeInstance1Opts struct
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

# **ResizeStorage**
> WrapContentListOrderResponse ResizeStorage(ctx, body, instanceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 
 **optional** | ***RelationalDatabaseAPIApiResizeStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelationalDatabaseAPIApiResizeStorageOpts struct
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

# **RestartDatabaseInstances1**
> WrapContentListActionDbInstancesResponse RestartDatabaseInstances1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDatabaseInstances1**
> WrapContentListActionDbInstancesResponse StartDatabaseInstances1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopDatabaseInstances1**
> WrapContentListActionDbInstancesResponse StopDatabaseInstances1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentListActionDbInstancesResponse**](WrapContentListActionDbInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseConfigGroup1**
> WrapContentObject UpdateDatabaseConfigGroup1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatabaseSetting1**
> WrapContentObject UpdateDatabaseSetting1(ctx, body, instanceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **instanceId** | **string**|  | 

### Return type

[**WrapContentObject**](WrapContentObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

