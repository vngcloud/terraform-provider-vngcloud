# TopicCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Topic Name | [optional] [default to null]
**Partitions** | **int32** | Partition count. Minimum 1 partition and maximum 2048 partition | [optional] [default to null]
**Replicas** | **int32** | Replication factor. Minimum 1 factor and not exceed the number of brokers | [optional] [default to null]
**RetentionSeconds** | **int64** | Retention time in seconds. Minimum 3600 seconds and maximum 7776000 seconds | [optional] [default to null]
**RetentionBytes** | **int64** | Retention byte. Minimum 1 byte and maximum 1099511627776 bytes. Kafka will create the topic with a default value if this field is empty. You can set this field to -1 for unlimited value. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

