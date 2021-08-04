terraform {
  required_providers {
    vngcloud = {
      version = "~> 0.2"
      source = "vngcloud.vn/terraform/vngcloud"
    }
  }
}

provider "vngcloud" {
  user_id = "11218"
  project_id = "pro-c4ee389a-582d-40b2-bcd4-9a733425f470"
  vdb_base_url = "https://vcmc-internal.vngcloud.tech/vdb-gateway"
  token_url = "https://vcmc.vngcloud.tech/user-api/v1/user/oauth2/token"
  client_id = "5V23JsOSOwM7puYnMPqxc6ibBqOlpbPI"
  client_secret = "Q4DRR473i2weaE2xmWHR5mR9fBwp8X9K"
}

//Create backup FULL
resource "vngcloud_vdb_backup" "backup" {
  action = "default"
  name = "test_terra_4"
  description = "Created from terraform"
  instance_id = "db-b22df26e-1fa5-4c90-8d7d-2cdd8c126cdc"
  backup_type = "FULL"
  engine_group = 1
}

//Create backup incremetal
resource "vngcloud_vdb_backup" "backup2" {
  action = "default"
  parent_id = resource.vngcloud_vdb_backup.backup.id
  name = "test_terra_5"
  description = "Created from terraform"
  instance_id = "db-b22df26e-1fa5-4c90-8d7d-2cdd8c126cdc"
  backup_type = "INCREMENTAL"
  engine_group = 1
}

data "vngcloud_vdb_package" "package" {
  engine_type = "MySQL"
  engine_version = "5.7"
  name = "v1.small1x1"
}

data "vngcloud_vdb_volume_type" "volume_type" {
  type = "ssd-iops200"
}

data "vngcloud_vdb_network" "network" {
  name = "network_nhontt"
}

data "vngcloud_vdb_subnet" "subnet" {
  name = "subnet_nhontt"
  network_id = data.vngcloud_vdb_network.network.id
}

// Resource database for restore backup (add field backup_id)
resource "vngcloud_vdb_database" "database" {
  action = "reboot"
  backup_auto = false
  backup_duration = 0
  backup_time = ""
  config_id = "cfg-6a7d017a-7687-478b-9b2f-a97c1be2c038"
  database_name = "reficul"
  datastore_type = data.vngcloud_vdb_package.package.engine_type
  datastore_version = data.vngcloud_vdb_package.package.engine_version
  engine_group = 1
  flavor_id = data.vngcloud_vdb_package.package.flavor_id
  name = "nhontt_tf1"
  network_id = data.vngcloud_vdb_subnet.subnet.id
  package_id = data.vngcloud_vdb_package.package.package_id
  password = "abcd1234"
  price_key = data.vngcloud_vdb_package.package.price_key
  ram = data.vngcloud_vdb_package.package.ram
  redis_password = ""
  redis_password_enabled = false
  replica_source_id = ""
  username = "reficul"
  vcpus = data.vngcloud_vdb_package.package.vcpus
  volume_size = 20
  volume_type = data.vngcloud_vdb_volume_type.volume_type.type
  volume_type_zone_id = data.vngcloud_vdb_volume_type.volume_type.volume_type_zone_id
  zone_id = data.vngcloud_vdb_package.package.zone_id
  public_access = false
  backup_id = "bk-2b310228-d4e7-4832-8056-1245af6da2a3"
  # backup_id = resource.vngcloud_vdb_backup.backup.id
}