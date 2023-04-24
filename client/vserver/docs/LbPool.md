# LbPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Algorithm of the pool. The algorithm can be \&quot;ROUND_ROBIN\&quot; or \&quot;LEAST_CONNECTIONS\&quot; or \&quot;SOURCE_IP\&quot; | [default to null]
**HealthCheckMethod** | **string** | Health check method for the HTTP protocol. | [optional] [default to null]
**HealthCheckPath** | **string** | Health check path for the HTTP protocol. | [optional] [default to null]
**HealthCheckProtocol** | **string** | Protocol for performing health check. | [default to null]
**HealthyThreshold** | **int64** | Healthy threshold. The value must be in range from 2 to 10. | [optional] [default to null]
**Interval** | **int64** | Health check interval. The value must be from 5 to 3600 seconds. | [optional] [default to null]
**Members** | [**[]CreateMemberRequest**](CreateMemberRequest.md) | List of members of the pool. | [optional] [default to null]
**PoolName** | **string** | Name of the pool. Only letters (a-z, A-Z, 0-9, &#39;_&#39;, &#39;.&#39;) are allowed and your input data must be between 6 and 20 characters. | [default to null]
**PoolProtocol** | **string** | Protocol of the pool. | [default to null]
**Stickiness** | **bool** | Enable sticky sessions. | [optional] [default to null]
**SuccessCode** | **string** | Health check success code for HTTP health check protocol. | [optional] [default to null]
**Timeout** | **int64** | Timeout of health check. The value must be from 2 to 120 seconds | [optional] [default to null]
**UnhealthyThreshold** | **int64** | Unhealthy threshold. The value must be in range from 2 to 10. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


