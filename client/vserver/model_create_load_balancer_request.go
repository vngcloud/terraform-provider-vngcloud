/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateLoadBalancerRequest struct {
	// Algorithm of the pool. The algorithm can be \"ROUND_ROBIN\" or \"LEAST_CONNECTIONS\" or \"SOURCE_IP\"
	Algorithm string `json:"algorithm,omitempty"`
	// Allowed cidr.
	AllowedCidrs string `json:"allowedCidrs,omitempty"`
	// List of certificate authority
	CertificateAuthorities []string `json:"certificateAuthorities,omitempty"`
	// Default certificate authority that the listener will use
	DefaultCertificateAuthority string `json:"defaultCertificateAuthority,omitempty"`
	// Health check method for the HTTP protocol.
	HealthCheckMethod string `json:"healthCheckMethod,omitempty"`
	// Health check path for the HTTP protocol.
	HealthCheckPath string `json:"healthCheckPath,omitempty"`
	// Protocol for performing health check. The protocol can be TCP or HTTP.
	HealthCheckProtocol string `json:"healthCheckProtocol,omitempty"`
	// Healthy threshold. The value must be in range from 2 to 10.
	HealthyThreshold int64 `json:"healthyThreshold,omitempty"`
	// Health check interval. The value must be from 5 to 3600 seconds.
	Interval int64 `json:"interval,omitempty"`
	// Name of the listener. Only letters (a-z, A-Z, 0-9, '_', '.') are allowed and your input data must be between 6 and 20 characters.
	ListenerName string `json:"listenerName,omitempty"`
	// Protocol of the listener.
	ListenerProtocol string `json:"listenerProtocol,omitempty"`
	// Port of the listener.
	ListenerProtocolPort int32 `json:"listenerProtocolPort,omitempty"`
	// List of members of the pool.
	Members []CreateMemberRequest `json:"members,omitempty"`
	// Load balancer's name. Only letters (a-z, A-Z, 0-9, '_', '.') are allowed and your input data must be between 6 and 20 characters.
	Name string `json:"name"`
	// Package ID of the load balancer.
	PackageId string `json:"packageId,omitempty"`
	// Name of the pool. Only letters (a-z, A-Z, 0-9, '_', '.') are allowed and your input data must be between 6 and 20 characters.
	PoolName string `json:"poolName,omitempty"`
	// Protocol of the pool.
	PoolProtocol string `json:"poolProtocol,omitempty"`
	// Schema of the load balancer, it may be Internet or Internal.
	Scheme string `json:"scheme"`
	// Enable sticky sessions.
	Stickiness bool `json:"stickiness,omitempty"`
	// Subnet ID for the load balancer.
	SubnetId string `json:"subnetId,omitempty"`
	// Health check success code for HTTP health check protocol.
	SuccessCode string `json:"successCode,omitempty"`
	// Timeout of health check. The value must be from 2 to 120 seconds
	Timeout int64 `json:"timeout,omitempty"`
	// Idle timeout of client. The value can be in range from 1 to 3600 seconds
	TimeoutClient int32 `json:"timeoutClient,omitempty"`
	// Idle timeout of connection. The value can be in range from 1 to 3600 seconds
	TimeoutConnection int32 `json:"timeoutConnection,omitempty"`
	// Idle timeout of member. The value can be in range from 1 to 3600 seconds
	TimeoutMember int32 `json:"timeoutMember,omitempty"`
	// Type of the load balancer. It may be Layer 4 or Layer 7
	Type_ string `json:"type,omitempty"`
	// Unhealthy threshold. The value must be in range from 2 to 10.
	UnhealthyThreshold int64 `json:"unhealthyThreshold,omitempty"`
}