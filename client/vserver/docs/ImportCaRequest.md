# ImportCaRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Certificate of the CA | [optional] [default to null]
**CertificateChain** | **string** | A certificate chain is an ordered list of certificates, containing an SSL/TLS Certificate and Certificate Authority (CA) Certificates, that enable the receiver to verify that the sender and all CA&#39;s are trustworthy. | [optional] [default to null]
**Expiration** | [**time.Time**](time.Time.md) | Expiration of CA | [optional] [default to null]
**Name** | **string** | Name of the CA | [default to null]
**Passphrase** | **string** | The password of the private key. | [optional] [default to null]
**PrivateKey** | **string** | Private key of the CA | [optional] [default to null]
**Type_** | **string** | Type of the certificate ( CA or TLS/SSL )  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


