terraform {
  required_providers {
    vdb = {
      version = "~> 0.2"
      source = "vngcloud.vn/terraform/vdb"
    }
  }
}

provider "vdb" {
  user_id = "11343"
  project_id = "pro-4f394dc3-d426-40e7-870f-395a0172f450"
  base_url = "http://localhost:8101"
}

resource "vdb_configuration_group" "test_config" {
  datastore_type = "MySQL"
  datastore_version = "5.7"
  name = "test configuration group"
}

output "config" {
  value = vdb_configuration_group.test_config
}