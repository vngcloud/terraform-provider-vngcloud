# CreatePoolRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Algorithm of the pool. The algorithm can be \&quot;ROUND_ROBIN\&quot; or \&quot;LEAST_CONNECTIONS\&quot; or \&quot;SOURCE_IP\&quot; | [default to null]
**HealthMonitor** | [***CreateHealthMonitorRequest**](CreateHealthMonitorRequest.md) | Request for creating health check monitor. | [optional] [default to null]
**LoadBalancerId** | **string** | Id of the load balancer. | [optional] [default to null]
**Members** | [**[]CreateMemberRequest**](CreateMemberRequest.md) | List of members of the pool. | [optional] [default to null]
**PoolName** | **string** | Name of the pool. Only letters (a-z, A-Z, 0-9, &#39;_&#39;, &#39;.&#39;) are allowed and your input data must be between 6 and 20 characters. | [default to null]
**PoolProtocol** | **string** | Protocol of the pool. | [default to null]
**Stickiness** | **bool** | Enable sticky sessions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


