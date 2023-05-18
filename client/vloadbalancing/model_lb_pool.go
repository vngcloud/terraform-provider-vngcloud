/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancing

type LbPool struct {
	// Algorithm of the pool. The algorithm can be \"ROUND_ROBIN\" or \"LEAST_CONNECTIONS\" or \"SOURCE_IP\"
	Algorithm string `json:"algorithm"`
	// Health check method for the HTTP protocol.
	HealthCheckMethod string `json:"healthCheckMethod"`
	// Health check path for the HTTP protocol.
	HealthCheckPath string `json:"healthCheckPath"`
	// Protocol for performing health check.
	HealthCheckProtocol string `json:"healthCheckProtocol"`
	// Healthy threshold. The value must be in range from 2 to 10.
	HealthyThreshold int64 `json:"healthyThreshold"`
	// Health check interval. The value must be from 5 to 3600 seconds.
	Interval int64 `json:"interval"`
	// List of members of the pool.
	Members []CreateMemberRequest `json:"members"`
	// Name of the pool. Only letters (a-z, A-Z, 0-9, '_', '.') are allowed and your input data must be between 6 and 20 characters.
	PoolName string `json:"poolName"`
	// Protocol of the pool.
	PoolProtocol string `json:"poolProtocol"`
	// Enable sticky sessions.
	Stickiness bool `json:"stickiness"`
	// Health check success code for HTTP health check protocol.
	SuccessCode string `json:"successCode"`
	// Timeout of health check. The value must be from 2 to 120 seconds
	Timeout int64 `json:"timeout"`
	// Unhealthy threshold. The value must be in range from 2 to 10.
	UnhealthyThreshold int64 `json:"unhealthyThreshold"`
}
