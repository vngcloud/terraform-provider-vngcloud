/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Create Bandwidth Request
type CreateBandwidthRequest struct {
	// Description of Bandwidth
	Description string `json:"description,omitempty"`
	// Ip of Bandwidth
	Ip string `json:"ip"`
	// Package Name Dom of Bandwidth
	PackageNameDom string `json:"packageNameDom"`
	// Package Name Int of Bandwidth
	PackageNameInt string `json:"packageNameInt"`
	// Package Name Pay Type of Bandwidth
	PackageNamePayG string `json:"packageNamePayG"`
	// Package Type of Bandwidth
	PackageType string `json:"packageType"`
	// Resource Type of Bandwidth
	ResourceType string `json:"resourceType"`
}
