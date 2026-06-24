# UpdateSecurityGroupRuleDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Secgroup rule ID. Set to null to insert new. | [optional] [default to null]
**PortRangeMin** | **int32** | Port range min. Allowed values: 3306 (MySQL, MariaDB), 5432 (PostgreSQL Single-node), 5432 or 15432 (PostgreSQL Cluster), 6379 (Redis) | [optional] [default to null]
**PortRangeMax** | **int32** | Port range max. Allowed values: 3306 (MySQL, MariaDB), 5432 (PostgreSQL Single-node), 5432 or 15432 (PostgreSQL Cluster), 6379 (Redis) | [optional] [default to null]
**RemoteIpPrefix** | **string** | IP Prefix | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

