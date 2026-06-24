# CreateConfigGroupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Config Group Name | [optional] [default to null]
**Description** | **string** | Config Group Description | [optional] [default to null]
**DatastoreType** | **string** | Config Group Datastore Type. Allowed values: MySQL, PostgreSQL, MariaDB | [optional] [default to null]
**DatastoreVersion** | **string** | Version of Datastore Type | [optional] [default to null]
**DeployType** | **string** | Optional. Currently only available for PostgreSQL. Allowed values: cluster, single_node. Default value: single_node.Config group with a &#x27;cluster&#x27; deploy type can only be used with a PostgreSQL Cluster. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

