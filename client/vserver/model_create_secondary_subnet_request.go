package vserver

// CreateSecondarySubnetRequest is one element of the secondary subnet list sent when syncing a subnet.
type CreateSecondarySubnetRequest struct {
	Name string `json:"name,omitempty"`
	Cidr string `json:"cidr,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
