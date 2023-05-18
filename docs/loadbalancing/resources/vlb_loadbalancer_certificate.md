
---
subcategory: "LB (Load Balancing)"
layout: "vngcloud"
page_title: "VNG: vng_lb_CA"
description: |-
ResourceCA is a Terraform resource that represents a certificate authority (CA). It can be used to import, and delete CAs in a target system.
---

# vngcloud_vlb_certificate (Resource)

ResourceCA is a Terraform resource that represents a certificate authority (CA). It can be used to import and delete CAs in a target system.

## Example Usage

```terraform

resource "vngcloud_vlb_certificate" "example" {
  project_id          = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  name                = "my-ca"
  type                = "TLS/SSL"
  certificate         = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
  private_key         = "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n"
  certificate_chain   = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n"
  expiration          = "2024-02-27T03:52:55.045Z"
}

```
Note: You can use the `file()` function to read the contents of a file into a string.
For example, the following code reads the contents of a file named `my-ca.pem` into the `certificate` argument:

```terraform
resource "vngcloud_vlb_certificate" "example" {
  # ...
  certificate = file("my-ca.pem")
}
```

## Argument Reference

The following arguments are supported:


* `project_id` -  (String, Required) This field represents the ID of the project in which the load balancer is created.
* `name` - (String, Required) The name of the CA.
* `type` -(String, Required) The type of the CA. It can be `TLS/SSL` or `CA`.
* `certificate` - (String, Required) The contents of the CA's certificate file.
* `private_key` - (String, Optional) The contents of the CA's private key file.
* `certificate_chain` - (String, Optional) The contents of the CA's certificate chain file.
* `passphrase` - (String, Optional) The passphrase used to encrypt the CA's private key.
* `expiration` - (String, Optional)The date and time at which the CA's certificate will expire.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - (String) The UUID of the CA.

## IAM Policy
### Create:
In order to **create Certificate**, user must have been granted permissions below:
- ImportCertificateAuthority
- GetCertificateAuthority

### Delete
In order to **delete Certificate**, user must have been granted permissions below:
- DeleteCertificateAuthority
