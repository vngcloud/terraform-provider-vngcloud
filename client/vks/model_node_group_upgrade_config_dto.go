/*
 * vks-api API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0-SNAPSHOT
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package vks

type NodeGroupUpgradeConfigDto struct {
	Strategy       string `json:"strategy"`
	MaxSurge       int32  `json:"maxSurge"`
	MaxUnavailable int32  `json:"maxUnavailable"`
}
