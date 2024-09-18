/*
 * vDB API
 *
 * API for working with vDB Instances (Relational Databases, MemoryStore Databases and Kafka Clusters).
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vdbv2

type DbBackupPackageDetail struct {
	PackageId      int32  `json:"packageId,omitempty"`
	Sku            string `json:"sku,omitempty"`
	ResourceTypeId int32  `json:"resourceTypeId,omitempty"`
	ResourceName   string `json:"resourceName,omitempty"`
	Description    string `json:"description,omitempty"`
	PackageName    string `json:"packageName,omitempty"`
	PackageQuota   string `json:"packageQuota,omitempty"`
	PackageConfig  string `json:"packageConfig,omitempty"`
	Price          int64  `json:"price,omitempty"`
}
