package vserver

// SecondarySubnetDto is a secondary CIDR range that belongs to a (primary) subnet.
type SecondarySubnetDto struct {
	Cidr string `json:"cidr,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
