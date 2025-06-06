/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type BackupStorageDetail struct {
	Id                    string  `json:"id,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Usage                 float32 `json:"usage,omitempty"`
	Quota                 int32   `json:"quota,omitempty"`
	Status                string  `json:"status,omitempty"`
	EngineGroup           int32   `json:"engineGroup,omitempty"`
	BackupPackageId       string  `json:"backupPackageId,omitempty"`
	BackupPackageName     string  `json:"backupPackageName,omitempty"`
	SkuDbaasBackupStorage string  `json:"skuDbaasBackupStorage,omitempty"`
}
