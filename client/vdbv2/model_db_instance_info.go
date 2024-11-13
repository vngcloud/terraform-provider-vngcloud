/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

import (
	"time"
)

type DbInstanceInfo struct {
	Id                   string                   `json:"id,omitempty"`
	DbBackendId          int32                    `json:"dbBackendId,omitempty"`
	Name                 string                   `json:"name,omitempty"`
	Status               string                   `json:"status,omitempty"`
	Ram                  int32                    `json:"ram,omitempty"`
	Vcpus                int32                    `json:"vcpus,omitempty"`
	DatastoreType        string                   `json:"datastoreType,omitempty"`
	PriceKey             string                   `json:"priceKey,omitempty"`
	ZoneType             string                   `json:"zoneType,omitempty"`
	EngineGroup          int32                    `json:"engineGroup,omitempty"`
	DatastoreVersion     string                   `json:"datastoreVersion,omitempty"`
	Region               string                   `json:"region,omitempty"`
	VolumeType           string                   `json:"volumeType,omitempty"`
	VolumeSize           int32                    `json:"volumeSize,omitempty"`
	VolumeUsed           float32                  `json:"volumeUsed,omitempty"`
	Hostname             string                   `json:"hostname,omitempty"`
	Ip                   []string                 `json:"ip,omitempty"`
	Created              string                   `json:"created,omitempty"`
	Updated              string                   `json:"updated,omitempty"`
	Role                 string                   `json:"role,omitempty"`
	Address              map[string][]AddressInfo `json:"address,omitempty"`
	Configuration        *TinyConfigInfo          `json:"configuration,omitempty"`
	PublicAccess         bool                     `json:"publicAccess,omitempty"`
	BackupAuto           bool                     `json:"backupAuto,omitempty"`
	BackupDuration       int32                    `json:"backupDuration,omitempty"`
	BackupTime           string                   `json:"backupTime,omitempty"`
	Port                 int32                    `json:"port,omitempty"`
	SecurityGroup        []SecurityGroupEntity    `json:"securityGroup,omitempty"`
	VolumeTypeZoneId     string                   `json:"volumeTypeZoneId,omitempty"`
	QuotaPackageId       string                   `json:"quotaPackageId,omitempty"`
	StartTime            time.Time                `json:"startTime,omitempty"`
	EndTime              time.Time                `json:"endTime,omitempty"`
	DisabledAt           time.Time                `json:"disabledAt,omitempty"`
	Period               int32                    `json:"period,omitempty"`
	Cost                 float64                  `json:"cost,omitempty"`
	Poc                  bool                     `json:"poc,omitempty"`
	License              string                   `json:"license,omitempty"`
	Bandwidth            int32                    `json:"bandwidth,omitempty"`
	PackageName          string                   `json:"packageName,omitempty"`
	BackupUsage          float32                  `json:"backupUsage,omitempty"`
	FreeBackupSize       int32                    `json:"freeBackupSize,omitempty"`
	ReplicaSourceId      string                   `json:"replicaSourceId,omitempty"`
	ReplicaSourceName    string                   `json:"replicaSourceName,omitempty"`
	Replicas             []string                 `json:"replicas,omitempty"`
	RedisPasswordEnabled bool                     `json:"redisPasswordEnabled,omitempty"`
	SharedBy             string                   `json:"sharedBy,omitempty"`
	Owner                bool                     `json:"owner,omitempty"`
	SharedActions        []string                 `json:"sharedActions,omitempty"`
	EnableAutoRenew      bool                     `json:"enableAutoRenew,omitempty"`
	AutoRenewPeriod      int32                    `json:"autoRenewPeriod,omitempty"`
	ZoneUUID             string                   `json:"zoneUUID,omitempty"`
	TrialDb              bool                     `json:"trialDb,omitempty"`
	SubnetId             string                   `json:"subnetId,omitempty"`
	ProjectId            string                   `json:"projectId,omitempty"`
	SkuFlavor            string                   `json:"skuFlavor,omitempty"`
	SkuVolume            string                   `json:"skuVolume,omitempty"`
	ConfigId             string                   `json:"configId,omitempty"`
	Deleted              bool                     `json:"deleted,omitempty"`
}