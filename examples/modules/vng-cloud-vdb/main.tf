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
  name = "nhontt-tf-config"
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
  engine_type = var.relational_entine_type
  engine_version = var.relational_entine_version
  name = "db.new.s-general-1x2"
}
data "vngcloud_vdb_database_volume_type" "Gen2_NVMe2_IOPS3000" {
  type = "Gen2-NVMe2-IOPS3000"
}
resource "vngcloud_vdb_relational_database" "mysql_8_db" {
  action = "start"
  backup_auto = true
  backup_duration = 2
  backup_time = "00:00"
  config_id = vngcloud_vdb_relational_config_group.mysql_8_config.id
  db_name = var.db_name
  engine_type = var.relational_entine_type
  engine_version = var.relational_entine_version
  name = "tf-db-instance"
  subnet_id = var.subnet_id
  package_id = data.vngcloud_vdb_database_package.db_new_s_general_1x2.id
  password = var.password
  public_access = false
  username = var.user_name
  volume_size = 20
  volume_type = data.vngcloud_vdb_database_volume_type.Gen2_NVMe2_IOPS3000.id
  allowed_ip_prefix = ["10.10.0.0/24"]
}




# Relational Backup
locals {
  mysql_8_db_id = "db-c9e07465-1678-4911-8e00-241cfd21f99a"
  redis_4_db_id = vngcloud_vdb_memstore_database.redis_4_db.id
}
resource "vngcloud_vdb_relational_backup" "mysql_8_backup" {
  name = "tf-backup1"
  instance_id = local.mysql_8_db_id
  backup_type = "FULL"
}
resource "vngcloud_vdb_relational_backup" "mysql_8_backup_incre" {
  name = "tf-backup2"
  instance_id = local.mysql_8_db_id
  backup_type = "INCREMENTAL"
  parent_id = vngcloud_vdb_relational_backup.mysql_8_backup.id
}


# Memory store Database
data "vngcloud_vdb_database_package" "db_new_s_general_1x2_mem" {
  engine_type = var.memmory_engine_type
  engine_version = var.memory_engine_version
  name = "db.new.s-general-1x2"
}
resource "vngcloud_vdb_memstore_database" "redis_4_db" {
  action = "start"
  backup_auto = true
  backup_duration = 2
  backup_time = "00:00"
  config_id = vngcloud_vdb_memstore_config_group.redis_4_config.id
  engine_type = var.memmory_engine_type
  engine_version = var.memory_engine_version
  name = "tf-db-redis"
  subnet_id = var.subnet_id
  package_id = data.vngcloud_vdb_database_package.db_new_s_general_1x2_mem.id
  public_access = true
  redis_password = var.redis_password
  redis_password_enabled = true
  allowed_ip_prefix = ["10.10.0.0/24"]
}




# Memory store Backup
resource "vngcloud_vdb_memstore_backup" "redis_4_backup" {
  name = "tf-redis-backup"
  instance_id = local.redis_4_db_id
}