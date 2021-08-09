resource "vngcloud_vdb_database" "database" {
  name = "MyInstance"

  engine_group      = 1
  datastore_type    = "MySQL"
  datastore_version = "5.7"

  price_key  = "general"
  package_id = "8"
  ram        = 1
  vcpus      = 1
  zone_id    = 1
  flavor_id  = "9"

  volume_type         = "ssd-iops200"
  volume_type_zone_id = "11B9FF75D-00F5-49CA-842D-CCD5D04834FB"
  volume_size         = 20

  database_name = "MyDatabase"
  username      = "username"
  password      = "password"

  redis_password_enabled = true
  redis_password         = "passwordpassword"

  network_id    = "sub-ed965ccb-faf0-4fd5-a685-f6642e4d1b21"
  public_access = true

  config_id = "cfg-6be4adf8-3a1d-46c9-875d-c3d31f7a62bf"

  backup_auto     = true
  backup_duration = 2
  backup_time     = "00:00"

  replica_source_id = "db-ddcdb8a0-aa0e-4f0e-874a-936c69b5dc79"

  action = "start"

  allowed_ip_prefix = ["0.0.0.0/0"]
}