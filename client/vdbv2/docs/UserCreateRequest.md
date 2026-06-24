# UserCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User Name | [optional] [default to null]
**ProduceTopicNames** | **[]string** | List of topic names that the created user has Produce permission on. The value of this field will be ignored if &#x27;produceAll&#x27; is true. | [optional] [default to null]
**ProduceAll** | **bool** | Indicates whether the created user has Produce permission on all topics. | [optional] [default to null]
**ConsumeTopicNames** | **[]string** | List of topic names that the created user has Consume permission on. The value of this field will be ignored if &#x27;consumeAll&#x27; is true. | [optional] [default to null]
**ConsumeAll** | **bool** | Indicates whether the created user has Consume permission on all topics. | [optional] [default to null]
**ProduceConsumeTopicNames** | **[]string** | List of topic names that the created user has Produce and Consume permissions on. The value of this field will be ignored if &#x27;produceConsumeAll&#x27; is true. | [optional] [default to null]
**ProduceConsumeAll** | **bool** | Indicates whether the created user has both Produce and Consume permissions on all topics. | [optional] [default to null]
**AdminTopicNames** | **[]string** | List of topic names that the created user has Admin permission on. The value of this field will be ignored if &#x27;adminAll&#x27; is true. | [optional] [default to null]
**AdminAll** | **bool** | Indicates whether the created user has Admin permission on all topics. | [optional] [default to null]
**MtlsAuthen** | **bool** | mTLS Authentication Enabled | [optional] [default to null]
**SaslAuthen** | **bool** | SASL Authentication Enabled | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

