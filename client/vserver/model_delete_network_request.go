/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type DeleteNetworkRequest struct {
	Extra interface{} `json:"extra,omitempty"`
	// Id of Network to be deleted.
	NetworkId string `json:"networkId"`
}
