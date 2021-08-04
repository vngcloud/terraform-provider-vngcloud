# CreateSecurityGroupRuleRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description for secgoup rule. | [optional] [default to null]
**Direction** | **string** | \&quot;ingress\&quot; or \&quot;egress\&quot; | [default to null]
**Ethertype** | **string** | \&quot;IPV4\&quot; or \&quot;IPv6\&quot; | [default to null]
**PortRangeMax** | **int32** | upper bound of range port. | [default to null]
**PortRangeMin** | **int32** | upper bound of range port. | [default to null]
**Protocol** | **string** | \&quot;TCP\&quot;, \&quot;UDP\&quot;, \&quot;ICMP\&quot;, ...  | [default to null]
**RemoteIpPrefix** | **string** | Ip Prefix of source/target. | [default to null]
**SecurityGroupId** | **string** | Id of security group that rules are attached. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


