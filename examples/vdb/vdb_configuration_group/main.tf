terraform {
  required_providers {
    vngcloud = {
      version = "~> 0.0.2"
      source = "vngcloud/vngcloud"
    }
  }
}

provider "vngcloud" {
  user_id = "11342"
  project_id = "pro-86a17fba-2530-43f1-890e-d3aaca497c14"
  vdb_base_url = "https://vcmc-internal.vngcloud.tech/vdb-gateway"
  token_url = "https://vcmc.vngcloud.tech/user-api/v1/user/oauth2/token"
  client_id = "5V23JsOSOwM7puYnMPqxc6ibBqOlpbPI"
  client_secret = "Q4DRR473i2weaE2xmWHR5mR9fBwp8X9K"
}

resource "vngcloud_vdb_configuration_group" "test_config" {
  datastore_type = "MySQL"
  datastore_version = "5.7"
  name = "test configuration group"
}

output "config" {
  value = vngcloud_vdb_configuration_group.test_config
}