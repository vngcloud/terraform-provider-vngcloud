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

# provider "vdb" {
#   user_id = "11317"
#   project_id = "pro-4a3fcd8d-2459-4a5d-b602-db6552205f78"
#   base_url = "http://localhost:8101"
# }

data "vdb_db" "db" {}

data "vdb_package" "package" {
  engine_type = "MySQL"
  engine_version = "5.7"
  name = "v1.small1x1"
}

data "vdb_volume_type" "volume_type" {
  type = "ssd-iops200"
}

resource "vdb_configuration_group" "config_group" {
  datastore_type = "MySQL"
  datastore_version = "5.7"
  name = "tmp config"
}

resource "vdb_database" "database" {
  action = "reboot"
  allowed_ip_prefix = ["192.168.1.0/30", "192.168.1.0/28", "192.168.1.0/29"]
  backup_auto = false
  backup_duration = 0
  backup_time = ""
  config_id = vdb_configuration_group.config_group.id
  database_name = "nyancat"
  datastore_type = data.vdb_package.package.engine_type
  datastore_version = data.vdb_package.package.engine_version
  engine_group = 1
  flavor_id = data.vdb_package.package.flavor_id
  name = "test db alo"
  network_id = "sub-0ff4f626-4665-422a-a073-8eeae3f02cda"
  package_id = data.vdb_package.package.package_id
  password = "abcd1234"
  price_key = data.vdb_package.package.price_key
  ram = data.vdb_package.package.ram
  redis_password = ""
  redis_password_enabled = false
  replica_source_id = ""
  username = "username"
  vcpus = data.vdb_package.package.vcpus
  volume_size = 20
  volume_type = data.vdb_volume_type.volume_type.type
  volume_type_zone_id = data.vdb_volume_type.volume_type.volume_type_zone_id
  zone_id = data.vdb_package.package.zone_id
}

output "my_db" {
  value = vdb_database.database
}

output "secgroup_rules" {
  value = vdb_database.database.allowed_ip_prefix
}