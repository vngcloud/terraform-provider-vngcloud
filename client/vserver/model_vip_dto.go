/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type VipDto struct {
	AddressPairIps []string  `json:"addressPairIps"`
	CreatedAt      time.Time `json:"createdAt"`
	Description    string    `json:"description"`
	IpAddress      string    `json:"ipAddress"`
	Name           string    `json:"name"`
	NetworkCIDR    string    `json:"networkCIDR"`
	NetworkId      string    `json:"networkId"`
	NetworkName    string    `json:"networkName"`
	SubnetCIDR     string    `json:"subnetCIDR"`
	SubnetId       string    `json:"subnetId"`
	SubnetName     string    `json:"subnetName"`
	Uuid           string    `json:"uuid"`
}
