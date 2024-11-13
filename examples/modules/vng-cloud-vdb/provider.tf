terraform {
  required_providers {
    vngcloud = {
      source  = "registry.terraform.nhontt/vngcloud/vngcloud"
      version = "1.0.0"
    }
  }
}

provider "vngcloud" {
  token_url        = "https://hcm-3.api.vngcloud.tech/iam/accounts-api/v2/auth/token"
  client_id        = var.client_id
  client_secret    = var.client_secret
  vserver_base_url = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway"
  vlb_base_url = "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway"
  vdb_base_url = "https://vdb-gateway.vngcloud.vn"
}
