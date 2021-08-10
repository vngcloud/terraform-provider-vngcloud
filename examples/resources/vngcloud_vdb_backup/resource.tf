resource "vngcloud_vdb_backup" "backup" {
  action       = "default"
  name         = "resource_backup"
  description  = "Created from terraform"
  instance_id  = "db-b22df26e-1fa5-4c90-8d7d-2cdd8c126cdc"
  backup_type  = "FULL"
  engine_group = 1
}