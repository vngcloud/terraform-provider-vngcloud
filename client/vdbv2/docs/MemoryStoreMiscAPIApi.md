# {{classname}}

All URIs are relative to *https:/vdb-gateway.vngcloud.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDatastore1**](MemoryStoreMiscAPIApi.md#GetAllDatastore1) | **Get** /vdb-memory/v1/database/datastore | 
[**GetAllInstanceFamily1**](MemoryStoreMiscAPIApi.md#GetAllInstanceFamily1) | **Get** /vdb-memory/v1/database/families | 
[**GetDBInstanceConfig1**](MemoryStoreMiscAPIApi.md#GetDBInstanceConfig1) | **Get** /vdb-memory/v1/database/configuration | 
[**GetEngine1**](MemoryStoreMiscAPIApi.md#GetEngine1) | **Get** /vdb-memory/v1/database/engine | 
[**GetFlavorCodes1**](MemoryStoreMiscAPIApi.md#GetFlavorCodes1) | **Get** /vdb-memory/v1/database/codes | 
[**GetFlavors1**](MemoryStoreMiscAPIApi.md#GetFlavors1) | **Get** /vdb-memory/v1/database/flavors | 
[**GetListNetwork1**](MemoryStoreMiscAPIApi.md#GetListNetwork1) | **Get** /vdb-memory/v1/database/networks | 
[**GetVolumeTypes1**](MemoryStoreMiscAPIApi.md#GetVolumeTypes1) | **Get** /vdb-memory/v1/database/volume-types | 
[**ListDatabaseInstanceStatus**](MemoryStoreMiscAPIApi.md#ListDatabaseInstanceStatus) | **Get** /vdb-memory/v1/database/status | 
[**ListSubnet1**](MemoryStoreMiscAPIApi.md#ListSubnet1) | **Get** /vdb-memory/v1/database/networks/subnets | 

# **GetAllDatastore1**
> WrapContentListEngineVersion GetAllDatastore1(ctx, portalUserId)


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

# **GetAllInstanceFamily1**
> WrapContentListInstanceFamily GetAllInstanceFamily1(ctx, portalUserId)


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

# **GetDBInstanceConfig1**
> WrapContentListItemConfigInfo GetDBInstanceConfig1(ctx, portalUserId)


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

# **GetEngine1**
> WrapContentSetEngine GetEngine1(ctx, portalUserId)


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
> WrapContentListFlavorInfo GetFlavors1(ctx, type_, version, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **version** | **string**|  | 
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreMiscAPIApiGetFlavors1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreMiscAPIApiGetFlavors1Opts struct
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

# **GetListNetwork1**
> WrapContentListNetworkResponse GetListNetwork1(ctx, portalUserId)


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

# **GetVolumeTypes1**
> WrapContentVolumeTypeGatewayResponse GetVolumeTypes1(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreMiscAPIApiGetVolumeTypes1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreMiscAPIApiGetVolumeTypes1Opts struct
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

# **ListDatabaseInstanceStatus**
> WrapContentSetString ListDatabaseInstanceStatus(ctx, portalUserId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 

### Return type

[**WrapContentSetString**](WrapContentSetString.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubnet1**
> WrapContentListNetworkResponseV2 ListSubnet1(ctx, portalUserId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portalUserId** | **int32**|  | 
 **optional** | ***MemoryStoreMiscAPIApiListSubnet1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemoryStoreMiscAPIApiListSubnet1Opts struct
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

