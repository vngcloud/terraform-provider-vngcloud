# CreateLoadBalancerRequestV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listener** | [***LbListener**](LbListener.md) | Listener of the load balancer. | [optional] [default to null]
**Name** | **string** | Load balancer&#39;s name. Only letters (a-z, A-Z, 0-9, &#39;_&#39;, &#39;.&#39;) are allowed and your input data must be between 6 and 20 characters. | [default to null]
**PackageId** | **string** | Package ID of the load balancer. | [default to null]
**Pool** | [***LbPool**](LbPool.md) | Pool of the load balancer. | [optional] [default to null]
**Scheme** | **string** | Schema of the load balancer, it may be Internet or Internal. | [default to null]
**SubnetId** | **string** | Subnet ID for the load balancer. | [default to null]
**Type_** | **string** | Type of the load balancer. It may be Layer 4 or Layer 7 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


