/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type BandwidthTrafficHistoryModel struct {
	BandwidthUuid string    `json:"bandwidthUuid,omitempty"`
	CreatedAt     string    `json:"createdAt,omitempty"`
	Id            int32     `json:"id,omitempty"`
	RemainVolume  int64     `json:"remainVolume,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	UsedVolume    int64     `json:"usedVolume,omitempty"`
}
