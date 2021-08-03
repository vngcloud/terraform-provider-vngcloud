terraform {
  required_providers {
    vdb = {
      version = "~> 0.2"
      source = "vngcloud.vn/terraform/vdb"
    }
  }
}

provider "vdb" {
  user_id = "11218"
  project_id = "pro-c4ee389a-582d-40b2-bcd4-9a733425f470"
  base_url = "http://localhost:8100"
}

# resource "vdb_backup" "backup" {
#   action = "default"
#   name = "test_terra_4"
#   description = "Test api"
#   instance_id = "db-894f1902-b2ac-44c4-8127-68bbffa2ff5e"
#   backup_type = "FULL"
#   engine_group = 1
# }

# resource "vdb_backup" "backup2" {
#   action = "default"
#   parent_id = "db-bk-78cb102d-58e6-4bbf-8d58-f35e2418f1e0"
#   name = "test_terra_3"
#   description = "Test api"
#   instance_id = "db-894f1902-b2ac-44c4-8127-68bbffa2ff5e"
#   backup_type = "INCREMENTAL"
#   engine_group = 1
# }

data "vdb_package" "package" {
  engine_type = "MySQL"
  engine_version = "5.7"
  name = "v1.small1x1"
}

data "vdb_volume_type" "volume_type" {
  type = "ssd-iops200"
}

data "vdb_network" "network" {
  name = "network_nhontt"
}

data "vdb_subnet" "subnet" {
  name = "subnet_nhontt"
  network_id = data.vdb_network.network.id
}

resource "vdb_database" "database" {
  action = "reboot"
  backup_auto = false
  backup_duration = 0
  backup_time = ""
  config_id = "cfg-6a7d017a-7687-478b-9b2f-a97c1be2c038"
  database_name = "reficul"
  datastore_type = data.vdb_package.package.engine_type
  datastore_version = data.vdb_package.package.engine_version
  engine_group = 1
  flavor_id = data.vdb_package.package.flavor_id
  name = "nhontt_tf1"
  network_id = data.vdb_subnet.subnet.id
  package_id = data.vdb_package.package.package_id
  password = "abcd1234"
  price_key = data.vdb_package.package.price_key
  ram = data.vdb_package.package.ram
  redis_password = ""
  redis_password_enabled = false
  replica_source_id = ""
  username = "reficul"
  vcpus = data.vdb_package.package.vcpus
  volume_size = 20
  volume_type = data.vdb_volume_type.volume_type.type
  volume_type_zone_id = data.vdb_volume_type.volume_type.volume_type_zone_id
  zone_id = data.vdb_package.package.zone_id
  backup_id = data.vdb_backup.backup.id
}