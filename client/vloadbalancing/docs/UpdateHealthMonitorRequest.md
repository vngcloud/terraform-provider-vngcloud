# UpdateHealthMonitorRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthCheckMethod** | **string** | Health check method for the HTTP protocol. | [optional] [default to null]
**HealthCheckPath** | **string** | Health check path for the HTTP protocol. | [optional] [default to null]
**HealthyThreshold** | **int64** | Healthy threshold. The value must be in range from 2 to 10. | [default to null]
**Interval** | **int64** | Health check interval. The value must be from 5 to 3600 seconds. | [default to null]
**SuccessCode** | **string** | Health check success code for HTTP health check protocol. | [optional] [default to null]
**Timeout** | **int64** | Timeout of health check. The value must be from 2 to 120 seconds | [default to null]
**UnhealthyThreshold** | **int64** | Unhealthy threshold. The value must be in range from 2 to 10. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


