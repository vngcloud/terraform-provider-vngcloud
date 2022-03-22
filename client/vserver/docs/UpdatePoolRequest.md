# UpdatePoolRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Algorithm of the pool. The algorithm can be \&quot;ROUND_ROBIN\&quot; or \&quot;LEAST_CONNECTIONS\&quot; or \&quot;SOURCE_IP\&quot; | [default to null]
**HealthMonitor** | [***UpdateHealthMonitorRequest**](UpdateHealthMonitorRequest.md) | Update request for health monitor | [optional] [default to null]
**PoolId** | **string** | Id of the pool to update. | [default to null]
**ProjectId** | **string** | Id of project | [optional] [default to null]
**Stickiness** | **bool** | Enable sticky sessions. | [optional] [default to null]
**UserId** | **int32** | Id of user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


