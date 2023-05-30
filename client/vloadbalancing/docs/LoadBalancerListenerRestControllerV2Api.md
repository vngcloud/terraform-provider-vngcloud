# \LoadBalancerListenerRestControllerV2Api

All URIs are relative to *https://localhost:8089*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListenerUsingPOST**](LoadBalancerListenerRestControllerV2Api.md#CreateListenerUsingPOST) | **Post** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners | Create a listener
[**CreatePolicyUsingPOST**](LoadBalancerListenerRestControllerV2Api.md#CreatePolicyUsingPOST) | **Post** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies | Create a l7 policy
[**DeleteListenerUsingDELETE**](LoadBalancerListenerRestControllerV2Api.md#DeleteListenerUsingDELETE) | **Delete** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId} | Delete a listener
[**DeletePolicyUsingDELETE**](LoadBalancerListenerRestControllerV2Api.md#DeletePolicyUsingDELETE) | **Delete** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies/{l7PolicyId} | Delete a policy
[**GetListL7PoliciesUsingGET**](LoadBalancerListenerRestControllerV2Api.md#GetListL7PoliciesUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies | Get list policies of a listener
[**GetListListenersByLoadBalancerUsingGET**](LoadBalancerListenerRestControllerV2Api.md#GetListListenersByLoadBalancerUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners | Get list of listeners of a Load Balancer
[**GetListenersUsingGET**](LoadBalancerListenerRestControllerV2Api.md#GetListenersUsingGET) | **Get** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId} | Get a specific listener
[**ReorderPoliciesUsingPUT**](LoadBalancerListenerRestControllerV2Api.md#ReorderPoliciesUsingPUT) | **Put** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/reorderL7Policies | Reorder the L7 policies
[**UpdateL7PolicyUsingPUT**](LoadBalancerListenerRestControllerV2Api.md#UpdateL7PolicyUsingPUT) | **Put** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies/{policyId} | Update a policy
[**UpdateListenerUsingPUT**](LoadBalancerListenerRestControllerV2Api.md#UpdateListenerUsingPUT) | **Put** /v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId} | Update a listener


# **CreateListenerUsingPOST**
> DataResponseListener CreateListenerUsingPOST(ctx, authorization, createListenerRequestV2, loadBalancerId, portalUserId, projectId)
Create a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **createListenerRequestV2** | [**CreateListenerRequestV2**](CreateListenerRequestV2.md)| createListenerRequestV2 | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListener**](DataResponse«Listener».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePolicyUsingPOST**
> DataResponseL7Policy CreatePolicyUsingPOST(ctx, authorization, createL7PolicyRequestV2, listenerId, portalUserId, projectId)
Create a l7 policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **createL7PolicyRequestV2** | [**CreateL7PolicyRequestV2**](CreateL7PolicyRequestV2.md)| createL7PolicyRequestV2 | 
  **listenerId** | **string**| The listener id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseL7Policy**](DataResponse«L7Policy».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListenerUsingDELETE**
> DeleteListenerUsingDELETE(ctx, authorization, listenerId, portalUserId, projectId)
Delete a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| listener id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicyUsingDELETE**
> DeletePolicyUsingDELETE(ctx, authorization, listenerId, policyId, portalUserId, projectId)
Delete a policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| listener id | 
  **policyId** | **string**| policy id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| project id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListL7PoliciesUsingGET**
> DataResponseListL7Policy GetListL7PoliciesUsingGET(ctx, authorization, listenerId, loadBalancerId, portalUserId, projectId)
Get list policies of a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| The listener id | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListL7Policy**](DataResponse«List«L7Policy»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListListenersByLoadBalancerUsingGET**
> DataResponseListListener GetListListenersByLoadBalancerUsingGET(ctx, authorization, loadBalancerId, portalUserId, projectId)
Get list of listeners of a Load Balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListListener**](DataResponse«List«Listener»».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListenersUsingGET**
> DataResponseListener GetListenersUsingGET(ctx, authorization, listenerId, loadBalancerId, portalUserId, projectId)
Get a specific listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| The listener id | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 

### Return type

[**DataResponseListener**](DataResponse«Listener».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReorderPoliciesUsingPUT**
> map[string]string ReorderPoliciesUsingPUT(ctx, authorization, listenerId, portalUserId, projectId, reorderPoliciesRequestV2)
Reorder the L7 policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| listener id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| project id | 
  **reorderPoliciesRequestV2** | [**ReorderPoliciesRequestV2**](ReorderPoliciesRequestV2.md)| reorderPoliciesRequestV2 | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateL7PolicyUsingPUT**
> DataResponseL7Policy UpdateL7PolicyUsingPUT(ctx, authorization, listenerId, policyId, portalUserId, projectId, updateL7PolicyRequestV2)
Update a policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| The listener id | 
  **policyId** | **string**| policy id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| project id | 
  **updateL7PolicyRequestV2** | [**UpdateL7PolicyRequestV2**](UpdateL7PolicyRequestV2.md)| updateL7PolicyRequestV2 | 

### Return type

[**DataResponseL7Policy**](DataResponse«L7Policy».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateListenerUsingPUT**
> DataResponseListener UpdateListenerUsingPUT(ctx, authorization, listenerId, loadBalancerId, portalUserId, projectId, updateListenerRequestV2)
Update a listener

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Access Token | 
  **listenerId** | **string**| The listener id | 
  **loadBalancerId** | **string**| The Load Balancer id | 
  **portalUserId** | **int32**| portal-user-id | 
  **projectId** | **string**| The project id | 
  **updateListenerRequestV2** | [**UpdateListenerRequestV2**](UpdateListenerRequestV2.md)| updateListenerRequestV2 | 

### Return type

[**DataResponseListener**](DataResponse«Listener».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

