data "vngcloud_vdb_backup_storage_package" "backup_storage_package" {
  engine_group = 1
  name         = "db.backup.quota.1"
}