/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

// Create Server Request
type CreateServerRequest struct {
	// Attach floating IP
	AttachFloating bool `json:"attachFloating,omitempty"`
	// Type encryption of data volume
	DataDiskEncryptionType string `json:"dataDiskEncryptionType,omitempty"`
	// Name of data volume
	DataDiskName string `json:"dataDiskName,omitempty"`
	DataDiskSize int32  `json:"dataDiskSize,omitempty"`
	// Id of data volume type
	DataDiskTypeId string `json:"dataDiskTypeId,omitempty"`
	// Encryption volume
	EncryptionVolume bool `json:"encryptionVolume"`
	// Skip change password: false, else: true
	ExpirePassword bool `json:"expirePassword,omitempty"`
	// Id of the flavor
	FlavorId string `json:"flavorId"`
	// Id of the image
	ImageId string `json:"imageId"`
	// Name of the server
	Name string `json:"name"`
	// Id of the network
	NetworkId string `json:"networkId"`
	// Licence of OS
	OsLicence bool `json:"osLicence,omitempty"`
	// Type encryption of boot volume
	RootDiskEncryptionType string `json:"rootDiskEncryptionType,omitempty"`
	// Size of boot volume
	RootDiskSize int32 `json:"rootDiskSize"`
	// Id of boot volume type
	RootDiskTypeId string `json:"rootDiskTypeId"`
	// Id of the SecGroups
	SecurityGroup []string `json:"securityGroup,omitempty"`
	// Server group id
	ServerGroupId string `json:"serverGroupId,omitempty"`
	// Id of SSH key
	SshKeyId string `json:"sshKeyId,omitempty"`
	// Id of the subnet
	SubnetId string `json:"subnetId"`
	// User data
	UserData string `json:"userData,omitempty"`
	// User data has already been base64 encoded
	UserDataBase64Encoded bool `json:"userDataBase64Encoded,omitempty"`
	// name of user
	UserName string `json:"userName,omitempty"`
	// password of user
	UserPassword string `json:"userPassword,omitempty"`
	// Zone of flavor. You can choose if having multiple zone
	HostGroupId string `json:"hostGroupId,omitempty"`
	// is POC
	IsPoc bool `json:"isPoc,omitempty"`
	// Zone id
	ZoneId string `json:"zoneId,omitempty"`
}
