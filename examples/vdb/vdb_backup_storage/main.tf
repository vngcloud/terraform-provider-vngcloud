terraform {
  required_providers {
    vngcloud = {
      version = "~> 0.2"
      source = "vngcloud.vn/terraform/vngcloud"
    }
  }
}

provider "vngcloud" {
  user_id = "11279"
  project_id = "pro-5c16c28a-07fa-4836-bc83-daba5a907d86"
  vdb_base_url = "https://vcmc-internal.vngcloud.tech/vdb-gateway"
  token_url = "https://vcmc.vngcloud.tech/user-api/v1/user/oauth2/token"
  client_id = "5V23JsOSOwM7puYnMPqxc6ibBqOlpbPI"
  client_secret = "Q4DRR473i2weaE2xmWHR5mR9fBwp8X9K"
}

// Backup storage for relational database
data "vngcloud_vdb_backup_storage_package" "backup_storage_package_1"{
  engine_group = 1
  name = "db.backup.quota.1"
}

resource "vngcloud_vdb_backup_storage" "backup_storage_1" {
  backup_storage_package_id = data.vngcloud_vdb_backup_storage_package.backup_storage_package_1.id
  backup_storage_package_name = data.vngcloud_vdb_backup_storage_package.backup_storage_package_1.name
  engine_group = data.vngcloud_vdb_backup_storage_package.backup_storage_package_1.engine_group
  quota = data.vngcloud_vdb_backup_storage_package.backup_storage_package_1.quota
}

// Backup storage for memory store database
data "vngcloud_vdb_backup_storage_package" "backup_storage_package_2"{
  engine_group = 2
  name = "db.backup.quota.1"
}

resource "vngcloud_vdb_backup_storage" "backup_storage_2" {
  backup_storage_package_id = data.vngcloud_vdb_backup_storage_package.backup_storage_package_2.id
  backup_storage_package_name = data.vngcloud_vdb_backup_storage_package.backup_storage_package_2.name
  engine_group = data.vngcloud_vdb_backup_storage_package.backup_storage_package_2.engine_group
  quota = data.vngcloud_vdb_backup_storage_package.backup_storage_package_2.quota
}
