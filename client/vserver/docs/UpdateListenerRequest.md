# UpdateListenerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCidrs** | **string** | Allowed cidr. | [default to null]
**ClientCertificate** | **string** | Client certificate | [optional] [default to null]
**DefaultCertificateAuthority** | **string** | Default certificate authority. | [optional] [default to null]
**DefaultPoolId** | **string** | Id of the pool that this listener will forward to. | [optional] [default to null]
**Headers** | **[]string** | Headers | [optional] [default to null]
**ListenerId** | **string** | Id of the listener to update | [default to null]
**TimeoutClient** | **int32** | Idle timeout of client. The value can be in range from 1 to 3600 seconds | [default to null]
**TimeoutConnection** | **int32** | Idle timeout of connection. The value can be in range from 1 to 3600 seconds | [default to null]
**TimeoutMember** | **int32** | Idle timeout of member. The value can be in range from 1 to 3600 seconds | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


