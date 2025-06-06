# Relational Config group
resource "vngcloud_vdb_relational_config_group" "mysql_8_config" {
  engine_type = var.relational_engine_type
  engine_version = var.relational_engine_version
  name = "tf-mysql8"
  values = {
    "character_set_server" = "dec8",
    "innodb_stats_method" = "nulls_equal"
  }
}

# Memory store Config group
resource "vngcloud_vdb_memstore_config_group" "redis_4_config" {
  engine_type = var.memmory_engine_type
  engine_version = var.memory_engine_version
  name = "tf-redis4"
  values = {
    "activedefrag" = "true"
  }
}




# Relational Backup storage
data "vngcloud_vdb_backup_storage_package" "db_backup_quota_100GB" {
  name = "db.backup.quota.100GB"
}
data "vngcloud_vdb_backup_storage_package" "db_backup_quota_200GB" {
  name = "db.backup.quota.200GB"
}
data "vngcloud_vdb_backup_storage_package" "db_backup_quota_500GB" {
  name = "db.backup.quota.500GB"
}
resource "vngcloud_vdb_relational_backup_storage" "rela_bk_storage" {
  backup_storage_package_id = data.vngcloud_vdb_backup_storage_package.db_backup_quota_200GB.id
}

# Memory store Backup storage
resource "vngcloud_vdb_memstore_backup_storage" "mem_bk_storage" {
  backup_storage_package_id = data.vngcloud_vdb_backup_storage_package.db_backup_quota_500GB.id
}




# Relational Database
data "vngcloud_vdb_database_package" "db_new_s_general_1x2" {
  engine_type = var.relational_engine_type
  engine_version = var.relational_engine_version
  name = "db.new.s-general-1x2"
}
data "vngcloud_vdb_database_package" "db_new_s_general_2x4" {
  engine_type = var.relational_engine_type
  engine_version = var.relational_engine_version
  name = "db.new.s-general-2x4"
}
data "vngcloud_vdb_database_volume_type" "Gen2_NVMe2_IOPS3000" {
  type = "Gen2-NVMe2-IOPS3000"
}
data "vngcloud_vdb_database_volume_type" "Gen2_NVMe2_IOPS5000" {
  type = "Gen2-NVMe2-IOPS5000"
}
resource "vngcloud_vdb_relational_database" "mysql_8_db" {
  action = ""
  backup_auto = true
  backup_duration = 7
  backup_time = "01:00"
  # config_id = vngcloud_vdb_relational_config_group.mysql_8_config.id
  config_id = ""
  db_name = var.db_name
  engine_type = var.relational_engine_type
  engine_version = var.relational_engine_version
  name = "tf-db-instance"
  subnet_id = var.subnet_id
  package_id = data.vngcloud_vdb_database_package.db_new_s_general_1x2.id
  password = var.password
  public_access = false
  username = var.user_name
  volume_size = 30
  volume_type = data.vngcloud_vdb_database_volume_type.Gen2_NVMe2_IOPS5000.id
  allowed_ip_prefix = ["0.0.0.0/0", "10.10.0.0/24"]
}

# resource "vngcloud_vdb_relational_database" "mysql_8_db_restore" {
#   action = "start"
#   backup_auto = true
#   backup_duration = 7
#   backup_time = "01:00"
#   # config_id = vngcloud_vdb_relational_config_group.mysql_8_config.id
#   config_id = ""
#   db_name = var.db_name
#   engine_type = var.relational_engine_type
#   engine_version = var.relational_engine_version
#   name = "tf-db-restore"
#   subnet_id = var.subnet_id
#   package_id = data.vngcloud_vdb_database_package.db_new_s_general_1x2.id
#   password = var.password
#   public_access = false
#   username = var.user_name
#   volume_size = 30
#   volume_type = data.vngcloud_vdb_database_volume_type.Gen2_NVMe2_IOPS5000.id
#   allowed_ip_prefix = ["0.0.0.0/0", "10.10.0.0/24"]
#   backup_id = local.backup_mysql_8_id
# }




# Relational Backup
locals {
  mysql_8_db_id = vngcloud_vdb_relational_database.mysql_8_db.id
  redis_4_db_id = vngcloud_vdb_memstore_database.redis_4_db.id
  backup_mysql_8_id = vngcloud_vdb_relational_backup.mysql_8_backup.id
  backup_redis_4_id = vngcloud_vdb_memstore_backup.redis_4_backup.id
}
resource "vngcloud_vdb_relational_backup" "mysql_8_backup" {
  name = "tf-backup1"
  instance_id = local.mysql_8_db_id
  backup_type = "FULL"
}
# resource "vngcloud_vdb_relational_backup" "mysql_8_backup_incre" {
#   name = "tf-backup2"
#   instance_id = local.mysql_8_db_id
#   backup_type = "INCREMENTAL"
#   parent_id = vngcloud_vdb_relational_backup.mysql_8_backup.id
# }


# Memory store Database
data "vngcloud_vdb_database_package" "db_new_s_general_1x2_mem" {
  engine_type = var.memmory_engine_type
  engine_version = var.memory_engine_version
  name = "db.new.s-general-1x2"
}
resource "vngcloud_vdb_memstore_database" "redis_4_db" {
  action = "start"
  backup_auto = false
  backup_duration = 2
  backup_time = "00:00"
  config_id = vngcloud_vdb_memstore_config_group.redis_4_config.id
  engine_type = var.memmory_engine_type
  engine_version = var.memory_engine_version
  name = "tf-db-redis"
  subnet_id = var.subnet_id
  package_id = data.vngcloud_vdb_database_package.db_new_s_general_1x2_mem.id
  public_access = false
  redis_password = var.redis_password
  redis_password_enabled = true
  allowed_ip_prefix = ["0.0.0.0/0", "10.10.0.0/24"]
}

import {
  id = "db-48f97f37-9fda-43e3-b01b-fb31b27f015c"
  to = vngcloud_vdb_memstore_database.redis_4_db_rep
}

resource "vngcloud_vdb_memstore_database" "redis_4_db_rep" {
  action = "start"
  backup_auto = false
  backup_duration = 0
  backup_time = ""
  config_id = vngcloud_vdb_memstore_config_group.redis_4_config.id
  engine_type = var.memmory_engine_type
  engine_version = var.memory_engine_version
  name = "tf-db-redis-rep"
  subnet_id = var.subnet_id
  package_id = data.vngcloud_vdb_database_package.db_new_s_general_1x2_mem.id
  public_access = false
  redis_password = var.redis_password2
  redis_password_enabled = true
  allowed_ip_prefix = ["0.0.0.0/0", "10.10.0.0/24"]
  replica_source_id = local.redis_4_db_id
}




# Memory store Backup
resource "vngcloud_vdb_memstore_backup" "redis_4_backup" {
  name = "tf-redis-backup"
  instance_id = local.redis_4_db_id
}
