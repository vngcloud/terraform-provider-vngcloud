# Go API client for swagger

API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://vdb-gateway.vngcloud.vn*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*KafkaClusterAPIApi* | [**CreateOrderCluster**](docs/KafkaClusterAPIApi.md#createordercluster) | **Post** /vdb-kafka/clusters | 
*KafkaClusterAPIApi* | [**CreateSecRule**](docs/KafkaClusterAPIApi.md#createsecrule) | **Post** /vdb-kafka/clusters/{clusterId}/security-group-rules | 
*KafkaClusterAPIApi* | [**CreateTopic**](docs/KafkaClusterAPIApi.md#createtopic) | **Post** /vdb-kafka/clusters/{clusterId}/topics | 
*KafkaClusterAPIApi* | [**CreateUser**](docs/KafkaClusterAPIApi.md#createuser) | **Post** /vdb-kafka/clusters/{clusterId}/users | 
*KafkaClusterAPIApi* | [**DeleteCluster**](docs/KafkaClusterAPIApi.md#deletecluster) | **Delete** /vdb-kafka/clusters/{clusterId} | 
*KafkaClusterAPIApi* | [**DeleteSecRule**](docs/KafkaClusterAPIApi.md#deletesecrule) | **Delete** /vdb-kafka/clusters/{clusterId}/security-group-rules/{secGroupRuleId} | 
*KafkaClusterAPIApi* | [**DeleteTopic**](docs/KafkaClusterAPIApi.md#deletetopic) | **Delete** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
*KafkaClusterAPIApi* | [**DeleteUser**](docs/KafkaClusterAPIApi.md#deleteuser) | **Delete** /vdb-kafka/clusters/{clusterId}/users/{userId} | 
*KafkaClusterAPIApi* | [**GetClusterById**](docs/KafkaClusterAPIApi.md#getclusterbyid) | **Get** /vdb-kafka/clusters/{clusterId} | 
*KafkaClusterAPIApi* | [**GetTopicById**](docs/KafkaClusterAPIApi.md#gettopicbyid) | **Get** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
*KafkaClusterAPIApi* | [**GetUserAuthenCreds**](docs/KafkaClusterAPIApi.md#getuserauthencreds) | **Get** /vdb-kafka/clusters/{clusterId}/users/{userId}/authen-creds | 
*KafkaClusterAPIApi* | [**GetUserById**](docs/KafkaClusterAPIApi.md#getuserbyid) | **Get** /vdb-kafka/clusters/{clusterId}/users/{userId} | 
*KafkaClusterAPIApi* | [**ListClusters**](docs/KafkaClusterAPIApi.md#listclusters) | **Get** /vdb-kafka/clusters | 
*KafkaClusterAPIApi* | [**ListHistory**](docs/KafkaClusterAPIApi.md#listhistory) | **Get** /vdb-kafka/clusters/{clusterId}/history | 
*KafkaClusterAPIApi* | [**ListTopic**](docs/KafkaClusterAPIApi.md#listtopic) | **Get** /vdb-kafka/clusters/{clusterId}/topics | 
*KafkaClusterAPIApi* | [**ListUser**](docs/KafkaClusterAPIApi.md#listuser) | **Get** /vdb-kafka/clusters/{clusterId}/users | 
*KafkaClusterAPIApi* | [**RegenerateUserAuthenCreds**](docs/KafkaClusterAPIApi.md#regenerateuserauthencreds) | **Put** /vdb-kafka/clusters/{clusterId}/users/{userId}/regenerate-creds | 
*KafkaClusterAPIApi* | [**UpdateAuthentication**](docs/KafkaClusterAPIApi.md#updateauthentication) | **Put** /vdb-kafka/clusters/{clusterId}/authentication | 
*KafkaClusterAPIApi* | [**UpdateBrokerCount**](docs/KafkaClusterAPIApi.md#updatebrokercount) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-broker-count | 
*KafkaClusterAPIApi* | [**UpdateConfigGroup**](docs/KafkaClusterAPIApi.md#updateconfiggroup) | **Put** /vdb-kafka/clusters/{clusterId}/config-group | 
*KafkaClusterAPIApi* | [**UpdatePublicAccess**](docs/KafkaClusterAPIApi.md#updatepublicaccess) | **Put** /vdb-kafka/clusters/{clusterId}/public-access | 
*KafkaClusterAPIApi* | [**UpdateStorageSize**](docs/KafkaClusterAPIApi.md#updatestoragesize) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-storage-size | 
*KafkaClusterAPIApi* | [**UpdateStorageType**](docs/KafkaClusterAPIApi.md#updatestoragetype) | **Put** /vdb-kafka/clusters/{clusterId}/kafka-storage-type | 
*KafkaClusterAPIApi* | [**UpdateTopic**](docs/KafkaClusterAPIApi.md#updatetopic) | **Put** /vdb-kafka/clusters/{clusterId}/topics/{topicId} | 
*KafkaClusterAPIApi* | [**UpdateUser**](docs/KafkaClusterAPIApi.md#updateuser) | **Put** /vdb-kafka/clusters/{clusterId}/users/{userId} | 
*KafkaConfigurationGroupAPIApi* | [**CreateConfigGroup**](docs/KafkaConfigurationGroupAPIApi.md#createconfiggroup) | **Post** /vdb-kafka/config-groups | 
*KafkaConfigurationGroupAPIApi* | [**CreateConfigGroupVersion**](docs/KafkaConfigurationGroupAPIApi.md#createconfiggroupversion) | **Post** /vdb-kafka/config-groups/{configGroupId}/versions | 
*KafkaConfigurationGroupAPIApi* | [**DeleteConfigGroup**](docs/KafkaConfigurationGroupAPIApi.md#deleteconfiggroup) | **Delete** /vdb-kafka/config-groups/{configGroupId} | 
*KafkaConfigurationGroupAPIApi* | [**GetConfigGroupById**](docs/KafkaConfigurationGroupAPIApi.md#getconfiggroupbyid) | **Get** /vdb-kafka/config-groups/{configGroupId} | 
*KafkaConfigurationGroupAPIApi* | [**GetConfigGroupVersionById**](docs/KafkaConfigurationGroupAPIApi.md#getconfiggroupversionbyid) | **Get** /vdb-kafka/config-groups/{configGroupId}/versions/{configGroupVersionId} | 
*KafkaConfigurationGroupAPIApi* | [**ListConfigGroup**](docs/KafkaConfigurationGroupAPIApi.md#listconfiggroup) | **Get** /vdb-kafka/config-groups | 
*KafkaMiscAPIApi* | [**GetVolumeTypes1**](docs/KafkaMiscAPIApi.md#getvolumetypes1) | **Get** /vdb-kafka/database/volume-types | 
*KafkaMiscAPIApi* | [**ListAppConfigs**](docs/KafkaMiscAPIApi.md#listappconfigs) | **Get** /vdb-kafka/database/configs | 
*KafkaMiscAPIApi* | [**ListCode**](docs/KafkaMiscAPIApi.md#listcode) | **Get** /vdb-kafka/database/codes | 
*KafkaMiscAPIApi* | [**ListFamily**](docs/KafkaMiscAPIApi.md#listfamily) | **Get** /vdb-kafka/database/families | 
*KafkaMiscAPIApi* | [**ListFlavor**](docs/KafkaMiscAPIApi.md#listflavor) | **Get** /vdb-kafka/database/flavors | 
*MemoryStoreBackupAPIApi* | [**CreateBackups**](docs/MemoryStoreBackupAPIApi.md#createbackups) | **Post** /vdb-memory/v1/backups/create | 
*MemoryStoreBackupAPIApi* | [**DeleteBackups**](docs/MemoryStoreBackupAPIApi.md#deletebackups) | **Post** /vdb-memory/v1/backups/delete | 
*MemoryStoreBackupAPIApi* | [**GetDetailBackupById**](docs/MemoryStoreBackupAPIApi.md#getdetailbackupbyid) | **Get** /vdb-memory/v1/backups/{backupId}/detail | 
*MemoryStoreBackupAPIApi* | [**GetFreeBackupUsage**](docs/MemoryStoreBackupAPIApi.md#getfreebackupusage) | **Get** /vdb-memory/v1/backups/free-backup | 
*MemoryStoreBackupAPIApi* | [**GetListBackups**](docs/MemoryStoreBackupAPIApi.md#getlistbackups) | **Get** /vdb-memory/v1/backups | 
*MemoryStoreBackupAPIApi* | [**RestoreBackup**](docs/MemoryStoreBackupAPIApi.md#restorebackup) | **Post** /vdb-memory/v1/backups/{backupId}/restore | 
*MemoryStoreBackupStorageAPIApi* | [**CreateMemoryStoreBackUpStorage**](docs/MemoryStoreBackupStorageAPIApi.md#creatememorystorebackupstorage) | **Post** /vdb-memory/v1/payment/backup-storages | 
*MemoryStoreBackupStorageAPIApi* | [**DeleteBackupStorage**](docs/MemoryStoreBackupStorageAPIApi.md#deletebackupstorage) | **Post** /vdb-memory/v1/backup-storages/actions/delete | 
*MemoryStoreBackupStorageAPIApi* | [**GetListBackupStorage**](docs/MemoryStoreBackupStorageAPIApi.md#getlistbackupstorage) | **Get** /vdb-memory/v1/backup-storages | 
*MemoryStoreBackupStorageAPIApi* | [**GetListQuotaPackage**](docs/MemoryStoreBackupStorageAPIApi.md#getlistquotapackage) | **Get** /vdb-memory/v1/backup-storages/packages | 
*MemoryStoreBackupStorageAPIApi* | [**RenewBackupStorage**](docs/MemoryStoreBackupStorageAPIApi.md#renewbackupstorage) | **Post** /vdb-memory/v1/backup-storages/actions/renew | 
*MemoryStoreBackupStorageAPIApi* | [**ResizeBackupStorage**](docs/MemoryStoreBackupStorageAPIApi.md#resizebackupstorage) | **Post** /vdb-memory/v1/backup-storages/actions/resize | 
*MemoryStoreConfigurationGroupAPIApi* | [**CreateConfig**](docs/MemoryStoreConfigurationGroupAPIApi.md#createconfig) | **Post** /vdb-memory/v1/configurations/create | 
*MemoryStoreConfigurationGroupAPIApi* | [**DeleteConfigs**](docs/MemoryStoreConfigurationGroupAPIApi.md#deleteconfigs) | **Post** /vdb-memory/v1/configurations/delete | 
*MemoryStoreConfigurationGroupAPIApi* | [**GetConfigParams**](docs/MemoryStoreConfigurationGroupAPIApi.md#getconfigparams) | **Get** /vdb-memory/v1/configurations/params | 
*MemoryStoreConfigurationGroupAPIApi* | [**GetConfigsById**](docs/MemoryStoreConfigurationGroupAPIApi.md#getconfigsbyid) | **Get** /vdb-memory/v1/configurations/{configId}/detail | 
*MemoryStoreConfigurationGroupAPIApi* | [**GetListConfigs**](docs/MemoryStoreConfigurationGroupAPIApi.md#getlistconfigs) | **Get** /vdb-memory/v1/configurations | 
*MemoryStoreConfigurationGroupAPIApi* | [**UpdateConfig**](docs/MemoryStoreConfigurationGroupAPIApi.md#updateconfig) | **Put** /vdb-memory/v1/configurations/update | 
*MemoryStoreDatabaseAPIApi* | [**CreateDatabaseInstanceReplica**](docs/MemoryStoreDatabaseAPIApi.md#createdatabaseinstancereplica) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/create-replicas | 
*MemoryStoreDatabaseAPIApi* | [**CreateMemoryStoreDatabaseInstance**](docs/MemoryStoreDatabaseAPIApi.md#creatememorystoredatabaseinstance) | **Post** /vdb-memory/v1/payment/database-instances | 
*MemoryStoreDatabaseAPIApi* | [**DeleteDatabaseInstances**](docs/MemoryStoreDatabaseAPIApi.md#deletedatabaseinstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/delete | 
*MemoryStoreDatabaseAPIApi* | [**DetachReplica**](docs/MemoryStoreDatabaseAPIApi.md#detachreplica) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/detach-replica | 
*MemoryStoreDatabaseAPIApi* | [**ExchangePocResource**](docs/MemoryStoreDatabaseAPIApi.md#exchangepocresource) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/exchange-poc | 
*MemoryStoreDatabaseAPIApi* | [**GetDatabaseInstancesById**](docs/MemoryStoreDatabaseAPIApi.md#getdatabaseinstancesbyid) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId} | 
*MemoryStoreDatabaseAPIApi* | [**GetDatabaseInstancesByUser**](docs/MemoryStoreDatabaseAPIApi.md#getdatabaseinstancesbyuser) | **Get** /vdb-memory/v1/database-instances | 
*MemoryStoreDatabaseAPIApi* | [**GetHistoryDB**](docs/MemoryStoreDatabaseAPIApi.md#gethistorydb) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/histories | 
*MemoryStoreDatabaseAPIApi* | [**GetListBackupsByInstanceId**](docs/MemoryStoreDatabaseAPIApi.md#getlistbackupsbyinstanceid) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/backups | 
*MemoryStoreDatabaseAPIApi* | [**GetListReplicas**](docs/MemoryStoreDatabaseAPIApi.md#getlistreplicas) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/replicas | 
*MemoryStoreDatabaseAPIApi* | [**GetSecurityRules1**](docs/MemoryStoreDatabaseAPIApi.md#getsecurityrules1) | **Get** /vdb-memory/v1/database-instances/{dbInstanceId}/secrules | 
*MemoryStoreDatabaseAPIApi* | [**RenewResource**](docs/MemoryStoreDatabaseAPIApi.md#renewresource) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/renew-resource | 
*MemoryStoreDatabaseAPIApi* | [**ResizeInstance**](docs/MemoryStoreDatabaseAPIApi.md#resizeinstance) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/resize-instance | 
*MemoryStoreDatabaseAPIApi* | [**RestartDatabaseInstances**](docs/MemoryStoreDatabaseAPIApi.md#restartdatabaseinstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/reboot | 
*MemoryStoreDatabaseAPIApi* | [**StartDatabaseInstances**](docs/MemoryStoreDatabaseAPIApi.md#startdatabaseinstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/start | 
*MemoryStoreDatabaseAPIApi* | [**StopDatabaseInstances**](docs/MemoryStoreDatabaseAPIApi.md#stopdatabaseinstances) | **Post** /vdb-memory/v1/database-instances/{dbInstanceId}/shutdown | 
*MemoryStoreDatabaseAPIApi* | [**UpdateDatabaseConfigGroup**](docs/MemoryStoreDatabaseAPIApi.md#updatedatabaseconfiggroup) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/update-config-group | 
*MemoryStoreDatabaseAPIApi* | [**UpdateDatabaseSetting**](docs/MemoryStoreDatabaseAPIApi.md#updatedatabasesetting) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/update-setting | 
*MemoryStoreDatabaseAPIApi* | [**UpdateSecurityRules**](docs/MemoryStoreDatabaseAPIApi.md#updatesecurityrules) | **Put** /vdb-memory/v1/database-instances/{dbInstanceId}/secrules | 
*MemoryStoreMiscAPIApi* | [**GetAllDatastore**](docs/MemoryStoreMiscAPIApi.md#getalldatastore) | **Get** /vdb-memory/v1/database/datastore | 
*MemoryStoreMiscAPIApi* | [**GetAllInstanceFamily**](docs/MemoryStoreMiscAPIApi.md#getallinstancefamily) | **Get** /vdb-memory/v1/database/families | 
*MemoryStoreMiscAPIApi* | [**GetDBInstanceConfig**](docs/MemoryStoreMiscAPIApi.md#getdbinstanceconfig) | **Get** /vdb-memory/v1/database/configuration | 
*MemoryStoreMiscAPIApi* | [**GetEngine**](docs/MemoryStoreMiscAPIApi.md#getengine) | **Get** /vdb-memory/v1/database/engine | 
*MemoryStoreMiscAPIApi* | [**GetFlavorCodes**](docs/MemoryStoreMiscAPIApi.md#getflavorcodes) | **Get** /vdb-memory/v1/database/codes | 
*MemoryStoreMiscAPIApi* | [**GetFlavors**](docs/MemoryStoreMiscAPIApi.md#getflavors) | **Get** /vdb-memory/v1/database/flavors | 
*MemoryStoreMiscAPIApi* | [**GetListNetwork**](docs/MemoryStoreMiscAPIApi.md#getlistnetwork) | **Get** /vdb-memory/v1/database/networks | 
*MemoryStoreMiscAPIApi* | [**GetVolumeTypes**](docs/MemoryStoreMiscAPIApi.md#getvolumetypes) | **Get** /vdb-memory/v1/database/volume-types | 
*MemoryStoreMiscAPIApi* | [**ListDatabaseInstanceStatus**](docs/MemoryStoreMiscAPIApi.md#listdatabaseinstancestatus) | **Get** /vdb-memory/v1/database/status | 
*MemoryStoreMiscAPIApi* | [**ListNetwork**](docs/MemoryStoreMiscAPIApi.md#listnetwork) | **Get** /vdb-memory/v1/database/networks/subnets | 
*PricingControllerApi* | [**CalculatePrice**](docs/PricingControllerApi.md#calculateprice) | **Post** /v1/price | 
*RelationalBackupAPIApi* | [**CreateBackups1**](docs/RelationalBackupAPIApi.md#createbackups1) | **Post** /v1/backups/create | 
*RelationalBackupAPIApi* | [**DeleteBackups1**](docs/RelationalBackupAPIApi.md#deletebackups1) | **Delete** /v1/backups/{backupId}/delete | 
*RelationalBackupAPIApi* | [**GetDetailBackupById1**](docs/RelationalBackupAPIApi.md#getdetailbackupbyid1) | **Get** /v1/backups/detail/{backupId} | 
*RelationalBackupAPIApi* | [**GetFreeBackupUsage1**](docs/RelationalBackupAPIApi.md#getfreebackupusage1) | **Get** /v1/backups/free-backup | 
*RelationalBackupAPIApi* | [**GetListBackups1**](docs/RelationalBackupAPIApi.md#getlistbackups1) | **Get** /v1/backups | 
*RelationalBackupAPIApi* | [**GetListBackupsByInstanceId1**](docs/RelationalBackupAPIApi.md#getlistbackupsbyinstanceid1) | **Get** /v1/backups/insId/{instanceId} | 
*RelationalBackupAPIApi* | [**RestoreBackup1**](docs/RelationalBackupAPIApi.md#restorebackup1) | **Post** /v1/backups/{id}/restore | 
*RelationalBackupStorageAPIApi* | [**CreateRelationalBackUpStorage**](docs/RelationalBackupStorageAPIApi.md#createrelationalbackupstorage) | **Post** /v1/payment/backup-storages | 
*RelationalBackupStorageAPIApi* | [**DeleteBackupStorage1**](docs/RelationalBackupStorageAPIApi.md#deletebackupstorage1) | **Post** /v1/backup-storages/actions/deletions | 
*RelationalBackupStorageAPIApi* | [**GetListBackupStorage1**](docs/RelationalBackupStorageAPIApi.md#getlistbackupstorage1) | **Get** /v1/backup-storages/information | 
*RelationalBackupStorageAPIApi* | [**GetListQuotaPackage1**](docs/RelationalBackupStorageAPIApi.md#getlistquotapackage1) | **Get** /v1/backup-storages | 
*RelationalBackupStorageAPIApi* | [**RenewBackupStorage1**](docs/RelationalBackupStorageAPIApi.md#renewbackupstorage1) | **Post** /v1/backup-storages/actions/renew | 
*RelationalBackupStorageAPIApi* | [**ResizeBackupStorage1**](docs/RelationalBackupStorageAPIApi.md#resizebackupstorage1) | **Post** /v1/backup-storages/actions/resize | 
*RelationalConfigurationGroupAPIApi* | [**CreateConfig1**](docs/RelationalConfigurationGroupAPIApi.md#createconfig1) | **Post** /v1/configurations/create | 
*RelationalConfigurationGroupAPIApi* | [**DeleteConfigs1**](docs/RelationalConfigurationGroupAPIApi.md#deleteconfigs1) | **Delete** /v1/configurations/delete | 
*RelationalConfigurationGroupAPIApi* | [**GetConfigParams1**](docs/RelationalConfigurationGroupAPIApi.md#getconfigparams1) | **Get** /v1/configurations/params | 
*RelationalConfigurationGroupAPIApi* | [**GetConfigsById1**](docs/RelationalConfigurationGroupAPIApi.md#getconfigsbyid1) | **Get** /v1/configurations/id | 
*RelationalConfigurationGroupAPIApi* | [**GetListConfigs1**](docs/RelationalConfigurationGroupAPIApi.md#getlistconfigs1) | **Get** /v1/configurations | 
*RelationalConfigurationGroupAPIApi* | [**UpdateConfig1**](docs/RelationalConfigurationGroupAPIApi.md#updateconfig1) | **Put** /v1/configurations/update | 
*RelationalDatabaseAPIApi* | [**CreateRelationalDatabaseInstance**](docs/RelationalDatabaseAPIApi.md#createrelationaldatabaseinstance) | **Post** /v1/payment/database-instances | 
*RelationalDatabaseAPIApi* | [**CreateRelationalDatabaseInstanceReplica**](docs/RelationalDatabaseAPIApi.md#createrelationaldatabaseinstancereplica) | **Post** /v1/database-instances/{id}/create-replicas | 
*RelationalDatabaseAPIApi* | [**DeleteDatabaseInstances1**](docs/RelationalDatabaseAPIApi.md#deletedatabaseinstances1) | **Post** /v1/database-instances/{instanceId}/delete | 
*RelationalDatabaseAPIApi* | [**DetachReplica1**](docs/RelationalDatabaseAPIApi.md#detachreplica1) | **Post** /v1/database-instances/{instanceId}/detach-replica | 
*RelationalDatabaseAPIApi* | [**ExchangePocResource1**](docs/RelationalDatabaseAPIApi.md#exchangepocresource1) | **Post** /v1/database-instances/exchange-poc | 
*RelationalDatabaseAPIApi* | [**GetAllDatastore1**](docs/RelationalDatabaseAPIApi.md#getalldatastore1) | **Get** /v1/database-instances/datastore | 
*RelationalDatabaseAPIApi* | [**GetAllInstanceFamily1**](docs/RelationalDatabaseAPIApi.md#getallinstancefamily1) | **Get** /v1/database-instances/families | 
*RelationalDatabaseAPIApi* | [**GetDBInstanceConfig1**](docs/RelationalDatabaseAPIApi.md#getdbinstanceconfig1) | **Get** /v1/database-instances/configuration | 
*RelationalDatabaseAPIApi* | [**GetDatabaseInstancesById1**](docs/RelationalDatabaseAPIApi.md#getdatabaseinstancesbyid1) | **Get** /v1/database-instances/id/{dbInstanceId} | 
*RelationalDatabaseAPIApi* | [**GetDatabaseInstancesByUser1**](docs/RelationalDatabaseAPIApi.md#getdatabaseinstancesbyuser1) | **Get** /v1/database-instances | 
*RelationalDatabaseAPIApi* | [**GetEngine1**](docs/RelationalDatabaseAPIApi.md#getengine1) | **Get** /v1/database-instances/engine | 
*RelationalDatabaseAPIApi* | [**GetFlavorCodes1**](docs/RelationalDatabaseAPIApi.md#getflavorcodes1) | **Get** /v1/database-instances/flavor_zones/codes | 
*RelationalDatabaseAPIApi* | [**GetFlavors1**](docs/RelationalDatabaseAPIApi.md#getflavors1) | **Get** /v1/database-instances/flavors | 
*RelationalDatabaseAPIApi* | [**GetHistoryDB1**](docs/RelationalDatabaseAPIApi.md#gethistorydb1) | **Get** /v1/database-instances/{instanceId}/histories | 
*RelationalDatabaseAPIApi* | [**GetListNetwork1**](docs/RelationalDatabaseAPIApi.md#getlistnetwork1) | **Get** /v1/database-instances/networks | 
*RelationalDatabaseAPIApi* | [**GetListReplicas1**](docs/RelationalDatabaseAPIApi.md#getlistreplicas1) | **Get** /v1/database-instances/{replicaSourceId}/replicas | 
*RelationalDatabaseAPIApi* | [**GetSecurityRules**](docs/RelationalDatabaseAPIApi.md#getsecurityrules) | **Put** /v1/database-instances/{instanceId}/secrules | 
*RelationalDatabaseAPIApi* | [**GetSecurityRules2**](docs/RelationalDatabaseAPIApi.md#getsecurityrules2) | **Get** /v1/database-instances/{instanceId}/secrules | 
*RelationalDatabaseAPIApi* | [**GetVolumeTypes2**](docs/RelationalDatabaseAPIApi.md#getvolumetypes2) | **Get** /v1/database-instances/volume/types | 
*RelationalDatabaseAPIApi* | [**ListNetwork1**](docs/RelationalDatabaseAPIApi.md#listnetwork1) | **Get** /v1/database-instances/networks/subnets | 
*RelationalDatabaseAPIApi* | [**RenewResource1**](docs/RelationalDatabaseAPIApi.md#renewresource1) | **Post** /v1/database-instances/{resourceId}/renew-resource | 
*RelationalDatabaseAPIApi* | [**ResizeInstance1**](docs/RelationalDatabaseAPIApi.md#resizeinstance1) | **Post** /v1/database-instances/{instanceId}/resize-instance | 
*RelationalDatabaseAPIApi* | [**ResizeStorage**](docs/RelationalDatabaseAPIApi.md#resizestorage) | **Post** /v1/database-instances/{instanceId}/resize-storage | 
*RelationalDatabaseAPIApi* | [**RestartDatabaseInstances1**](docs/RelationalDatabaseAPIApi.md#restartdatabaseinstances1) | **Post** /v1/database-instances/{instanceId}/reboot | 
*RelationalDatabaseAPIApi* | [**StartDatabaseInstances1**](docs/RelationalDatabaseAPIApi.md#startdatabaseinstances1) | **Post** /v1/database-instances/{instanceId}/start | 
*RelationalDatabaseAPIApi* | [**StopDatabaseInstances1**](docs/RelationalDatabaseAPIApi.md#stopdatabaseinstances1) | **Post** /v1/database-instances/{instanceId}/shutdown | 
*RelationalDatabaseAPIApi* | [**UpdateDatabaseConfigGroup1**](docs/RelationalDatabaseAPIApi.md#updatedatabaseconfiggroup1) | **Put** /v1/database-instances/{instanceId}/update/config-group | 
*RelationalDatabaseAPIApi* | [**UpdateDatabaseSetting1**](docs/RelationalDatabaseAPIApi.md#updatedatabasesetting1) | **Put** /v1/database-instances/{instanceId}/update/setting | 
*ResourceControllerApi* | [**GetListResource**](docs/ResourceControllerApi.md#getlistresource) | **Get** /v1/resources | 
*ResourceControllerApi* | [**UpdateAutoRenew**](docs/ResourceControllerApi.md#updateautorenew) | **Put** /v1/resources/{artifactId}/autoRenew | 

## Documentation For Models

 - [ActionDbInstancesResponse](docs/ActionDbInstancesResponse.md)
 - [AddressInfo](docs/AddressInfo.md)
 - [AutoRenewInfo](docs/AutoRenewInfo.md)
 - [AutoRenewRequest](docs/AutoRenewRequest.md)
 - [BackupInfo](docs/BackupInfo.md)
 - [BackupInfoGatewayResponse](docs/BackupInfoGatewayResponse.md)
 - [BackupStorageDetail](docs/BackupStorageDetail.md)
 - [BillingElement](docs/BillingElement.md)
 - [CalculatePriceRequest](docs/CalculatePriceRequest.md)
 - [CalculatePriceResponse](docs/CalculatePriceResponse.md)
 - [ConfigInfoGatewayResponse](docs/ConfigInfoGatewayResponse.md)
 - [ConfigurationParamInfo](docs/ConfigurationParamInfo.md)
 - [CreateBackupResponse](docs/CreateBackupResponse.md)
 - [CreateKafkaClusterRequest](docs/CreateKafkaClusterRequest.md)
 - [DatabaseInstancesGateway](docs/DatabaseInstancesGateway.md)
 - [DatabaseInstancesGatewayResponse](docs/DatabaseInstancesGatewayResponse.md)
 - [DbBackupPackageDetail](docs/DbBackupPackageDetail.md)
 - [DbBackupPackageResponse](docs/DbBackupPackageResponse.md)
 - [DbInstanceInfo](docs/DbInstanceInfo.md)
 - [DbInstancesHistoryGatewayResponse](docs/DbInstancesHistoryGatewayResponse.md)
 - [DeleteBackupResponse](docs/DeleteBackupResponse.md)
 - [DeleteConfigResponse](docs/DeleteConfigResponse.md)
 - [Engine](docs/Engine.md)
 - [EngineVersion](docs/EngineVersion.md)
 - [FilterRequest](docs/FilterRequest.md)
 - [FlavorCode](docs/FlavorCode.md)
 - [FlavorInfo](docs/FlavorInfo.md)
 - [FreeBackupStorageInfo](docs/FreeBackupStorageInfo.md)
 - [HistoryResponse](docs/HistoryResponse.md)
 - [InstanceEntity](docs/InstanceEntity.md)
 - [InstanceFamily](docs/InstanceFamily.md)
 - [ItemConfigInfo](docs/ItemConfigInfo.md)
 - [NetworkResponse](docs/NetworkResponse.md)
 - [NetworkResponseV2](docs/NetworkResponseV2.md)
 - [OrderResponse](docs/OrderResponse.md)
 - [PageObject](docs/PageObject.md)
 - [PageObjectWithLimitSize](docs/PageObjectWithLimitSize.md)
 - [PropertyPrice](docs/PropertyPrice.md)
 - [ResourceBilling](docs/ResourceBilling.md)
 - [ResourcesBillingInfo](docs/ResourcesBillingInfo.md)
 - [SecurityGroupEntity](docs/SecurityGroupEntity.md)
 - [SecurityGroupRuleEntity](docs/SecurityGroupRuleEntity.md)
 - [SubnetResponse](docs/SubnetResponse.md)
 - [TinyConfigInfo](docs/TinyConfigInfo.md)
 - [TinyDbInstanceInfo](docs/TinyDbInstanceInfo.md)
 - [VolumeTypeGatewayResponse](docs/VolumeTypeGatewayResponse.md)
 - [VolumeTypeInfo](docs/VolumeTypeInfo.md)
 - [WrapContentBackupInfo](docs/WrapContentBackupInfo.md)
 - [WrapContentBackupInfoGatewayResponse](docs/WrapContentBackupInfoGatewayResponse.md)
 - [WrapContentConfigInfoGatewayResponse](docs/WrapContentConfigInfoGatewayResponse.md)
 - [WrapContentCreateBackupResponse](docs/WrapContentCreateBackupResponse.md)
 - [WrapContentDatabaseInstancesGatewayResponse](docs/WrapContentDatabaseInstancesGatewayResponse.md)
 - [WrapContentDbInstanceInfo](docs/WrapContentDbInstanceInfo.md)
 - [WrapContentDbInstancesHistoryGatewayResponse](docs/WrapContentDbInstancesHistoryGatewayResponse.md)
 - [WrapContentFreeBackupStorageInfo](docs/WrapContentFreeBackupStorageInfo.md)
 - [WrapContentItemConfigInfo](docs/WrapContentItemConfigInfo.md)
 - [WrapContentListActionDbInstancesResponse](docs/WrapContentListActionDbInstancesResponse.md)
 - [WrapContentListBackupInfo](docs/WrapContentListBackupInfo.md)
 - [WrapContentListBackupStorageDetail](docs/WrapContentListBackupStorageDetail.md)
 - [WrapContentListConfigurationParamInfo](docs/WrapContentListConfigurationParamInfo.md)
 - [WrapContentListDbBackupPackageResponse](docs/WrapContentListDbBackupPackageResponse.md)
 - [WrapContentListDeleteBackupResponse](docs/WrapContentListDeleteBackupResponse.md)
 - [WrapContentListDeleteConfigResponse](docs/WrapContentListDeleteConfigResponse.md)
 - [WrapContentListEngineVersion](docs/WrapContentListEngineVersion.md)
 - [WrapContentListFlavorInfo](docs/WrapContentListFlavorInfo.md)
 - [WrapContentListInstanceFamily](docs/WrapContentListInstanceFamily.md)
 - [WrapContentListItemConfigInfo](docs/WrapContentListItemConfigInfo.md)
 - [WrapContentListNetworkResponse](docs/WrapContentListNetworkResponse.md)
 - [WrapContentListNetworkResponseV2](docs/WrapContentListNetworkResponseV2.md)
 - [WrapContentListOrderResponse](docs/WrapContentListOrderResponse.md)
 - [WrapContentListSecurityGroupRuleEntity](docs/WrapContentListSecurityGroupRuleEntity.md)
 - [WrapContentMapStringString](docs/WrapContentMapStringString.md)
 - [WrapContentObject](docs/WrapContentObject.md)
 - [WrapContentResourcesBillingInfo](docs/WrapContentResourcesBillingInfo.md)
 - [WrapContentSetEngine](docs/WrapContentSetEngine.md)
 - [WrapContentSetFlavorCode](docs/WrapContentSetFlavorCode.md)
 - [WrapContentSetString](docs/WrapContentSetString.md)
 - [WrapContentVolumeTypeGatewayResponse](docs/WrapContentVolumeTypeGatewayResponse.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


