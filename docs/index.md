---
page_title: "VNG Cloud Provider"
description: |-
    The VNG Cloud provider is used to interact with the resources supported by VNG Cloud. The provider needs to be configured with the clientID and client secret before it can be used.
---

# VNGCloud Provider



## Example Usage

```terraform
provider "vngcloud" {
    token_url = "https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token"
    client_id = "client_id"
    client_secret = "client_secret"
    base_url = "https://vserverapi.vngcloud.vn/vserver-gateway"
}
```
## Argument Reference

The following arguments are supported:

* `client_id` - (Required) This is your clientID to authenticate with VNG Cloud.  Alternatively, this can also be specified using environment variables ordered by precedence: `CLIENT_ID`

* `client_secret` - (Required) This is your client secret to authenticate with VNG Cloud.  Alternatively, this can also be specified using environment variables ordered by precedence: `CLIENT_SECRET`

* `token_url` - (Required) This is endpoint to authencation with VNG Cloud. Alternatively, this can also be specified using environment variables ordered by precedence: `TOKEN_ADDRESS`. Suggested value is [https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token](https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token)
  
* `api_endpoint` - (Required) This is endpoint to interactive with VNG Cloud's resource. Alternatively, this can also be specified using environment variables ordered by precedence: `BASE_URL`. Suggested value is [https://vserverapi.vngcloud.vn/vserver-gateway](https://vserverapi.vngcloud.vn/vserver-gateway)