/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancing

type CreateLoadBalancerRequestV2 struct {
	// Listener of the load balancer.
	Listener *LbListener `json:"listener"`
	// Load balancer's name. Only letters (a-z, A-Z, 0-9, '_', '.') are allowed and your input data must be between 6 and 20 characters.
	Name string `json:"name"`
	// Package ID of the load balancer.
	PackageId string `json:"packageId"`
	// Pool of the load balancer.
	Pool *LbPool `json:"pool"`
	// Schema of the load balancer, it may be Internet or Internal.
	Scheme string `json:"scheme"`
	// Subnet ID for the load balancer.
	SubnetId string `json:"subnetId"`
	// Type of the load balancer. It may be Layer 4 or Layer 7
	Type_ string `json:"type"`
}
