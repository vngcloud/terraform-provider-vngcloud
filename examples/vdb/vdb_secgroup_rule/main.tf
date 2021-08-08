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

data "vngcloud_vdb_package" "package" {
  engine_type = "MySQL"
  engine_version = "5.7"
  name = "v1.small1x1"
}

data "vngcloud_vdb_volume_type" "volume_type" {
  type = "ssd-iops200"
}

resource "vngcloud_vdb_configuration_group" "config_group" {
  datastore_type = "MySQL"
  datastore_version = "5.7"
  name = "tmp config"
}

resource "vngcloud_vdb_database" "database" {
  action = "reboot"
  allowed_ip_prefix = ["192.168.1.0/30", "192.168.1.0/28", "192.168.1.0/29"]
  backup_auto = false
  backup_duration = 0
  backup_time = ""
  config_id = vngcloud_vdb_configuration_group.config_group.id
  database_name = "nyancat"
  datastore_type = data.vngcloud_vdb_package.package.engine_type
  datastore_version = data.vngcloud_vdb_package.package.engine_version
  engine_group = 1
  flavor_id = data.vngcloud_vdb_package.package.flavor_id
  name = "test db alo"
  network_id = "sub-97b8258b-e64f-48cc-a659-88c8889fd4c3"
  package_id = data.vngcloud_vdb_package.package.package_id
  password = "abcd1234"
  price_key = data.vngcloud_vdb_package.package.price_key
  public_access = true
  ram = data.vngcloud_vdb_package.package.ram
  redis_password = ""
  redis_password_enabled = false
  replica_source_id = ""
  username = "username"
  vcpus = data.vngcloud_vdb_package.package.vcpus
  volume_size = 20
  volume_type = data.vngcloud_vdb_volume_type.volume_type.type
  volume_type_zone_id = data.vngcloud_vdb_volume_type.volume_type.volume_type_zone_id
  zone_id = data.vngcloud_vdb_package.package.zone_id
}

output "my_db" {
  value = vngcloud_vdb_database.database
}

output "secgroup_rules" {
  value = vngcloud_vdb_database.database.allowed_ip_prefix
}