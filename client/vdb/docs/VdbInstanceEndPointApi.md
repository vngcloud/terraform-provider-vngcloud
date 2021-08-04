# \VdbInstanceEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitorMetadataUsingPOST**](VdbInstanceEndPointApi.md#AddMonitorMetadataUsingPOST) | **Post** /v1/trv/{projectId}/instances/metadata | addMonitorMetadata
[**CreateDbInstanceUsingPOST**](VdbInstanceEndPointApi.md#CreateDbInstanceUsingPOST) | **Post** /v1/trv/{projectId}/instances | createDbInstance
[**DeleteDbInstanceInTrashUsingDELETE**](VdbInstanceEndPointApi.md#DeleteDbInstanceInTrashUsingDELETE) | **Delete** /v1/trv/{projectId}/instances/{dbInstanceId}/delete-in-trash | deleteDbInstanceInTrash
[**DeleteDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#DeleteDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/delete | deleteDbInstance
[**DetachReplicaUsingPOST**](VdbInstanceEndPointApi.md#DetachReplicaUsingPOST) | **Post** /v1/trv/{projectId}/instances/{dbInstanceId}/detach-replica | detachReplica
[**GetAllAvailableZoneUsingGET**](VdbInstanceEndPointApi.md#GetAllAvailableZoneUsingGET) | **Get** /v1/trv/{projectId}/instances/all-zone | getAllAvailableZone
[**GetDbInstanceDetailFromDatabaseUsingGET**](VdbInstanceEndPointApi.md#GetDbInstanceDetailFromDatabaseUsingGET) | **Get** /v1/trv/{projectId}/instances/from-database/{dbInstanceId} | getDbInstanceDetailFromDatabase
[**GetDbInstanceDetailUsingGET**](VdbInstanceEndPointApi.md#GetDbInstanceDetailUsingGET) | **Get** /v1/trv/{projectId}/instances/{dbInstanceId} | getDbInstanceDetail
[**GetDbInstanceHistoryUsingGET**](VdbInstanceEndPointApi.md#GetDbInstanceHistoryUsingGET) | **Get** /v1/trv/{projectId}/instances/{dbInstanceId}/history | getDbInstanceHistory
[**GetDbInstanceSecurityGroupRuleUsingGET**](VdbInstanceEndPointApi.md#GetDbInstanceSecurityGroupRuleUsingGET) | **Get** /v1/trv/{projectId}/instances/{dbInstanceId}/security-group-rule | getDbInstanceSecurityGroupRule
[**GetListDbInstanceByEngineGroupUsingGET**](VdbInstanceEndPointApi.md#GetListDbInstanceByEngineGroupUsingGET) | **Get** /v1/trv/{projectId}/instances | getListDbInstanceByEngineGroup
[**GetListPartialDbInstanceUsingGET**](VdbInstanceEndPointApi.md#GetListPartialDbInstanceUsingGET) | **Get** /v1/trv/{projectId}/instances/partial | getListPartialDbInstance
[**RecoverDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#RecoverDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/recover | recoverDbInstance
[**RemoveMonitorMetadataUsingDELETE**](VdbInstanceEndPointApi.md#RemoveMonitorMetadataUsingDELETE) | **Delete** /v1/trv/{projectId}/instances/metadata | removeMonitorMetadata
[**RenewDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#RenewDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/renew | renewDbInstance
[**RestartDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#RestartDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/restart | restartDbInstance
[**StartDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#StartDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/start | startDbInstance
[**StopDbInstancePocUsingPOST**](VdbInstanceEndPointApi.md#StopDbInstancePocUsingPOST) | **Post** /v1/trv/{projectId}/instances/{dbInstanceId}/poc/stop | stopDbInstancePoc
[**StopDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#StopDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/stop | stopDbInstance
[**UpdateDbInstanceAutoRenewUsingPUT**](VdbInstanceEndPointApi.md#UpdateDbInstanceAutoRenewUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/auto-renew | updateDbInstanceAutoRenew
[**UpdateDbInstanceSecurityGroupRuleUsingPUT**](VdbInstanceEndPointApi.md#UpdateDbInstanceSecurityGroupRuleUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/security-group-rule | updateDbInstanceSecurityGroupRule
[**UpdateDbInstanceUsingPUT**](VdbInstanceEndPointApi.md#UpdateDbInstanceUsingPUT) | **Put** /v1/trv/{projectId}/instances/{dbInstanceId}/update | updateDbInstance


# **AddMonitorMetadataUsingPOST**
> BaseResponse AddMonitorMetadataUsingPOST(ctx, metadataRequest, projectId)
addMonitorMetadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metadataRequest** | [**MetadataRequest**](MetadataRequest.md)| metadataRequest | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDbInstanceUsingPOST**
> DbInstanceResponse CreateDbInstanceUsingPOST(ctx, createDbInstanceRequest, projectId)
createDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createDbInstanceRequest** | [**CreateDbInstanceRequest**](CreateDbInstanceRequest.md)| createDbInstanceRequest | 
  **projectId** | **string**| projectId | 

### Return type

[**DbInstanceResponse**](DbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDbInstanceInTrashUsingDELETE**
> DbInstanceResponse DeleteDbInstanceInTrashUsingDELETE(ctx, dbInstanceId, deleteDBInstanceInTrashRequest, projectId)
deleteDbInstanceInTrash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **deleteDBInstanceInTrashRequest** | [**DeleteDbInstanceInTrashRequest**](DeleteDbInstanceInTrashRequest.md)| deleteDBInstanceInTrashRequest | 
  **projectId** | **string**| projectId | 

### Return type

[**DbInstanceResponse**](DbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDbInstanceUsingPUT**
> DbInstanceResponse DeleteDbInstanceUsingPUT(ctx, dbInstanceId, deleteDBInstanceRequest, projectId)
deleteDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **deleteDBInstanceRequest** | [**DeleteDbInstanceRequest**](DeleteDbInstanceRequest.md)| deleteDBInstanceRequest | 
  **projectId** | **string**| projectId | 

### Return type

[**DbInstanceResponse**](DbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachReplicaUsingPOST**
> BaseResponse DetachReplicaUsingPOST(ctx, dbInstanceId, projectId)
detachReplica

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAvailableZoneUsingGET**
> BaseResponse GetAllAvailableZoneUsingGET(ctx, projectId)
getAllAvailableZone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDbInstanceDetailFromDatabaseUsingGET**
> BaseResponse GetDbInstanceDetailFromDatabaseUsingGET(ctx, dbInstanceId, projectId)
getDbInstanceDetailFromDatabase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDbInstanceDetailUsingGET**
> DbInstanceResponse GetDbInstanceDetailUsingGET(ctx, dbInstanceId, projectId)
getDbInstanceDetail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**DbInstanceResponse**](DbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDbInstanceHistoryUsingGET**
> BaseResponse GetDbInstanceHistoryUsingGET(ctx, dbInstanceId, projectId)
getDbInstanceHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDbInstanceSecurityGroupRuleUsingGET**
> DbSecurityGroupRuleResponse GetDbInstanceSecurityGroupRuleUsingGET(ctx, dbInstanceId, projectId)
getDbInstanceSecurityGroupRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**DbSecurityGroupRuleResponse**](DbSecurityGroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListDbInstanceByEngineGroupUsingGET**
> ListDbInstanceResponse GetListDbInstanceByEngineGroupUsingGET(ctx, engineGroup, projectId)
getListDbInstanceByEngineGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineGroup** | **int32**| engineGroup | 
  **projectId** | **string**| projectId | 

### Return type

[**ListDbInstanceResponse**](ListDbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListPartialDbInstanceUsingGET**
> BaseResponse GetListPartialDbInstanceUsingGET(ctx, projectId)
getListPartialDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecoverDbInstanceUsingPUT**
> BaseResponse RecoverDbInstanceUsingPUT(ctx, dbInstanceId, projectId, recoverDBRequest)
recoverDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 
  **recoverDBRequest** | [**RecoverDbRequest**](RecoverDbRequest.md)| recoverDBRequest | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMonitorMetadataUsingDELETE**
> BaseResponse RemoveMonitorMetadataUsingDELETE(ctx, metadataRequest, projectId)
removeMonitorMetadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metadataRequest** | [**MetadataRequest**](MetadataRequest.md)| metadataRequest | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewDbInstanceUsingPUT**
> BaseResponse RenewDbInstanceUsingPUT(ctx, dbInstanceId, projectId, updateDbInstanceRequest)
renewDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 
  **updateDbInstanceRequest** | [**UpdateDbInstanceRequest**](UpdateDbInstanceRequest.md)| updateDbInstanceRequest | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartDbInstanceUsingPUT**
> BaseResponse RestartDbInstanceUsingPUT(ctx, dbInstanceId, projectId)
restartDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDbInstanceUsingPUT**
> BaseResponse StartDbInstanceUsingPUT(ctx, dbInstanceId, projectId)
startDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopDbInstancePocUsingPOST**
> BaseResponse StopDbInstancePocUsingPOST(ctx, dbInstanceId, projectId, request)
stopDbInstancePoc

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 
  **request** | [**StopDbPocRequest**](StopDbPocRequest.md)| request | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopDbInstanceUsingPUT**
> BaseResponse StopDbInstanceUsingPUT(ctx, dbInstanceId, projectId)
stopDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDbInstanceAutoRenewUsingPUT**
> BaseResponse UpdateDbInstanceAutoRenewUsingPUT(ctx, dbAutoRenewChangeRequest, dbInstanceId, projectId)
updateDbInstanceAutoRenew

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbAutoRenewChangeRequest** | [**DbAutoRenewChangeRequest**](DbAutoRenewChangeRequest.md)| dbAutoRenewChangeRequest | 
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDbInstanceSecurityGroupRuleUsingPUT**
> DbSecurityGroupRuleResponse UpdateDbInstanceSecurityGroupRuleUsingPUT(ctx, dbInstanceId, projectId, rules)
updateDbInstanceSecurityGroupRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 
  **rules** | [**[]SecurityGroupRuleInfo**](SecurityGroupRuleInfo.md)| rules | 

### Return type

[**DbSecurityGroupRuleResponse**](DbSecurityGroupRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDbInstanceUsingPUT**
> DbInstanceResponse UpdateDbInstanceUsingPUT(ctx, dbInstanceId, projectId, updateDbInstanceRequest)
updateDbInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbInstanceId** | **string**| dbInstanceId | 
  **projectId** | **string**| projectId | 
  **updateDbInstanceRequest** | [**UpdateDbInstanceRequest**](UpdateDbInstanceRequest.md)| updateDbInstanceRequest | 

### Return type

[**DbInstanceResponse**](DbInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

