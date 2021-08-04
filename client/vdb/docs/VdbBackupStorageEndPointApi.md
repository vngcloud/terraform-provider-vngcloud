# \VdbBackupStorageEndPointApi

All URIs are relative to *https://localhost:8101*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupStorageUsingPOST**](VdbBackupStorageEndPointApi.md#CreateBackupStorageUsingPOST) | **Post** /v1/trv/{projectId}/{engineGroup}/backup-storage | createBackupStorage
[**DeleteBackupStorageUsingPUT**](VdbBackupStorageEndPointApi.md#DeleteBackupStorageUsingPUT) | **Put** /v1/trv/{projectId}/{engineGroup}/backup-storage/delete | deleteBackupStorage
[**DeleteInTrashBackupStorageUsingDELETE**](VdbBackupStorageEndPointApi.md#DeleteInTrashBackupStorageUsingDELETE) | **Delete** /v1/trv/{projectId}/{engineGroup}/backup-storage/delete-in-trash | deleteInTrashBackupStorage
[**GetBackupStorageByEngineGroupUsingGET**](VdbBackupStorageEndPointApi.md#GetBackupStorageByEngineGroupUsingGET) | **Get** /v1/trv/{projectId}/{engineGroup}/backup-storage/ | getBackupStorageByEngineGroup
[**GetBackupStorageByIdUsingGET**](VdbBackupStorageEndPointApi.md#GetBackupStorageByIdUsingGET) | **Get** /v1/trv/{projectId}/{engineGroup}/backup-storage/{backupStorageId} | getBackupStorageById
[**GetListPackageByEngineGroupUsingGET**](VdbBackupStorageEndPointApi.md#GetListPackageByEngineGroupUsingGET) | **Get** /v1/trv/{projectId}/{engineGroup}/backup-storage/backup-package | getListPackageByEngineGroup
[**RecoverBackupStorageUsingPUT**](VdbBackupStorageEndPointApi.md#RecoverBackupStorageUsingPUT) | **Put** /v1/trv/{projectId}/{engineGroup}/backup-storage/recover | recoverBackupStorage
[**RenewBackupStorageUsingPUT**](VdbBackupStorageEndPointApi.md#RenewBackupStorageUsingPUT) | **Put** /v1/trv/{projectId}/{engineGroup}/backup-storage/renew | renewBackupStorage
[**ResizeBackupStorageUsingPUT**](VdbBackupStorageEndPointApi.md#ResizeBackupStorageUsingPUT) | **Put** /v1/trv/{projectId}/{engineGroup}/backup-storage/resize | resizeBackupStorage


# **CreateBackupStorageUsingPOST**
> BackupStorageResponse CreateBackupStorageUsingPOST(ctx, engineGroup, projectId, request)
createBackupStorage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineGroup** | **int32**| engineGroup | 
  **projectId** | **string**| projectId | 
  **request** | [**CreateBackupStorageRequest**](CreateBackupStorageRequest.md)| request | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupStorageUsingPUT**
> BackupStorageResponse DeleteBackupStorageUsingPUT(ctx, projectId, request)
deleteBackupStorage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**DeleteBackupStorageRequest**](DeleteBackupStorageRequest.md)| request | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInTrashBackupStorageUsingDELETE**
> BackupStorageResponse DeleteInTrashBackupStorageUsingDELETE(ctx, projectId, request)
deleteInTrashBackupStorage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**DeleteInTrashBackupStorageRequest**](DeleteInTrashBackupStorageRequest.md)| request | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupStorageByEngineGroupUsingGET**
> BackupStorageResponse GetBackupStorageByEngineGroupUsingGET(ctx, engineGroup, projectId)
getBackupStorageByEngineGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineGroup** | **int32**| engineGroup | 
  **projectId** | **string**| projectId | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupStorageByIdUsingGET**
> BackupStorageResponse GetBackupStorageByIdUsingGET(ctx, backupStorageId, projectId)
getBackupStorageById

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupStorageId** | **string**| backupStorageId | 
  **projectId** | **string**| projectId | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListPackageByEngineGroupUsingGET**
> ListBackupStoragePackageResponse GetListPackageByEngineGroupUsingGET(ctx, engineGroup, projectId)
getListPackageByEngineGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineGroup** | **int32**| engineGroup | 
  **projectId** | **string**| projectId | 

### Return type

[**ListBackupStoragePackageResponse**](ListBackupStoragePackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecoverBackupStorageUsingPUT**
> BackupStorageResponse RecoverBackupStorageUsingPUT(ctx, projectId, request)
recoverBackupStorage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**RecoverBackupStorageRequest**](RecoverBackupStorageRequest.md)| request | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewBackupStorageUsingPUT**
> BackupStorageResponse RenewBackupStorageUsingPUT(ctx, projectId, request)
renewBackupStorage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**RenewBackupStorageRequest**](RenewBackupStorageRequest.md)| request | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeBackupStorageUsingPUT**
> BackupStorageResponse ResizeBackupStorageUsingPUT(ctx, projectId, request)
resizeBackupStorage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| projectId | 
  **request** | [**ResizeBackupStorageRequest**](ResizeBackupStorageRequest.md)| request | 

### Return type

[**BackupStorageResponse**](BackupStorageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

