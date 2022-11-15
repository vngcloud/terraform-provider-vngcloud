# \ServerRestControllerV2Api

All URIs are relative to *https://virtserver.swaggerhub.com/manhtu1997/vserver/1.0.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachExternalNetworkInterfaceUsingPOST**](ServerRestControllerV2Api.md#AttachExternalNetworkInterfaceUsingPOST) | **Post** /v2/{projectId}/servers/{serverId}/external-network-interfaces | Attach External Network Interface
[**AttachNetworkInterfaceUsingPOST**](ServerRestControllerV2Api.md#AttachNetworkInterfaceUsingPOST) | **Post** /v2/{projectId}/servers/{serverId}/internal-network-interfaces | Attach Internal Network Interface
[**AttachWanIPUsingPUT**](ServerRestControllerV2Api.md#AttachWanIPUsingPUT) | **Put** /v2/{projectId}/servers/{serverId}/wan-ips/{wanIpId}/attach | Attach Wan IP
[**CreateServerUsingPOST1**](ServerRestControllerV2Api.md#CreateServerUsingPOST1) | **Post** /v2/{projectId}/servers | Create server
[**DeleteServerUsingDELETE1**](ServerRestControllerV2Api.md#DeleteServerUsingDELETE1) | **Delete** /v2/{projectId}/servers/{serverId} | Delete Server
[**DetachExternalNetworkInterfaceUsingDELETE**](ServerRestControllerV2Api.md#DetachExternalNetworkInterfaceUsingDELETE) | **Delete** /v2/{projectId}/servers/{serverId}/external-network-interfaces | Detach External Network Interface
[**DetachNetworkInterfaceUsingDELETE**](ServerRestControllerV2Api.md#DetachNetworkInterfaceUsingDELETE) | **Delete** /v2/{projectId}/servers/{serverId}/internal-network-interfaces | Detach Internal Network Interface
[**DetachWanIPUsingPUT**](ServerRestControllerV2Api.md#DetachWanIPUsingPUT) | **Put** /v2/{projectId}/servers/{serverId}/wan-ips/{wanIpId}/detach | Detach Wan IP
[**GetConsoleLogUsingGET**](ServerRestControllerV2Api.md#GetConsoleLogUsingGET) | **Get** /v2/{projectId}/servers/{serverId}/console-log | Get console log
[**GetConsoleUrlUsingGET**](ServerRestControllerV2Api.md#GetConsoleUrlUsingGET) | **Get** /v2/{projectId}/servers/{serverId}/console-url | Get console url
[**GetServerUsingGET1**](ServerRestControllerV2Api.md#GetServerUsingGET1) | **Get** /v2/{projectId}/servers/{serverId} | Get server by id
[**ListActionUsingGET**](ServerRestControllerV2Api.md#ListActionUsingGET) | **Get** /v2/{projectId}/servers/{serverId}/actions | List Action Of Server
[**ListNetworkInterfaceUsingGET**](ServerRestControllerV2Api.md#ListNetworkInterfaceUsingGET) | **Get** /v2/{projectId}/servers/{serverId}/network-interfaces | List Network Interface Of Server
[**ListSecGroupUsingGET**](ServerRestControllerV2Api.md#ListSecGroupUsingGET) | **Get** /v2/{projectId}/servers/{serverId}/sec-groups | List Sec Group Of Server
[**ListServerUsingGET1**](ServerRestControllerV2Api.md#ListServerUsingGET1) | **Get** /v2/{projectId}/servers | List server
[**RebootServerUsingPUT1**](ServerRestControllerV2Api.md#RebootServerUsingPUT1) | **Put** /v2/{projectId}/servers/{serverId}/reboot | Reboot server
[**RenameServerUsingPUT**](ServerRestControllerV2Api.md#RenameServerUsingPUT) | **Put** /v2/{projectId}/servers/{serverId}/rename | Rename Server
[**ResizeServerUsingPUT1**](ServerRestControllerV2Api.md#ResizeServerUsingPUT1) | **Put** /v2/{projectId}/servers/{serverId}/resize | Change flavor of server
[**StartServerUsingPUT1**](ServerRestControllerV2Api.md#StartServerUsingPUT1) | **Put** /v2/{projectId}/servers/{serverId}/start | Start server
[**StopServerUsingPUT1**](ServerRestControllerV2Api.md#StopServerUsingPUT1) | **Put** /v2/{projectId}/servers/{serverId}/stop | Stop server
[**UpdateSecGroupServerUsingPUT1**](ServerRestControllerV2Api.md#UpdateSecGroupServerUsingPUT1) | **Put** /v2/{projectId}/servers/{serverId}/update-sec-group | Update SecGroups of server


# **AttachExternalNetworkInterfaceUsingPOST**
> DataResponse AttachExternalNetworkInterfaceUsingPOST(ctx, attachExternalNetworkInterfaceRequest, projectId, serverId)
Attach External Network Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachExternalNetworkInterfaceRequest** | [**AttachExternalNetworkInterfaceRequest**](AttachExternalNetworkInterfaceRequest.md)| attachExternalNetworkInterfaceRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachNetworkInterfaceUsingPOST**
> DataResponse AttachNetworkInterfaceUsingPOST(ctx, attachNetworkInterfaceRequest, projectId, serverId)
Attach Internal Network Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachNetworkInterfaceRequest** | [**AttachNetworkInterfaceRequest**](AttachNetworkInterfaceRequest.md)| attachNetworkInterfaceRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachWanIPUsingPUT**
> AttachWanIPUsingPUT(ctx, attachDetachWanIPRequest, projectId, serverId, wanIpId)
Attach Wan IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachDetachWanIPRequest** | [**AttachDetachWanIpRequest**](AttachDetachWanIpRequest.md)| attachDetachWanIPRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 
  **wanIpId** | **string**| The wan ip id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServerUsingPOST1**
> DataResponse CreateServerUsingPOST1(ctx, createServerRequest, projectId)
Create server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createServerRequest** | [**CreateServerRequest**](CreateServerRequest.md)| createServerRequest | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerUsingDELETE1**
> DataResponse DeleteServerUsingDELETE1(ctx, deleteServerRequest, projectId, serverId)
Delete Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteServerRequest** | [**DeleteServerRequest**](DeleteServerRequest.md)| deleteServerRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachExternalNetworkInterfaceUsingDELETE**
> DetachExternalNetworkInterfaceUsingDELETE(ctx, detachExternalNetworkInterfaceRequest, projectId, serverId)
Detach External Network Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **detachExternalNetworkInterfaceRequest** | [**DetachExternalNetworkInterfaceRequest**](DetachExternalNetworkInterfaceRequest.md)| detachExternalNetworkInterfaceRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachNetworkInterfaceUsingDELETE**
> DetachNetworkInterfaceUsingDELETE(ctx, detachNetworkInterfaceRequest, projectId, serverId)
Detach Internal Network Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **detachNetworkInterfaceRequest** | [**DetachNetworkInterfaceRequest**](DetachNetworkInterfaceRequest.md)| detachNetworkInterfaceRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachWanIPUsingPUT**
> DetachWanIPUsingPUT(ctx, attachDetachWanIPRequest, projectId, serverId, wanIpId)
Detach Wan IP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachDetachWanIPRequest** | [**AttachDetachWanIpRequest**](AttachDetachWanIpRequest.md)| attachDetachWanIPRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 
  **wanIpId** | **string**| The wan ip id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsoleLogUsingGET**
> DataResponsestring GetConsoleLogUsingGET(ctx, projectId, serverId)
Get console log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponsestring**](DataResponse«string».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsoleUrlUsingGET**
> DataResponsestring GetConsoleUrlUsingGET(ctx, projectId, serverId)
Get console url

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponsestring**](DataResponse«string».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerUsingGET1**
> DataResponseServer GetServerUsingGET1(ctx, projectId, serverId)
Get server by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponseServer**](DataResponse«Server».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActionUsingGET**
> DataResponseListServerAction ListActionUsingGET(ctx, projectId, serverId)
List Action Of Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponseListServerAction**](DataResponse«List«ServerAction»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkInterfaceUsingGET**
> DataResponseServerNetworkInterfaceDetail ListNetworkInterfaceUsingGET(ctx, projectId, serverId)
List Network Interface Of Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponseServerNetworkInterfaceDetail**](DataResponse«ServerNetworkInterfaceDetail».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSecGroupUsingGET**
> DataResponseServerSecGroupDetail ListSecGroupUsingGET(ctx, projectId, serverId)
List Sec Group Of Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponseServerSecGroupDetail**](DataResponse«ServerSecGroupDetail».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServerUsingGET1**
> PagingServer ListServerUsingGET1(ctx, projectId, optional)
List server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
 **optional** | ***ServerRestControllerV2ApiListServerUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerRestControllerV2ApiListServerUsingGET1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name | 
 **page** | **optional.Int32**| page | 
 **size** | **optional.Int32**| size | 

### Return type

[**PagingServer**](Paging«Server».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootServerUsingPUT1**
> DataResponse RebootServerUsingPUT1(ctx, projectId, serverId)
Reboot server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameServerUsingPUT**
> DataResponseServer RenameServerUsingPUT(ctx, projectId, renameServerRequest, serverId)
Rename Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **renameServerRequest** | [**RenameServerRequest**](RenameServerRequest.md)| renameServerRequest | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponseServer**](DataResponse«Server».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResizeServerUsingPUT1**
> DataResponse ResizeServerUsingPUT1(ctx, projectId, resizeServerRequest, serverId)
Change flavor of server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **resizeServerRequest** | [**ResizeServerRequest**](ResizeServerRequest.md)| resizeServerRequest | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartServerUsingPUT1**
> DataResponse StartServerUsingPUT1(ctx, projectId, serverId)
Start server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopServerUsingPUT1**
> DataResponse StopServerUsingPUT1(ctx, projectId, serverId)
Stop server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The server id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecGroupServerUsingPUT1**
> DataResponse UpdateSecGroupServerUsingPUT1(ctx, changeSecGroupRequest, projectId, serverId)
Update SecGroups of server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeSecGroupRequest** | [**ChangeSecGroupRequest**](ChangeSecGroupRequest.md)| changeSecGroupRequest | 
  **projectId** | **string**| The project id | 
  **serverId** | **string**| The project id | 

### Return type

[**DataResponse**](DataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

