resource "vngcloud_vdb_backup_storage" "backup_storage" {
  backup_storage_package_id   = 1
  backup_storage_package_name = "db.backup.quota.1"
  engine_group                = 1
  quota                       = 10
}
