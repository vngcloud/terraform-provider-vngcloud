# CreateListenerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCidrs** | **string** | Allowed cidr. | [default to null]
**CertificateAuthorities** | **[]string** | List of certificate authority | [optional] [default to null]
**DefaultPoolId** | **string** | Id of the pool that this listener will forward to. | [default to null]
**ListenerName** | **string** | Name of the listener. Only letters (a-z, A-Z, 0-9, &#39;_&#39;, &#39;.&#39;) are allowed and your input data must be between 6 and 20 characters. | [default to null]
**ListenerProtocol** | **string** | Protocol of the listener. | [default to null]
**ListenerProtocolPort** | **int32** | Port of the listener. | [default to null]
**LoadBalancerId** | **string** | Id of the load balancer. | [optional] [default to null]
**ProjectId** | **string** | Id of project | [optional] [default to null]
**TimeoutClient** | **int32** | Idle timeout of client. The value can be in range from 1 to 3600 seconds | [default to null]
**TimeoutConnection** | **int32** | Idle timeout of connection. The value can be in range from 1 to 3600 seconds | [default to null]
**TimeoutMember** | **int32** | Idle timeout of member. The value can be in range from 1 to 3600 seconds | [default to null]
**UserId** | **int32** | Id of user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


