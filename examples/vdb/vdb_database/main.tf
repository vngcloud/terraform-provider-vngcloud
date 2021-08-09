`terraform {
  required_providers {
    vdb = {
      version = "~> 0.2"
      source  = "vngcloud.vn/terraform/vdb"
    }
  }
}

provider "vdb" {
  user_id       = "11317"
  project_id    = "pro-4a3fcd8d-2459-4a5d-b602-db6552205f78"
  base_url      = "http://localhost:8101"
  token_url     = "https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token"
  client_id     = "eQZyB8nzxfz1cCIQBYMNtGh8nRGF0KV9"
  client_secret = "2fj1zljDbLCupfhsDxCkOJId1cuU1ICn"
}

data "vdb_db" "db" {}

data "vdb_package" "package" {
  engine_type    = "MySQL"
  engine_version = "5.7"
  name           = "v1.small1x1"
}

data "vdb_volume_type" "volume_type" {
  type = "ssd-iops200"
}

data "vdb_network" "network" {
  name = "TriNetwork"
}

data "vdb_subnet" "subnet" {
  name       = "TriSubnet"
  network_id = data.vdb_network.network.id
}

resource "vdb_configuration_group" "config_group" {
  datastore_type    = "MySQL"
  datastore_version = "5.7"
  name              = "trihm4_1_3_C_2"
}

resource "vdb_database" "database" {
  action                 = "reboot"
  backup_auto            = false
  backup_duration        = 0
  backup_time            = ""
  config_id              = vdb_configuration_group.config_group.id
  database_name          = "Tri-Database"
  datastore_type         = data.vdb_package.package.engine_type
  datastore_version      = data.vdb_package.package.engine_version
  engine_group           = 1
  flavor_id              = data.vdb_package.package.flavor_id
  name                   = "trihm4_1_D_3_M_11"
  network_id             = data.vdb_subnet.subnet.id
  package_id             = data.vdb_package.package.package_id
  password               = "password"
  price_key              = data.vdb_package.package.price_key
  public_access          = true
  ram                    = data.vdb_package.package.ram
  redis_password         = "passwordpassword"
  redis_password_enabled = false
  replica_source_id      = ""
  username               = "username"
  vcpus                  = data.vdb_package.package.vcpus
  volume_size            = 20
  volume_type            = data.vdb_volume_type.volume_type.type
  volume_type_zone_id    = data.vdb_volume_type.volume_type.volume_type_zone_id
  zone_id                = data.vdb_package.package.zone_id
}

resource "vdb_database" "database2" {
  action                 = "reboot"
  backup_auto            = false
  backup_duration        = 0
  backup_time            = ""
  config_id              = vdb_database.database.config_id
  database_name          = "Tri-Database"
  datastore_type         = vdb_database.database.datastore_type
  datastore_version      = vdb_database.database.datastore_version
  engine_group           = vdb_database.database.engine_group
  flavor_id              = data.vdb_package.package.flavor_id
  name                   = "trihm4_1_3_M_11_R_1"
  network_id             = vdb_database.database.network_id
  package_id             = data.vdb_package.package.package_id
  password               = "password"
  price_key              = data.vdb_package.package.price_key
  public_access          = false
  ram                    = data.vdb_package.package.ram
  redis_password         = "passwordpassword"
  redis_password_enabled = false
  replica_source_id      = vdb_database.database.id
  username               = "username"
  vcpus                  = data.vdb_package.package.vcpus
  volume_size            = 20
  volume_type            = data.vdb_volume_type.volume_type.type
  volume_type_zone_id    = data.vdb_volume_type.volume_type.volume_type_zone_id
  zone_id                = data.vdb_package.package.zone_id
}
