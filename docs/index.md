---
page_title: "VNG Cloud Provider"
description: |-
The VNG Cloud provider is used to interact with the resources supported by VNG Cloud. The provider needs to be configured with the clientID and client secret before it can be used.
---

# VNGCloud Provider



## Example Usage

```terraform
terraform {
  required_providers {
    vngcloud = {
      source  = "vngcloud/vngcloud"
      version = "= 1.0.0"
    }
  }
}
provider "vngcloud" {
  token_url        = "https://iamapis.vngcloud.vn/accounts-api/v2/auth/token"
  client_id        = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  client_secret    = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  vserver_base_url = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway"
  vlb_base_url     = "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway"
}
```
## Argument Reference

The following arguments are supported:

* `client_id` - (Required) This is your clientID to authenticate with VNG Cloud.  Alternatively, this can also be specified using environment variables ordered by precedence: `CLIENT_ID`

* `client_secret` - (Required) This is your client secret to authenticate with VNG Cloud.  Alternatively, this can also be specified using environment variables ordered by precedence: `CLIENT_SECRET`

* `token_url` - (Required) This is endpoint to authentication with VNG Cloud. Alternatively, this can also be specified using environment variables ordered by precedence: `TOKEN_ADDRESS`. Suggested value is [https://iamapis.vngcloud.vn/accounts-api/v2/auth/token](https://iamapis.vngcloud.vn/accounts-api/v2/auth/token)

* `vserver_base_url` - (Required) This is endpoint to interactive with VNG Cloud's resource. Alternatively, this can also be specified using environment variables ordered by precedence: `VSERVER_BASE_URL`. Suggested value is [https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway](https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway)

* `vlb_base_url` - (Required) This is endpoint to interactive with VNG Cloud's resource. Alternatively, this can also be specified using environment variables ordered by precedence: `VLB_BASE_URL`. Suggested value is [https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway](https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway)
