/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Create Network Request
type CreateNetworkRequest struct {
	// Subnet of Network (/16).
	Cidr string `json:"cidr"`
	Extra *interface{} `json:"extra,omitempty"`
	// Name of the Network
	Name string `json:"name"`
	// Id of default route table.
	RouteTableDefaultId string `json:"routeTableDefaultId,omitempty"`
}
