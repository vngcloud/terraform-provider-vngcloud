# CreateServerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachFloating** | **bool** | Attach floating IP | [optional] [default to null]
**DataDiskEncryptionType** | **string** | Type encryption of data volume | [optional] [default to null]
**DataDiskName** | **string** | Name of data volume | [optional] [default to null]
**DataDiskSize** | **int32** |  | [optional] [default to null]
**DataDiskTypeId** | **string** | Id of data volume type | [optional] [default to null]
**EncryptionVolume** | **bool** | Encryption volume | [default to null]
**ExpirePassword** | **bool** | Skip change password: false, else: true | [optional] [default to null]
**FlavorId** | **string** | Id of the flavor | [default to null]
**ImageId** | **string** | Id of the image | [default to null]
**Name** | **string** | Name of the server | [default to null]
**NetworkId** | **string** | Id of the network | [default to null]
**OsLicence** | **bool** | Licence of OS | [optional] [default to null]
**RootDiskEncryptionType** | **string** | Type encryption of boot volume | [optional] [default to null]
**RootDiskSize** | **int32** | Size of boot volume | [default to null]
**RootDiskTypeId** | **string** | Id of boot volume type | [default to null]
**SecurityGroup** | **[]string** | Id of the SecGroups | [optional] [default to null]
**ServerGroupId** | **string** | Server group id | [optional] [default to null]
**SshKeyId** | **string** | Id of SSH key | [optional] [default to null]
**SubnetId** | **string** | Id of the subnet | [default to null]
**UserData** | **string** | User data | [optional] [default to null]
**UserDataBase64Encoded** | **bool** | User data has already been base64 encoded | [optional] [default to null]
**UserName** | **string** | name of user | [optional] [default to null]
**UserPassword** | **string** | password of user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


