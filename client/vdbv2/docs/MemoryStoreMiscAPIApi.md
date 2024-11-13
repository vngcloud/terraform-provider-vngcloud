# {{classname}}

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDatastore**](MemoryStoreMiscAPIApi.md#GetAllDatastore) | **Get** /vdb-memory/v1/database/datastore | 
[**GetAllInstanceFamily**](MemoryStoreMiscAPIApi.md#GetAllInstanceFamily) | **Get** /vdb-memory/v1/database/families | 
[**GetDBInstanceConfig**](MemoryStoreMiscAPIApi.md#GetDBInstanceConfig) | **Get** /vdb-memory/v1/database/configuration | 
[**GetEngine**](MemoryStoreMiscAPIApi.md#GetEngine) | **Get** /vdb-memory/v1/database/engine | 
[**GetFlavorCodes**](MemoryStoreMiscAPIApi.md#GetFlavorCodes) | **Get** /vdb-memory/v1/database/codes | 
[**GetFlavors**](MemoryStoreMiscAPIApi.md#GetFlavors) | **Get** /vdb-memory/v1/database/flavors | 
[**GetListNetwork**](MemoryStoreMiscAPIApi.md#GetListNetwork) | **Get** /vdb-memory/v1/database/networks | 
[**GetVolumeTypes**](MemoryStoreMiscAPIApi.md#GetVolumeTypes) | **Get** /vdb-memory/v1/database/volume-types | 
[**ListDatabaseInstanceStatus**](MemoryStoreMiscAPIApi.md#ListDatabaseInstanceStatus) | **Get** /vdb-memory/v1/database/status | 
[**ListNetwork**](MemoryStoreMiscAPIApi.md#ListNetwork) | **Get** /vdb-memory/v1/database/networks/subnets | 

# **GetAllDatastore**
> WrapContentListEngineVersion GetAllDatastore(ctx, )


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

# **GetAllInstanceFamily**
> WrapContentListInstanceFamily GetAllInstanceFamily(ctx, )


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

# **GetDBInstanceConfig**
> WrapContentListItemConfigInfo GetDBInstanceConfig(ctx, )


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

# **GetEngine**
> WrapContentSetEngine GetEngine(ctx, )


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
> WrapContentListFlavorInfo GetFlavors(ctx, type_, version, portalUserId)


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

# **GetListNetwork**
> WrapContentListNetworkResponse GetListNetwork(ctx, )


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

# **GetVolumeTypes**
> WrapContentVolumeTypeGatewayResponse GetVolumeTypes(ctx, portalUserId)


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

# **ListDatabaseInstanceStatus**
> WrapContentSetString ListDatabaseInstanceStatus(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WrapContentSetString**](WrapContentSetString.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetwork**
> WrapContentListNetworkResponseV2 ListNetwork(ctx, )


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

